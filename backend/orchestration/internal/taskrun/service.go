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
	"reflect"
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
	run, err := s.runtime.GetTaskRun(ctx, runID)
	if err != nil {
		return TaskRunView{}, err
	}
	return s.refreshRun(ctx, run)
}

func (s *Service) UpdateRun(ctx context.Context, runID string, update TaskRunUpdate) (TaskRunView, error) {
	if s.runtime == nil {
		return TaskRunView{}, fmt.Errorf("task runtime backend is not configured")
	}
	current, err := s.runtime.GetTaskRun(ctx, runID)
	if err != nil {
		return TaskRunView{}, err
	}
	now := s.now()
	if update.FollowUp == nil {
		update.FollowUp = derivedFollowUp(current, update, now)
	}
	projected := projectRun(current, update, now)
	if update.Actions == nil {
		update.Actions = actionsForRun(projected, now)
	}
	if update.Attention == nil {
		attention := attentionForRun(projected, now)
		update.Attention = &attention
	}
	return s.runtime.UpdateTaskRun(ctx, runID, update)
}

func (s *Service) PokeRun(ctx context.Context, runID string) (TaskRunView, error) {
	run, err := s.Run(ctx, runID)
	if err != nil {
		return TaskRunView{}, err
	}
	availability := run.Actions[ActionPoke]
	if !availability.Allowed {
		return TaskRunView{}, fmt.Errorf("poke blocked: %s", summarizeBlockReasons(availability.BlockReasons))
	}

	now := s.now()
	update := TaskRunUpdate{
		State:               StateSleepingOrStalled,
		ReasonCode:          "poke_requested",
		StateSummary:        "Run was poked and is waiting for a fresh backend progress signal.",
		NextOwner:           "backend",
		NextExpectedEvent:   "Execution worker records a fresh progress update or explicit wait reason.",
		SuspiciousAfter:     now.Add(10 * time.Minute),
		LastProgressSummary: "Backend requested a fresh status update for the stalled run.",
		FollowUp: &RunFollowUp{
			Kind:        "poke_worker_check",
			Owner:       "backend_worker",
			Status:      "pending",
			Summary:     "Execution worker should acknowledge the poke with fresh progress or an explicit wait reason.",
			RequestedAt: now,
			DueAt:       now.Add(5 * time.Minute),
		},
	}
	return s.UpdateRun(ctx, runID, update)
}

func (s *Service) InterruptRun(ctx context.Context, runID string) (TaskRunView, error) {
	run, err := s.Run(ctx, runID)
	if err != nil {
		return TaskRunView{}, err
	}
	availability := run.Actions[ActionInterrupt]
	if !availability.Allowed {
		return TaskRunView{}, fmt.Errorf("interrupt blocked: %s", summarizeBlockReasons(availability.BlockReasons))
	}

	repoLane, resetErr := s.restoreOwnedLane(run.RepoLane)
	if resetErr != nil {
		update := TaskRunUpdate{
			State:               StateBlocked,
			ReasonCode:          "interrupt_cleanup_blocked",
			StateSummary:        "Run interrupt could not restore the owned checkout.",
			NextOwner:           "human_or_supervisor",
			NextExpectedEvent:   "Review cleanup failure and resolve the owned checkout manually.",
			LastProgressSummary: "Interrupt cleanup failed and the owned checkout needs manual review.",
			FollowUp: &RunFollowUp{
				Kind:        "cleanup_repair",
				Owner:       "human_or_supervisor",
				Status:      "pending",
				Summary:     "Repair the cleanup-blocked owned checkout before attempting another interrupt or redispatch.",
				RequestedAt: s.now(),
				DueAt:       s.now().Add(24 * time.Hour),
			},
			RepoLane:       &repoLane,
			FailureSummary: resetErr.Error(),
		}
		return s.UpdateRun(ctx, runID, update)
	}

	now := s.now()
	update := TaskRunUpdate{
		State:               StateInterrupted,
		ReasonCode:          "interrupt_requested",
		StateSummary:        "Run was interrupted and the owned checkout was restored.",
		NextOwner:           "human_or_supervisor",
		NextExpectedEvent:   "Review the interrupted run and decide whether to dispatch again.",
		SuspiciousAfter:     now,
		LastProgressSummary: "Interrupt restored the owned checkout to its recorded restore commit.",
		FollowUp: &RunFollowUp{
			Kind:        "interrupt_review",
			Owner:       "human_or_supervisor",
			Status:      "pending",
			Summary:     "Review the interrupted run and decide whether to redispatch, revise the task docs, or close the attempt.",
			RequestedAt: now,
			DueAt:       now.Add(24 * time.Hour),
		},
		RepoLane:    &repoLane,
		CompletedAt: now,
	}
	return s.UpdateRun(ctx, runID, update)
}

func (s *Service) RetryCleanupRun(ctx context.Context, runID string) (TaskRunView, error) {
	run, err := s.Run(ctx, runID)
	if err != nil {
		return TaskRunView{}, err
	}
	if run.StateEnvelope.State != StateBlocked || run.StateEnvelope.ReasonCode != "interrupt_cleanup_blocked" {
		return TaskRunView{}, fmt.Errorf("cleanup retry blocked: run is not waiting on cleanup repair")
	}

	repoLane, resetErr := s.restoreOwnedLane(run.RepoLane)
	now := s.now()
	if resetErr != nil {
		update := TaskRunUpdate{
			State:               StateBlocked,
			ReasonCode:          "interrupt_cleanup_blocked",
			StateSummary:        "Cleanup retry could not restore the owned checkout.",
			NextOwner:           "human_or_supervisor",
			NextExpectedEvent:   "Repair the owned checkout or retry cleanup again.",
			LastProgressSummary: "Backend cleanup retry failed and the owned checkout still needs repair.",
			FollowUp: &RunFollowUp{
				Kind:        "cleanup_repair",
				Owner:       "human_or_supervisor",
				Status:      "pending",
				Summary:     "Repair the cleanup-blocked owned checkout or retry cleanup again after the restore target is valid.",
				RequestedAt: now,
				DueAt:       now.Add(24 * time.Hour),
			},
			RepoLane:       &repoLane,
			FailureSummary: resetErr.Error(),
		}
		return s.UpdateRun(ctx, runID, update)
	}

	update := TaskRunUpdate{
		State:               StateInterrupted,
		ReasonCode:          "interrupt_cleanup_repaired",
		StateSummary:        "Cleanup retry restored the owned checkout and the run now needs interrupt review.",
		NextOwner:           "human_or_supervisor",
		NextExpectedEvent:   "Review the repaired interrupted run and decide whether to dispatch again.",
		SuspiciousAfter:     now,
		LastProgressSummary: "Cleanup retry restored the owned checkout to its recorded restore commit.",
		FollowUp: &RunFollowUp{
			Kind:        "interrupt_review",
			Owner:       "human_or_supervisor",
			Status:      "pending",
			Summary:     "Cleanup repair completed; review the interrupted run and decide whether to redispatch, revise the task docs, or close the attempt.",
			RequestedAt: now,
			DueAt:       now.Add(24 * time.Hour),
		},
		RepoLane:    &repoLane,
		CompletedAt: now,
	}
	return s.UpdateRun(ctx, runID, update)
}

func (s *Service) ResolveInterruptReview(ctx context.Context, runID string, resolution InterruptReviewResolution) (TaskRunView, error) {
	run, err := s.Run(ctx, runID)
	if err != nil {
		return TaskRunView{}, err
	}
	if !hasPendingInterruptReview(run) {
		return TaskRunView{}, fmt.Errorf("interrupt review resolution blocked: run is not waiting on interrupt review")
	}

	now := s.now()
	resolvedBy := strings.TrimSpace(resolution.ResolvedBy)
	if resolvedBy == "" {
		resolvedBy = "human_or_supervisor"
	}

	decision := strings.TrimSpace(resolution.Decision)
	switch decision {
	case "redispatch_ready":
		summary := strings.TrimSpace(resolution.Summary)
		if summary == "" {
			summary = "Interrupt review approved the run for a later redispatch."
		}
		return s.UpdateRun(ctx, runID, TaskRunUpdate{
			State:               StateInterrupted,
			ReasonCode:          "interrupt_review_resolved_redispatch_ready",
			StateSummary:        "Interrupt review approved the run for redispatch.",
			NextOwner:           "backend",
			NextExpectedEvent:   "Dispatch a new run when the task is ready.",
			LastProgressSummary: summary,
			FollowUp: &RunFollowUp{
				Kind:        "interrupt_review",
				Owner:       "human_or_supervisor",
				Status:      "completed",
				Summary:     summary,
				RequestedAt: run.FollowUp.RequestedAt,
				DueAt:       run.FollowUp.DueAt,
				CompletedAt: now,
			},
			Resolution: &RunResolution{
				Kind:       "interrupt_review",
				Decision:   decision,
				Summary:    summary,
				ResolvedBy: resolvedBy,
				ResolvedAt: now,
			},
		})
	case "keep_closed":
		summary := strings.TrimSpace(resolution.Summary)
		if summary == "" {
			summary = "Interrupt review closed this interrupted attempt without redispatch."
		}
		return s.UpdateRun(ctx, runID, TaskRunUpdate{
			State:               StateInterrupted,
			ReasonCode:          "interrupt_review_resolved_keep_closed",
			StateSummary:        "Interrupt review closed this interrupted attempt.",
			NextOwner:           "none",
			NextExpectedEvent:   "No further action is required for this run.",
			LastProgressSummary: summary,
			FollowUp: &RunFollowUp{
				Kind:        "interrupt_review",
				Owner:       "human_or_supervisor",
				Status:      "completed",
				Summary:     summary,
				RequestedAt: run.FollowUp.RequestedAt,
				DueAt:       run.FollowUp.DueAt,
				CompletedAt: now,
			},
			Resolution: &RunResolution{
				Kind:       "interrupt_review",
				Decision:   decision,
				Summary:    summary,
				ResolvedBy: resolvedBy,
				ResolvedAt: now,
			},
		})
	default:
		return TaskRunView{}, fmt.Errorf("interrupt review resolution blocked: unsupported decision %q", decision)
	}
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
	run, err = s.refreshRun(ctx, run)
	if err != nil {
		return TaskView{}, err
	}

	view.LatestRun = &run
	if runOwnsLiveStory(run) {
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
	} else {
		view.CurrentStory = StoryOwnership{
			Status: "no_active_run",
			Reason: "The latest task run is terminal and no run currently owns the live story.",
		}
		if hasPendingInterruptReview(run) {
			view.StateEnvelope = StateEnvelope{
				State:             StateWaitingForHuman,
				ReasonCode:        "interrupt_review_pending",
				StateSummary:      "Task is waiting for interrupt review before redispatch.",
				EvidenceRefs:      metadata.evidenceRef,
				NextOwner:         "human_or_supervisor",
				NextExpectedEvent: "Resolve the interrupted run review decision.",
				SuspiciousAfter:   run.FollowUp.DueAt,
			}
			view.DispatchReadiness = DispatchReadiness{
				Ready: false,
				BlockReasons: []ActionBlockReason{{
					Code:    "interrupt_review_pending",
					Summary: "Dispatch stays blocked until the interrupted run review is resolved.",
				}},
			}
			view.Attention = attentionForRun(run, s.now())
			view.Actions = defaultActions(view.DispatchReadiness)
			view.StateEnvelope.ActionBlockReasons = collectActionBlockReasons(view.Actions)
		} else {
			view.DispatchReadiness = s.deriveDispatchReadiness(metadata, false)
			view.Actions = defaultActions(view.DispatchReadiness)
			view.StateEnvelope.ActionBlockReasons = collectActionBlockReasons(view.Actions)
		}
	}

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

func (s *Service) refreshRun(ctx context.Context, run TaskRunView) (TaskRunView, error) {
	if s.runtime == nil {
		return run, nil
	}
	update := s.derivedRunUpdate(run)
	if update == nil {
		return run, nil
	}
	return s.runtime.UpdateTaskRun(ctx, run.RunID, *update)
}

func (s *Service) derivedRunUpdate(run TaskRunView) *TaskRunUpdate {
	now := s.now()
	desiredActions := actionsForRun(run, now)
	desiredAttention := attentionForRun(run, now)
	desiredFollowUp := desiredFollowUp(run, now)

	var update TaskRunUpdate
	changed := false

	if staleUpdate, ok := staleRunUpdate(run, now, desiredActions, desiredAttention); ok {
		return &staleUpdate
	}

	if !reflect.DeepEqual(run.Actions, desiredActions) {
		update.Actions = desiredActions
		changed = true
	}
	if run.Attention != desiredAttention {
		update.Attention = &desiredAttention
		changed = true
	}
	if !reflect.DeepEqual(run.FollowUp, desiredFollowUp) {
		update.FollowUp = desiredFollowUp
		changed = true
	}
	if !changed {
		return nil
	}
	return &update
}

func (s *Service) provisionOwnedLane(taskID string) (RepoLane, error) {
	baselineCommit := gitRevision(s.declaredWorktreeRoot)
	if baselineCommit == "" {
		return RepoLane{}, fmt.Errorf("resolve baseline commit for %s", taskID)
	}

	if err := os.MkdirAll(s.ownedLaneRoot, 0o755); err != nil {
		return RepoLane{}, fmt.Errorf("create owned lane root: %w", err)
	}
	laneDir, err := os.MkdirTemp(s.ownedLaneRoot, shortTaskSegment(taskID)+"-")
	if err != nil {
		return RepoLane{}, fmt.Errorf("create owned lane temp dir: %w", err)
	}
	ownedRepoRoot := filepath.Join(laneDir, "w")

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

func (s *Service) restoreOwnedLane(repoLane RepoLane) (RepoLane, error) {
	now := s.now()
	repoLane.LastResetAt = now
	restoreCommit := repoLane.ApprovedRestoreCommit
	if restoreCommit == "" {
		restoreCommit = repoLane.BaselineCommit
	}
	repoLane.LastResetTargetCommit = restoreCommit
	repoLane.ResetFailureSummary = ""
	if repoLane.OwnedRepoRoot == "" {
		repoLane.ResetStatus = "cleanup_blocked"
		repoLane.ResetFailureSummary = "Owned repo root is missing."
		return repoLane, fmt.Errorf("owned repo root is missing")
	}
	if restoreCommit == "" {
		repoLane.ResetStatus = "cleanup_blocked"
		repoLane.ResetFailureSummary = "Restore commit is missing."
		return repoLane, fmt.Errorf("restore commit is missing")
	}
	if !pathWithinRoot(repoLane.OwnedRepoRoot, s.ownedLaneRoot) {
		repoLane.ResetStatus = "cleanup_blocked"
		repoLane.ResetFailureSummary = fmt.Sprintf("Owned repo root %q is outside the backend-owned lane root.", repoLane.OwnedRepoRoot)
		return repoLane, fmt.Errorf("owned repo root %q is outside the backend-owned lane root", repoLane.OwnedRepoRoot)
	}
	if err := gitInWorktree(repoLane.OwnedRepoRoot, "reset", "--hard", restoreCommit); err != nil {
		repoLane.ResetStatus = "cleanup_blocked"
		repoLane.ResetFailureSummary = fmt.Sprintf("Reset to %s failed.", restoreCommit)
		return repoLane, fmt.Errorf("reset owned lane to %s: %w", restoreCommit, err)
	}
	if err := gitInWorktree(repoLane.OwnedRepoRoot, "clean", "-fd"); err != nil {
		repoLane.ResetStatus = "cleanup_blocked"
		repoLane.ResetFailureSummary = "Git clean failed while restoring the owned checkout."
		return repoLane, fmt.Errorf("clean owned lane: %w", err)
	}
	repoLane.ResetStatus = "restored"
	repoLane.ApprovedRestoreCommit = restoreCommit
	repoLane.ResetFailureSummary = ""
	return repoLane, nil
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
	pokeUnavailable := []ActionBlockReason{{Code: "poke_not_allowed_for_state", Summary: "Poke is not allowed in the current run state."}}
	interruptAllowed := ActionAvailability{Allowed: true}
	interruptUnavailable := []ActionBlockReason{{Code: "interrupt_not_allowed_for_state", Summary: "Interrupt is not allowed in the current run state."}}

	switch state {
	case StateRunning, StateDispatching:
		return map[string]ActionAvailability{
			ActionDispatch:  {Allowed: false, BlockReasons: dispatchBlocked},
			ActionPoke:      {Allowed: false, BlockReasons: pokeUnavailable},
			ActionInterrupt: interruptAllowed,
		}
	case StateSleepingOrStalled:
		return map[string]ActionAvailability{
			ActionDispatch:  {Allowed: false, BlockReasons: dispatchBlocked},
			ActionPoke:      {Allowed: true},
			ActionInterrupt: interruptAllowed,
		}
	case StateWaitingForHuman, StateBlocked:
		return map[string]ActionAvailability{
			ActionDispatch:  {Allowed: false, BlockReasons: dispatchBlocked},
			ActionPoke:      {Allowed: false, BlockReasons: pokeUnavailable},
			ActionInterrupt: interruptAllowed,
		}
	case StateCompleted, StateFailed, StateInterrupted:
		return map[string]ActionAvailability{
			ActionDispatch: {Allowed: false, BlockReasons: dispatchBlocked},
			ActionPoke: {Allowed: false, BlockReasons: []ActionBlockReason{{
				Code:    "run_terminal",
				Summary: "Poke is not allowed after the run has already ended.",
			}}},
			ActionInterrupt: {Allowed: false, BlockReasons: []ActionBlockReason{{
				Code:    "run_terminal",
				Summary: "Interrupt is not allowed after the run has already ended.",
			}}},
		}
	default:
		return map[string]ActionAvailability{
			ActionDispatch:  {Allowed: false, BlockReasons: dispatchBlocked},
			ActionPoke:      {Allowed: false, BlockReasons: pokeUnavailable},
			ActionInterrupt: {Allowed: false, BlockReasons: interruptUnavailable},
		}
	}
}

func actionsForRun(run TaskRunView, now time.Time) map[string]ActionAvailability {
	actions := actionsForRunState(run.StateEnvelope.State)
	if run.FollowUp != nil && run.FollowUp.Status == "pending" && run.FollowUp.Kind == "poke_worker_check" {
		actions[ActionPoke] = ActionAvailability{
			Allowed: false,
			BlockReasons: []ActionBlockReason{{
				Code:    "follow_up_pending",
				Summary: "Poke is already waiting on a backend-worker follow-up.",
			}},
		}
	}
	if run.StateEnvelope.State == StateRunning || run.StateEnvelope.State == StateDispatching {
		if !run.StateEnvelope.SuspiciousAfter.IsZero() && now.After(run.StateEnvelope.SuspiciousAfter) {
			actions[ActionPoke] = ActionAvailability{Allowed: true}
		} else {
			actions[ActionPoke] = ActionAvailability{
				Allowed: false,
				BlockReasons: []ActionBlockReason{{
					Code:    "run_not_suspicious_yet",
					Summary: "Poke stays blocked until the run misses its next expected progress deadline.",
				}},
			}
		}
	}
	if run.StateEnvelope.State == StateWaitingForHuman && run.WaitContract != nil && !run.WaitContract.StaleAfter.IsZero() && now.After(run.WaitContract.StaleAfter) {
		actions[ActionPoke] = ActionAvailability{
			Allowed: false,
			BlockReasons: []ActionBlockReason{{
				Code:    "waiting_for_human_stale",
				Summary: "Poke does not replace the required human action on a stale human wait.",
			}},
		}
	}
	return actions
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

func attentionForRun(run TaskRunView, now time.Time) AttentionPriority {
	if run.FollowUp != nil && run.FollowUp.Status == "overdue" {
		switch run.FollowUp.Owner {
		case "backend_worker":
			return AttentionPriority{Level: AttentionUrgent, Reason: "A backend-worker follow-up is overdue.", SortKey: "11-follow_up_overdue"}
		default:
			return AttentionPriority{Level: AttentionUrgent, Reason: "A required follow-up is overdue.", SortKey: "13-follow_up_overdue"}
		}
	}
	if run.StateEnvelope.State == StateWaitingForHuman && run.WaitContract != nil && !run.WaitContract.StaleAfter.IsZero() && now.After(run.WaitContract.StaleAfter) {
		return AttentionPriority{Level: AttentionUrgent, Reason: "Run is still waiting on a human action past its stale deadline.", SortKey: "12-waiting_stale"}
	}
	if hasPendingInterruptReview(run) {
		return AttentionPriority{Level: AttentionNeedsAttention, Reason: "Interrupted run is still waiting on review resolution.", SortKey: "25-interrupt_review_pending"}
	}
	if run.StateEnvelope.State == StateInterrupted && run.Resolution != nil {
		return AttentionPriority{Level: AttentionNone, Reason: "Interrupted run review is resolved.", SortKey: "85-interrupt_review_resolved"}
	}
	return attentionForRunState(run.StateEnvelope.State)
}

func desiredFollowUp(run TaskRunView, now time.Time) *RunFollowUp {
	if run.FollowUp == nil {
		return nil
	}
	followUp := *run.FollowUp
	if followUp.Status == "pending" && !followUp.DueAt.IsZero() && now.After(followUp.DueAt) {
		followUp.Status = "overdue"
		return &followUp
	}
	return run.FollowUp
}

func hasPendingInterruptReview(run TaskRunView) bool {
	return run.StateEnvelope.State == StateInterrupted &&
		run.FollowUp != nil &&
		run.FollowUp.Kind == "interrupt_review" &&
		(run.FollowUp.Status == "pending" || run.FollowUp.Status == "overdue")
}

func staleRunUpdate(run TaskRunView, now time.Time, actions map[string]ActionAvailability, attention AttentionPriority) (TaskRunUpdate, bool) {
	if (run.StateEnvelope.State == StateRunning || run.StateEnvelope.State == StateDispatching) &&
		!run.StateEnvelope.SuspiciousAfter.IsZero() &&
		now.After(run.StateEnvelope.SuspiciousAfter) {
		return TaskRunUpdate{
			State:               StateSleepingOrStalled,
			ReasonCode:          "progress_stale",
			StateSummary:        "Run has gone quiet past its expected progress window.",
			NextOwner:           "backend",
			NextExpectedEvent:   "Poke or interrupt the run.",
			SuspiciousAfter:     run.StateEnvelope.SuspiciousAfter,
			LastProgressSummary: "Supervision marked the run as sleeping or stalled.",
			Attention:           &attention,
			Actions:             actionsForRunState(StateSleepingOrStalled),
		}, true
	}
	if run.StateEnvelope.State == StateWaitingForHuman &&
		run.WaitContract != nil &&
		!run.WaitContract.StaleAfter.IsZero() &&
		now.After(run.WaitContract.StaleAfter) &&
		run.StateEnvelope.ReasonCode != "human_wait_stale" {
		return TaskRunUpdate{
			State:               StateWaitingForHuman,
			ReasonCode:          "human_wait_stale",
			StateSummary:        "Run is still waiting for human input and the wait has gone stale.",
			NextOwner:           "human_or_supervisor",
			NextExpectedEvent:   "Review the stale human wait or interrupt the run.",
			SuspiciousAfter:     run.WaitContract.StaleAfter,
			LastProgressSummary: "Supervision marked the human wait as stale.",
			Attention:           &attention,
			Actions:             actions,
		}, true
	}
	return TaskRunUpdate{}, false
}

func projectRun(current TaskRunView, update TaskRunUpdate, now time.Time) TaskRunView {
	projected := current
	if update.State != "" {
		projected.StateEnvelope.State = update.State
	}
	if update.ReasonCode != "" {
		projected.StateEnvelope.ReasonCode = update.ReasonCode
	}
	if update.StateSummary != "" {
		projected.StateEnvelope.StateSummary = update.StateSummary
	}
	if update.NextOwner != "" {
		projected.StateEnvelope.NextOwner = update.NextOwner
	}
	if update.NextExpectedEvent != "" {
		projected.StateEnvelope.NextExpectedEvent = update.NextExpectedEvent
	}
	if !update.SuspiciousAfter.IsZero() {
		projected.StateEnvelope.SuspiciousAfter = update.SuspiciousAfter
	}
	if update.WaitContract != nil {
		projected.WaitContract = update.WaitContract
	} else if update.State != "" && update.State != StateWaitingForHuman {
		projected.WaitContract = nil
	}
	if update.RepoLane != nil {
		projected.RepoLane = *update.RepoLane
	}
	if update.FollowUp != nil {
		if isEmptyRunFollowUp(update.FollowUp) {
			projected.FollowUp = nil
		} else {
			projected.FollowUp = update.FollowUp
		}
	}
	if update.Resolution != nil {
		projected.Resolution = update.Resolution
	}
	if update.LastProgressSummary != "" {
		projected.LastProgressSummary = update.LastProgressSummary
		projected.LastProgressAt = now
	}
	if update.FailureSummary != "" {
		projected.FailureSummary = update.FailureSummary
	} else if update.State != "" && update.State != StateBlocked && update.State != StateFailed {
		projected.FailureSummary = ""
	}
	if !update.CompletedAt.IsZero() {
		projected.LastProgressAt = update.CompletedAt
	}
	switch projected.StateEnvelope.State {
	case StateCompleted:
		projected.Status = "completed"
	case StateFailed:
		projected.Status = "failed"
	case StateInterrupted:
		projected.Status = "interrupted"
	default:
		projected.Status = "active"
	}
	return projected
}

func derivedFollowUp(current TaskRunView, update TaskRunUpdate, now time.Time) *RunFollowUp {
	if update.FollowUp != nil {
		return update.FollowUp
	}
	if current.FollowUp == nil {
		return nil
	}
	if current.FollowUp.Owner == "backend_worker" && current.FollowUp.Status != "completed" {
		effectiveState := current.StateEnvelope.State
		if update.State != "" {
			effectiveState = update.State
		}
		if update.LastProgressSummary != "" && update.ReasonCode != "poke_requested" && effectiveState != StateSleepingOrStalled {
			completed := *current.FollowUp
			completed.Status = "completed"
			completed.CompletedAt = now
			completed.Summary = "Backend worker follow-up completed with a fresh runtime update."
			return &completed
		}
	}
	return current.FollowUp
}

func isEmptyRunFollowUp(followUp *RunFollowUp) bool {
	return followUp != nil &&
		followUp.Kind == "" &&
		followUp.Owner == "" &&
		followUp.Status == "" &&
		followUp.Summary == "" &&
		followUp.RequestedAt.IsZero() &&
		followUp.DueAt.IsZero() &&
		followUp.CompletedAt.IsZero()
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

func runOwnsLiveStory(run TaskRunView) bool {
	return run.Status != "completed" && run.Status != "failed" && run.Status != "interrupted"
}

func pathWithinRoot(path string, root string) bool {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return false
	}
	absRoot, err := filepath.Abs(root)
	if err != nil {
		return false
	}
	rel, err := filepath.Rel(absRoot, absPath)
	if err != nil {
		return false
	}
	return rel != ".." && !strings.HasPrefix(rel, ".."+string(filepath.Separator))
}

func gitInWorktree(worktreeRoot string, args ...string) error {
	argv := []string{}
	if runtime.GOOS == "windows" {
		argv = append(argv, "-c", "core.longpaths=true")
	}
	argv = append(argv, "-C", worktreeRoot)
	argv = append(argv, args...)
	cmd := exec.Command("git", argv...)
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("%w: %s", err, strings.TrimSpace(string(output)))
	}
	return nil
}
