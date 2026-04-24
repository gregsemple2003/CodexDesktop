package taskrun

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"time"
)

type Service struct {
	declaredWorktreeRoot string
	trackingRoot         string
	ownedLaneRoot        string
	runtime              Runtime
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

type parsedTask struct {
	state       taskStateFile
	title       string
	meaning     string
	snapshot    TaskDefinitionSnapshot
	evidenceRef []EvidenceRef
	taskRoot    string
}

func NewService(declaredWorktreeRoot string, runsRoot string, runtime Runtime) *Service {
	return &Service{
		declaredWorktreeRoot: declaredWorktreeRoot,
		trackingRoot:         filepath.Join(declaredWorktreeRoot, "Tracking"),
		ownedLaneRoot:        defaultOwnedLaneRoot(runsRoot),
		runtime:              runtime,
		now: func() time.Time {
			return time.Now().UTC()
		},
	}
}

func (s *Service) ListTasks(ctx context.Context) ([]TaskView, error) {
	entries, err := os.ReadDir(s.trackingRoot)
	if err != nil {
		return nil, fmt.Errorf("read tracking root: %w", err)
	}

	tasks := make([]TaskView, 0)
	for _, entry := range entries {
		if !entry.IsDir() || !strings.HasPrefix(entry.Name(), "Task-") {
			continue
		}
		task, err := s.readTask(ctx, filepath.Join(s.trackingRoot, entry.Name()))
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

func (s *Service) Task(ctx context.Context, taskID string) (TaskView, error) {
	taskRoot := filepath.Join(s.trackingRoot, taskID)
	if _, err := os.Stat(taskRoot); err != nil {
		if os.IsNotExist(err) {
			return TaskView{}, fmt.Errorf("task %q not found", taskID)
		}
		return TaskView{}, err
	}
	return s.readTask(ctx, taskRoot)
}

func (s *Service) Dispatch(ctx context.Context, taskID string) (TaskRunView, error) {
	if s.runtime == nil {
		return TaskRunView{}, fmt.Errorf("task runtime backend is not configured")
	}

	task, err := s.Task(ctx, taskID)
	if err != nil {
		return TaskRunView{}, err
	}
	if !task.DispatchReadiness.Ready {
		return TaskRunView{}, fmt.Errorf("dispatch blocked: %s", summarizeBlockReasons(task.DispatchReadiness.BlockReasons))
	}

	repoLane, err := s.provisionOwnedLane(task.TaskID)
	if err != nil {
		return TaskRunView{}, err
	}

	request := StartTaskRunRequest{
		RunID:          ActiveRunID(task.TaskID),
		TaskID:         task.TaskID,
		Title:          task.Title,
		MeaningSummary: task.MeaningSummary,
		CapturedTaskSnapshot: TaskDefinitionSnapshot{
			DeclaredWorktreeRoot: task.DeclaredWorktreeRoot,
			DeclaredTaskRoot:     task.DeclaredTaskRoot,
			DeclaredTaskRevision: task.DeclaredTaskRevision,
			DeclaredGitRevision:  task.DeclaredGitRevision,
			CapturedAt:           s.now(),
			Files:                nil,
		},
		RepoLane:            repoLane,
		DispatchRequestedAt: s.now(),
	}

	metadata, err := s.loadTask(taskRootForID(s.trackingRoot, task.TaskID))
	if err != nil {
		_ = s.cleanupOwnedLane(repoLane)
		return TaskRunView{}, err
	}
	request.CapturedTaskSnapshot = metadata.snapshot

	run, err := s.runtime.StartTaskRun(ctx, request)
	if err != nil {
		_ = s.cleanupOwnedLane(repoLane)
		return TaskRunView{}, err
	}
	return run, nil
}

func (s *Service) Run(ctx context.Context, runID string) (TaskRunView, error) {
	if s.runtime == nil {
		return TaskRunView{}, fmt.Errorf("task runtime backend is not configured")
	}
	return s.runtime.GetTaskRun(ctx, runID)
}

func (s *Service) UpdateRun(ctx context.Context, runID string, update TaskRunUpdate) (TaskRunView, error) {
	if s.runtime == nil {
		return TaskRunView{}, fmt.Errorf("task runtime backend is not configured")
	}
	if update.Actions == nil && update.State != "" {
		update.Actions = actionsForRunState(update.State)
	}
	if update.Attention == nil && update.State != "" {
		attention := attentionForRunState(update.State)
		update.Attention = &attention
	}
	return s.runtime.UpdateTaskRun(ctx, runID, update)
}

func (s *Service) readTask(ctx context.Context, taskRoot string) (TaskView, error) {
	metadata, err := s.loadTask(taskRoot)
	if err != nil {
		return TaskView{}, err
	}

	view := TaskView{
		TaskID:               metadata.state.TaskID,
		Title:                metadata.title,
		MeaningSummary:       metadata.meaning,
		DeclaredWorktreeRoot: metadata.snapshot.DeclaredWorktreeRoot,
		DeclaredTaskRoot:     metadata.snapshot.DeclaredTaskRoot,
		DeclaredTaskRevision: metadata.snapshot.DeclaredTaskRevision,
		DeclaredGitRevision:  metadata.snapshot.DeclaredGitRevision,
		CurrentStory: StoryOwnership{
			Status: "no_active_run",
			Reason: "No task run currently owns the live story.",
		},
		CurrentGate:  metadata.state.CurrentGate,
		CurrentPass:  metadata.state.CurrentPass,
		Phase:        metadata.state.Phase,
		PlanApproved: metadata.state.PlanApproved,
		Blockers:     append([]string(nil), metadata.state.Blockers...),
		UpdatedAt:    metadata.state.UpdatedAt,
	}

	view.StateEnvelope = s.deriveStateEnvelope(metadata)
	view.DispatchReadiness = s.deriveDispatchReadiness(metadata, false)
	view.Attention = s.deriveAttention(view.StateEnvelope.State, view.DispatchReadiness.Ready)
	view.Actions = defaultActions(view.DispatchReadiness)
	view.StateEnvelope.ActionBlockReasons = collectActionBlockReasons(view.Actions)

	if s.runtime == nil {
		return view, nil
	}

	run, err := s.runtime.GetActiveTaskRun(ctx, metadata.state.TaskID)
	if err != nil {
		if errors.Is(err, ErrRunNotFound) {
			return view, nil
		}
		return TaskView{}, err
	}

	if run.CapturedTaskSnapshot.DeclaredTaskRevision != metadata.snapshot.DeclaredTaskRevision {
		reconciled, reconcileErr := s.runtime.ReconcileTaskSnapshot(ctx, run.RunID, metadata.snapshot)
		if reconcileErr == nil {
			run = reconciled
		}
	}

	view.LatestRun = &run
	view.CurrentStory = StoryOwnership{
		OwnerRunID: run.RunID,
		Status:     "active_run",
		Reason:     "An active task run owns the current live story.",
	}
	view.StateEnvelope = run.StateEnvelope
	view.Attention = run.Attention
	view.Actions = run.Actions
	view.DispatchReadiness = s.deriveDispatchReadiness(metadata, true)
	view.StateEnvelope.ActionBlockReasons = collectActionBlockReasons(view.Actions)

	return view, nil
}

func (s *Service) loadTask(taskRoot string) (parsedTask, error) {
	taskMDPath := filepath.Join(taskRoot, "TASK.md")
	taskRaw, err := os.ReadFile(taskMDPath)
	if err != nil {
		return parsedTask{}, fmt.Errorf("read %s: %w", taskMDPath, err)
	}

	taskStatePath := filepath.Join(taskRoot, "TASK-STATE.json")
	taskStateRaw, err := os.ReadFile(taskStatePath)
	if err != nil {
		return parsedTask{}, fmt.Errorf("read %s: %w", taskStatePath, err)
	}

	var state taskStateFile
	if err := json.Unmarshal(taskStateRaw, &state); err != nil {
		return parsedTask{}, fmt.Errorf("decode %s: %w", taskStatePath, err)
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
		return parsedTask{}, err
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

	return parsedTask{
		state:       state,
		title:       title,
		meaning:     meaning,
		snapshot:    snapshot,
		evidenceRef: evidenceRefs,
		taskRoot:    taskRoot,
	}, nil
}

func (s *Service) deriveStateEnvelope(metadata parsedTask) StateEnvelope {
	now := s.now()
	envelope := StateEnvelope{
		EvidenceRefs: metadata.evidenceRef,
		NextOwner:    "backend",
	}

	switch {
	case len(metadata.state.Blockers) > 0:
		envelope.State = StateBlocked
		envelope.ReasonCode = "task_blocked"
		envelope.StateSummary = "Task is blocked on recorded constraints."
		envelope.NextOwner = "human_or_supervisor"
		envelope.NextExpectedEvent = "Resolve blockers and reassess dispatch readiness."
		envelope.SuspiciousAfter = now.Add(24 * time.Hour)
	case !metadata.state.PlanApproved:
		envelope.State = StateWaitingForHuman
		envelope.ReasonCode = "plan_approval_required"
		envelope.StateSummary = "Task is waiting for plan approval."
		envelope.NextOwner = "human"
		envelope.NextExpectedEvent = "Approve PLAN.md."
		envelope.SuspiciousAfter = now.Add(72 * time.Hour)
	case metadata.state.Status == "done" || metadata.state.Status == "completed" || metadata.state.Status == "closed":
		envelope.State = StateCompleted
		envelope.ReasonCode = "task_complete"
		envelope.StateSummary = "Task is complete."
		envelope.NextOwner = "none"
		envelope.NextExpectedEvent = "No further action is required."
	case metadata.state.Phase == "implementation" && metadata.state.CurrentGate == "implementation":
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
			URI:   fileURI(filepath.Join(metadata.taskRoot, "PLAN.md")),
		})
	}

	return envelope
}

func (s *Service) deriveDispatchReadiness(metadata parsedTask, activeRunExists bool) DispatchReadiness {
	readiness := DispatchReadiness{
		Ready:        false,
		BlockReasons: []ActionBlockReason{},
	}

	if s.runtime == nil {
		readiness.BlockReasons = append(readiness.BlockReasons, ActionBlockReason{
			Code:    "dispatch_runtime_not_implemented",
			Summary: "The durable dispatch lane is not implemented yet.",
		})
	}
	if !metadata.state.PlanApproved {
		readiness.BlockReasons = append(readiness.BlockReasons, ActionBlockReason{
			Code:    "plan_not_approved",
			Summary: "Dispatch requires an approved plan.",
		})
	}
	if len(metadata.state.Blockers) > 0 {
		readiness.BlockReasons = append(readiness.BlockReasons, ActionBlockReason{
			Code:    "task_blockers_present",
			Summary: "Dispatch is blocked until the recorded task blockers are cleared.",
		})
	}
	if activeRunExists {
		readiness.BlockReasons = append(readiness.BlockReasons, ActionBlockReason{
			Code:    "active_run_exists",
			Summary: "Dispatch is blocked while another active run owns the current live story.",
		})
	}
	if _, err := os.Stat(filepath.Join(metadata.taskRoot, "PLAN.md")); err != nil {
		readiness.BlockReasons = append(readiness.BlockReasons, ActionBlockReason{
			Code:    "plan_missing",
			Summary: "Dispatch requires PLAN.md to be present in the declared task root.",
		})
	}
	if _, err := os.Stat(filepath.Join(metadata.taskRoot, "TASK.md")); err != nil {
		readiness.BlockReasons = append(readiness.BlockReasons, ActionBlockReason{
			Code:    "task_missing",
			Summary: "Dispatch requires TASK.md to be present in the declared task root.",
		})
	}
	if metadata.snapshot.DeclaredGitRevision == "" {
		readiness.BlockReasons = append(readiness.BlockReasons, ActionBlockReason{
			Code:    "baseline_commit_unavailable",
			Summary: "Dispatch requires a resolvable git baseline for the declared worktree.",
		})
	}

	if len(readiness.BlockReasons) == 0 {
		readiness.Ready = true
		readiness.ExpectedFirstSignal = "Create a durable backend task run with an owned checkout and captured baseline commit."
		readiness.FirstSuspiciousAfter = s.now().Add(15 * time.Minute)
	}

	return readiness
}

func (s *Service) deriveAttention(state string, dispatchReady bool) AttentionPriority {
	switch {
	case state == StateWaitingForHuman:
		return AttentionPriority{Level: AttentionNeedsAttention, Reason: "Task needs an explicit human action.", SortKey: "20-waiting_for_human"}
	case state == StateBlocked:
		return AttentionPriority{Level: AttentionNeedsAttention, Reason: "Task is blocked and needs review.", SortKey: "30-blocked"}
	case dispatchReady:
		return AttentionPriority{Level: AttentionNeedsAttention, Reason: "Task is ready for dispatch.", SortKey: "40-ready"}
	case state == StateCompleted:
		return AttentionPriority{Level: AttentionNone, Reason: "Task is complete.", SortKey: "90-complete"}
	default:
		return AttentionPriority{Level: AttentionWatch, Reason: "Task should remain visible for backend follow-up.", SortKey: "50-watch"}
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

func (s *Service) provisionOwnedLane(taskID string) (RepoLane, error) {
	baselineCommit := gitRevision(s.declaredWorktreeRoot)
	if baselineCommit == "" {
		return RepoLane{}, fmt.Errorf("resolve baseline commit for %s", taskID)
	}

	stamp := fmt.Sprintf("%x", s.now().UnixNano())
	ownedRepoRoot := filepath.Join(s.ownedLaneRoot, shortTaskSegment(taskID), stamp, "w")
	if err := os.MkdirAll(filepath.Dir(ownedRepoRoot), 0o755); err != nil {
		return RepoLane{}, fmt.Errorf("create owned lane parent: %w", err)
	}

	args := []string{"-C", s.declaredWorktreeRoot}
	if runtime.GOOS == "windows" {
		args = append([]string{"-c", "core.longpaths=true"}, args...)
	}
	args = append(args, "worktree", "add", "--detach", ownedRepoRoot, baselineCommit)
	cmd := exec.Command("git", args...)
	if output, err := cmd.CombinedOutput(); err != nil {
		return RepoLane{}, fmt.Errorf("create owned worktree: %w: %s", err, strings.TrimSpace(string(output)))
	}

	return RepoLane{
		OwnedRepoRoot:         ownedRepoRoot,
		CheckoutMode:          "git_worktree_detached",
		BaselineCommit:        baselineCommit,
		ApprovedRestoreCommit: baselineCommit,
		ResetStatus:           "not_run",
	}, nil
}

func (s *Service) cleanupOwnedLane(repoLane RepoLane) error {
	if repoLane.OwnedRepoRoot == "" {
		return nil
	}
	cmd := exec.Command("git", "-C", s.declaredWorktreeRoot, "worktree", "remove", "--force", repoLane.OwnedRepoRoot)
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("remove owned worktree: %w: %s", err, strings.TrimSpace(string(output)))
	}
	return nil
}

func defaultActions(readiness DispatchReadiness) map[string]ActionAvailability {
	return map[string]ActionAvailability{
		ActionDispatch: {
			Allowed:      readiness.Ready,
			BlockReasons: append([]ActionBlockReason(nil), readiness.BlockReasons...),
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
}

func actionsForRunState(state string) map[string]ActionAvailability {
	dispatchBlocked := []ActionBlockReason{{
		Code:    "active_run_exists",
		Summary: "Dispatch is blocked while this run owns the current live story.",
	}}
	pokeUnavailable := []ActionBlockReason{{
		Code:    "poke_not_implemented",
		Summary: "Poke is not implemented yet for task runs.",
	}}
	interruptUnavailable := []ActionBlockReason{{
		Code:    "interrupt_not_implemented",
		Summary: "Interrupt is not implemented yet for task runs.",
	}}

	switch state {
	case StateRunning, StateDispatching, StateSleepingOrStalled:
		return map[string]ActionAvailability{
			ActionDispatch:  {Allowed: false, BlockReasons: dispatchBlocked},
			ActionPoke:      {Allowed: false, BlockReasons: pokeUnavailable},
			ActionInterrupt: {Allowed: false, BlockReasons: interruptUnavailable},
		}
	case StateWaitingForHuman, StateBlocked, StateCompleted, StateFailed, StateInterrupted:
		return map[string]ActionAvailability{
			ActionDispatch: {
				Allowed:      false,
				BlockReasons: dispatchBlocked,
			},
			ActionPoke: {
				Allowed: false,
				BlockReasons: []ActionBlockReason{{
					Code:    "poke_not_allowed_for_state",
					Summary: "Poke is not allowed in the current run state.",
				}},
			},
			ActionInterrupt: {
				Allowed: false,
				BlockReasons: []ActionBlockReason{{
					Code:    "interrupt_not_allowed_for_state",
					Summary: "Interrupt is not allowed in the current run state.",
				}},
			},
		}
	default:
		return map[string]ActionAvailability{
			ActionDispatch:  {Allowed: false, BlockReasons: dispatchBlocked},
			ActionPoke:      {Allowed: false, BlockReasons: pokeUnavailable},
			ActionInterrupt: {Allowed: false, BlockReasons: interruptUnavailable},
		}
	}
}

func attentionForRunState(state string) AttentionPriority {
	switch state {
	case StateWaitingForHuman:
		return AttentionPriority{Level: AttentionNeedsAttention, Reason: "Run is waiting on a human action.", SortKey: "20-waiting_for_human"}
	case StateBlocked:
		return AttentionPriority{Level: AttentionNeedsAttention, Reason: "Run is blocked and needs review.", SortKey: "30-blocked"}
	case StateSleepingOrStalled:
		return AttentionPriority{Level: AttentionUrgent, Reason: "Run appears stalled.", SortKey: "10-stalled"}
	case StateCompleted:
		return AttentionPriority{Level: AttentionNone, Reason: "Run is complete.", SortKey: "90-complete"}
	case StateFailed:
		return AttentionPriority{Level: AttentionUrgent, Reason: "Run failed.", SortKey: "15-failed"}
	case StateInterrupted:
		return AttentionPriority{Level: AttentionWatch, Reason: "Run was interrupted.", SortKey: "60-interrupted"}
	default:
		return AttentionPriority{Level: AttentionWatch, Reason: "Run is active.", SortKey: "50-active"}
	}
}

func collectActionBlockReasons(actions map[string]ActionAvailability) map[string][]ActionBlockReason {
	blockReasons := map[string][]ActionBlockReason{}
	for action, availability := range actions {
		blockReasons[action] = append([]ActionBlockReason(nil), availability.BlockReasons...)
	}
	return blockReasons
}

func summarizeBlockReasons(reasons []ActionBlockReason) string {
	if len(reasons) == 0 {
		return "unknown reason"
	}
	summaries := make([]string, 0, len(reasons))
	for _, reason := range reasons {
		summaries = append(summaries, reason.Summary)
	}
	return strings.Join(summaries, "; ")
}

func ActiveRunID(taskID string) string {
	return "taskrun--" + sanitizePathSegment(taskID) + "--active"
}

func taskRootForID(trackingRoot string, taskID string) string {
	return filepath.Join(trackingRoot, taskID)
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

func sanitizePathSegment(value string) string {
	replacer := strings.NewReplacer("\\", "_", "/", "_", ":", "_", " ", "_")
	return replacer.Replace(value)
}

func defaultOwnedLaneRoot(runsRoot string) string {
	if runtime.GOOS == "windows" {
		return filepath.Join(os.TempDir(), "cdxow")
	}
	return filepath.Join(runsRoot, "task-owned-checkouts")
}

func shortTaskSegment(taskID string) string {
	hash := sha256.Sum256([]byte(taskID))
	return sanitizePathSegment(taskID) + "-" + hex.EncodeToString(hash[:4])
}
