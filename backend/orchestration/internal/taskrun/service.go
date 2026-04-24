package taskrun

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type Service struct {
	declaredWorktreeRoot string
	trackingRoot         string
	now                  func() time.Time
}

type taskStateFile struct {
	TaskID       string   `json:"task_id"`
	Status       string   `json:"status"`
	Phase        string   `json:"phase"`
	PlanApproved bool     `json:"plan_approved"`
	CurrentPass  string   `json:"current_pass"`
	CurrentGate  string   `json:"current_gate"`
	Blockers     []string `json:"blockers"`
	UpdatedAt    string   `json:"updated_at"`
}

func NewService(declaredWorktreeRoot string) *Service {
	return &Service{
		declaredWorktreeRoot: declaredWorktreeRoot,
		trackingRoot:         filepath.Join(declaredWorktreeRoot, "Tracking"),
		now: func() time.Time {
			return time.Now().UTC()
		},
	}
}

func (s *Service) ListTasks(_ context.Context) ([]TaskView, error) {
	entries, err := os.ReadDir(s.trackingRoot)
	if err != nil {
		return nil, fmt.Errorf("read tracking root: %w", err)
	}

	tasks := make([]TaskView, 0)
	for _, entry := range entries {
		if !entry.IsDir() || !strings.HasPrefix(entry.Name(), "Task-") {
			continue
		}
		task, err := s.readTask(filepath.Join(s.trackingRoot, entry.Name()))
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].TaskID < tasks[j].TaskID
	})
	return tasks, nil
}

func (s *Service) Task(_ context.Context, taskID string) (TaskView, error) {
	taskRoot := filepath.Join(s.trackingRoot, taskID)
	if _, err := os.Stat(taskRoot); err != nil {
		if os.IsNotExist(err) {
			return TaskView{}, fmt.Errorf("task %q not found", taskID)
		}
		return TaskView{}, err
	}
	return s.readTask(taskRoot)
}

func (s *Service) readTask(taskRoot string) (TaskView, error) {
	taskMDPath := filepath.Join(taskRoot, "TASK.md")
	taskRaw, err := os.ReadFile(taskMDPath)
	if err != nil {
		return TaskView{}, fmt.Errorf("read %s: %w", taskMDPath, err)
	}

	taskStatePath := filepath.Join(taskRoot, "TASK-STATE.json")
	taskStateRaw, err := os.ReadFile(taskStatePath)
	if err != nil {
		return TaskView{}, fmt.Errorf("read %s: %w", taskStatePath, err)
	}

	var state taskStateFile
	if err := json.Unmarshal(taskStateRaw, &state); err != nil {
		return TaskView{}, fmt.Errorf("decode %s: %w", taskStatePath, err)
	}

	title := extractMarkdownSection(string(taskRaw), "Title")
	if title == "" {
		title = strings.TrimSpace(state.TaskID)
	}
	meaning := firstParagraph(extractMarkdownSection(string(taskRaw), "Summary"))
	if meaning == "" {
		meaning = title
	}

	snapshot, err := s.captureSnapshot(taskRoot)
	if err != nil {
		return TaskView{}, err
	}

	evidenceRefs := []EvidenceRef{
		taskArtifactRef("TASK.md", taskMDPath),
		taskArtifactRef("TASK-STATE.json", taskStatePath),
	}
	if _, err := os.Stat(filepath.Join(taskRoot, "PLAN.md")); err == nil {
		evidenceRefs = append(evidenceRefs, taskArtifactRef("PLAN.md", filepath.Join(taskRoot, "PLAN.md")))
	}
	if _, err := os.Stat(filepath.Join(taskRoot, "HANDOFF.md")); err == nil {
		evidenceRefs = append(evidenceRefs, taskArtifactRef("HANDOFF.md", filepath.Join(taskRoot, "HANDOFF.md")))
	}
	if _, err := os.Stat(filepath.Join(taskRoot, "CONSTRAINTS.md")); err == nil {
		evidenceRefs = append(evidenceRefs, taskArtifactRef("CONSTRAINTS.md", filepath.Join(taskRoot, "CONSTRAINTS.md")))
	}

	view := TaskView{
		TaskID:               state.TaskID,
		Title:                title,
		MeaningSummary:       meaning,
		DeclaredWorktreeRoot: snapshot.DeclaredWorktreeRoot,
		DeclaredTaskRoot:     snapshot.DeclaredTaskRoot,
		DeclaredTaskRevision: snapshot.DeclaredTaskRevision,
		DeclaredGitRevision:  snapshot.DeclaredGitRevision,
		CurrentStory: StoryOwnership{
			Status: "no_active_run",
			Reason: "No task run currently owns the live story.",
		},
		CurrentGate:  state.CurrentGate,
		CurrentPass:  state.CurrentPass,
		Phase:        state.Phase,
		PlanApproved: state.PlanApproved,
		Blockers:     append([]string(nil), state.Blockers...),
		UpdatedAt:    state.UpdatedAt,
	}

	view.StateEnvelope = s.deriveStateEnvelope(view, state, evidenceRefs, taskRoot)
	view.DispatchReadiness = s.deriveDispatchReadiness(state, taskRoot)
	view.Attention = s.deriveAttention(view.StateEnvelope.State, view.DispatchReadiness.Ready)
	view.Actions = map[string]ActionAvailability{
		ActionDispatch: {
			Allowed:      view.DispatchReadiness.Ready,
			BlockReasons: append([]ActionBlockReason(nil), view.DispatchReadiness.BlockReasons...),
		},
		ActionPoke: {
			Allowed: false,
			BlockReasons: []ActionBlockReason{{
				Code:    "no_active_run",
				Summary: "Poke is unavailable until a task run exists.",
			}},
		},
		ActionInterrupt: {
			Allowed: false,
			BlockReasons: []ActionBlockReason{{
				Code:    "no_active_run",
				Summary: "Interrupt is unavailable until a task run exists.",
			}},
		},
	}
	view.StateEnvelope.ActionBlockReasons = map[string][]ActionBlockReason{
		ActionDispatch:  append([]ActionBlockReason(nil), view.Actions[ActionDispatch].BlockReasons...),
		ActionPoke:      append([]ActionBlockReason(nil), view.Actions[ActionPoke].BlockReasons...),
		ActionInterrupt: append([]ActionBlockReason(nil), view.Actions[ActionInterrupt].BlockReasons...),
	}

	return view, nil
}

func (s *Service) deriveStateEnvelope(view TaskView, state taskStateFile, evidenceRefs []EvidenceRef, taskRoot string) StateEnvelope {
	now := s.now()
	envelope := StateEnvelope{
		EvidenceRefs: evidenceRefs,
		NextOwner:    "backend",
	}

	switch {
	case len(state.Blockers) > 0:
		envelope.State = StateBlocked
		envelope.ReasonCode = "task_blocked"
		envelope.StateSummary = "Task is blocked on recorded constraints."
		envelope.NextOwner = "human_or_supervisor"
		envelope.NextExpectedEvent = "Resolve blockers and reassess dispatch readiness."
		envelope.SuspiciousAfter = now.Add(24 * time.Hour)
	case !state.PlanApproved:
		envelope.State = StateWaitingForHuman
		envelope.ReasonCode = "plan_approval_required"
		envelope.StateSummary = "Task is waiting for plan approval."
		envelope.NextOwner = "human"
		envelope.NextExpectedEvent = "Approve PLAN.md."
		envelope.SuspiciousAfter = now.Add(72 * time.Hour)
	case state.Status == "done" || state.Status == "completed" || state.Status == "closed":
		envelope.State = StateCompleted
		envelope.ReasonCode = "task_complete"
		envelope.StateSummary = "Task is complete."
		envelope.NextOwner = "none"
		envelope.NextExpectedEvent = "No further action is required."
	case state.Phase == "implementation" && state.CurrentGate == "implementation":
		envelope.State = StateReady
		envelope.ReasonCode = "ready_for_dispatch"
		envelope.StateSummary = "Task is ready for backend dispatch."
		envelope.NextExpectedEvent = "Dispatch a backend task run."
		envelope.SuspiciousAfter = now.Add(12 * time.Hour)
	default:
		envelope.State = StateBlocked
		envelope.ReasonCode = "task_state_unmapped"
		envelope.StateSummary = "Task needs backend review before dispatch."
		envelope.NextOwner = "backend"
		envelope.NextExpectedEvent = "Review task docs and runtime contract."
		envelope.SuspiciousAfter = now.Add(24 * time.Hour)
	}

	if envelope.State == StateWaitingForHuman && envelope.ReasonCode == "plan_approval_required" {
		envelope.EvidenceRefs = append(envelope.EvidenceRefs, EvidenceRef{
			Type:  "task_artifact",
			Label: "PLAN approval target",
			URI:   fileURI(filepath.Join(taskRoot, "PLAN.md")),
		})
	}

	return envelope
}

func (s *Service) deriveDispatchReadiness(state taskStateFile, taskRoot string) DispatchReadiness {
	readiness := DispatchReadiness{
		Ready:        false,
		BlockReasons: []ActionBlockReason{},
	}

	if !state.PlanApproved {
		readiness.BlockReasons = append(readiness.BlockReasons, ActionBlockReason{
			Code:    "plan_not_approved",
			Summary: "Dispatch requires an approved plan.",
		})
	}
	if len(state.Blockers) > 0 {
		readiness.BlockReasons = append(readiness.BlockReasons, ActionBlockReason{
			Code:    "task_blockers_present",
			Summary: "Dispatch is blocked until the recorded task blockers are cleared.",
		})
	}
	if _, err := os.Stat(filepath.Join(taskRoot, "PLAN.md")); err != nil {
		readiness.BlockReasons = append(readiness.BlockReasons, ActionBlockReason{
			Code:    "plan_missing",
			Summary: "Dispatch requires PLAN.md to be present in the declared task root.",
		})
	}
	if _, err := os.Stat(filepath.Join(taskRoot, "TASK.md")); err != nil {
		readiness.BlockReasons = append(readiness.BlockReasons, ActionBlockReason{
			Code:    "task_missing",
			Summary: "Dispatch requires TASK.md to be present in the declared task root.",
		})
	}
	if len(readiness.BlockReasons) == 0 {
		readiness.BlockReasons = append(readiness.BlockReasons, ActionBlockReason{
			Code:    "dispatch_runtime_not_implemented",
			Summary: "The durable dispatch lane is not implemented yet.",
		})
	}

	return readiness
}

func (s *Service) deriveAttention(state string, dispatchReady bool) AttentionPriority {
	switch {
	case state == StateWaitingForHuman:
		return AttentionPriority{
			Level:   AttentionNeedsAttention,
			Reason:  "Task needs an explicit human action.",
			SortKey: "20-waiting_for_human",
		}
	case state == StateBlocked:
		return AttentionPriority{
			Level:   AttentionNeedsAttention,
			Reason:  "Task is blocked and needs review.",
			SortKey: "30-blocked",
		}
	case dispatchReady:
		return AttentionPriority{
			Level:   AttentionNeedsAttention,
			Reason:  "Task is ready for dispatch.",
			SortKey: "40-ready",
		}
	case state == StateCompleted:
		return AttentionPriority{
			Level:   AttentionNone,
			Reason:  "Task is complete.",
			SortKey: "90-complete",
		}
	default:
		return AttentionPriority{
			Level:   AttentionWatch,
			Reason:  "Task should remain visible for backend follow-up.",
			SortKey: "50-watch",
		}
	}
}

func (s *Service) captureSnapshot(taskRoot string) (TaskDefinitionSnapshot, error) {
	paths := []string{
		filepath.Join(taskRoot, "TASK.md"),
		filepath.Join(taskRoot, "PLAN.md"),
		filepath.Join(taskRoot, "HANDOFF.md"),
		filepath.Join(taskRoot, "TASK-STATE.json"),
		filepath.Join(taskRoot, "CONSTRAINTS.md"),
	}

	digests := make([]TaskArtifactDigest, 0, len(paths))
	hash := sha256.New()
	for _, path := range paths {
		raw, err := os.ReadFile(path)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return TaskDefinitionSnapshot{}, fmt.Errorf("read snapshot file %s: %w", path, err)
		}
		relativePath, err := filepath.Rel(taskRoot, path)
		if err != nil {
			return TaskDefinitionSnapshot{}, fmt.Errorf("rel snapshot path %s: %w", path, err)
		}
		fileHash := sha256.Sum256(raw)
		digests = append(digests, TaskArtifactDigest{
			RelativePath: filepath.ToSlash(relativePath),
			SHA256:       hex.EncodeToString(fileHash[:]),
		})
		hash.Write([]byte(filepath.ToSlash(relativePath)))
		hash.Write(fileHash[:])
	}

	return TaskDefinitionSnapshot{
		DeclaredWorktreeRoot: s.declaredWorktreeRoot,
		DeclaredTaskRoot:     taskRoot,
		DeclaredTaskRevision: hex.EncodeToString(hash.Sum(nil)),
		DeclaredGitRevision:  gitRevision(s.declaredWorktreeRoot),
		CapturedAt:           s.now(),
		Files:                digests,
	}, nil
}

func extractMarkdownSection(markdown string, heading string) string {
	lines := strings.Split(markdown, "\n")
	header := "## " + heading
	capture := false
	section := make([]string, 0)
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == header {
			capture = true
			continue
		}
		if capture && strings.HasPrefix(trimmed, "## ") {
			break
		}
		if capture {
			section = append(section, line)
		}
	}
	return strings.TrimSpace(strings.Join(section, "\n"))
}

func firstParagraph(section string) string {
	lines := strings.Split(section, "\n")
	paragraph := make([]string, 0)
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			if len(paragraph) > 0 {
				break
			}
			continue
		}
		paragraph = append(paragraph, trimmed)
	}
	return strings.Join(paragraph, " ")
}

func taskArtifactRef(label string, path string) EvidenceRef {
	return EvidenceRef{
		Type:  "task_artifact",
		Label: label,
		URI:   fileURI(path),
	}
}

func fileURI(path string) string {
	value := filepath.ToSlash(path)
	return (&url.URL{Scheme: "file", Path: "/" + strings.TrimPrefix(value, "/")}).String()
}

func gitRevision(worktreeRoot string) string {
	out, err := exec.Command("git", "-C", worktreeRoot, "rev-parse", "HEAD").Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(out))
}
