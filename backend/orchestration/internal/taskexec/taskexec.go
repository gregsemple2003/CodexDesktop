package taskexec

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"

	"github.com/gregsemple2003/CodexDesktop/backend/orchestration/internal/taskrun"
)

const (
	TaskRunWorkflowName         = "codex.task.run"
	TaskRunStateQueryName       = "taskrun.current_state"
	ReconcileSnapshotSignalName = "taskrun.reconcile_snapshot"
	UpdateRunSignalName         = "taskrun.update_state"
	RunExecutionPreflightName   = "taskrun.execution_preflight"
)

func Register(w worker.Worker) {
	w.RegisterWorkflowWithOptions(TaskRunWorkflow, workflow.RegisterOptions{Name: TaskRunWorkflowName})
	w.RegisterActivityWithOptions(runExecutionPreflight, activity.RegisterOptions{Name: RunExecutionPreflightName})
}

func TaskRunWorkflow(ctx workflow.Context, request taskrun.StartTaskRunRequest) (taskrun.TaskRunView, error) {
	now := workflow.Now(ctx).UTC()
	if request.DispatchRequestedAt.IsZero() {
		request.DispatchRequestedAt = now
	}

	info := workflow.GetInfo(ctx)
	view := InitialView(request, info.WorkflowExecution.ID, info.WorkflowExecution.RunID)

	if err := workflow.SetQueryHandler(ctx, TaskRunStateQueryName, func() (taskrun.TaskRunView, error) {
		return view, nil
	}); err != nil {
		return taskrun.TaskRunView{}, err
	}

	if request.RepoLane.RunArtifactRoot != "" {
		activityCtx := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
			StartToCloseTimeout: 2 * time.Minute,
			RetryPolicy: &temporal.RetryPolicy{
				MaximumAttempts: 1,
			},
		})
		var preflight executionPreflightResult
		if err := workflow.ExecuteActivity(activityCtx, RunExecutionPreflightName, request).Get(activityCtx, &preflight); err != nil {
			applyUpdate(&view, taskrun.TaskRunUpdate{
				State:               taskrun.StateBlocked,
				ReasonCode:          "execution_preflight_failed",
				StateSummary:        "Run could not complete owned-lane execution preflight.",
				NextOwner:           "human_or_supervisor",
				NextExpectedEvent:   "Review the preflight failure before continuing execution.",
				SuspiciousAfter:     workflow.Now(ctx).UTC(),
				LastProgressSummary: "Execution preflight failed before the first workload step.",
				Attention: &taskrun.AttentionPriority{
					Level:   taskrun.AttentionUrgent,
					Reason:  "Run could not prepare the owned lane for execution.",
					SortKey: "14-execution_preflight_failed",
				},
				Actions:        actionsForState(taskrun.StateBlocked),
				FailureSummary: err.Error(),
			}, workflow.Now(ctx).UTC())
		} else {
			repoLane := view.RepoLane
			if preflight.CurrentCommit != "" {
				repoLane.CurrentCommit = preflight.CurrentCommit
			}
			if preflight.PreflightArtifactPath != "" {
				repoLane.PreflightArtifactPath = preflight.PreflightArtifactPath
			}
			applyUpdate(&view, taskrun.TaskRunUpdate{
				State:               taskrun.StateRunning,
				ReasonCode:          "execution_preflight_complete",
				StateSummary:        "Run completed owned-lane execution preflight.",
				NextOwner:           "backend_worker",
				NextExpectedEvent:   "Execution worker records the first workload step.",
				SuspiciousAfter:     workflow.Now(ctx).UTC().Add(15 * time.Minute),
				LastProgressSummary: preflight.ProgressSummary,
				Attention: &taskrun.AttentionPriority{
					Level:   taskrun.AttentionWatch,
					Reason:  "Run completed execution preflight and is ready for the first workload step.",
					SortKey: "44-execution_preflight_complete",
				},
				RepoLane: &repoLane,
				Actions:  actionsForState(taskrun.StateRunning),
			}, workflow.Now(ctx).UTC())
		}
	}

	reconcileCh := workflow.GetSignalChannel(ctx, ReconcileSnapshotSignalName)
	updateCh := workflow.GetSignalChannel(ctx, UpdateRunSignalName)
	for {
		selector := workflow.NewSelector(ctx)
		selector.AddReceive(reconcileCh, func(c workflow.ReceiveChannel, more bool) {
			var snapshot taskrun.TaskDefinitionSnapshot
			c.Receive(ctx, &snapshot)
			view.CapturedTaskSnapshot = snapshot
			view.DocRuntimeDivergenceStatus = "reconciled"
			view.DocRuntimeDivergenceSummary = "Runtime captured newer task docs during task readback."
			view.LastProgressAt = workflow.Now(ctx).UTC()
			view.LastProgressSummary = "Reconciled declared task docs into runtime state."
			view.StateEnvelope.SuspiciousAfter = workflow.Now(ctx).UTC().Add(15 * time.Minute)
		})
		selector.AddReceive(updateCh, func(c workflow.ReceiveChannel, more bool) {
			var update taskrun.TaskRunUpdate
			c.Receive(ctx, &update)
			applyUpdate(&view, update, workflow.Now(ctx).UTC())
		})
		selector.Select(ctx)
		if shouldExit(view) {
			return view, nil
		}
	}
}

func InitialView(request taskrun.StartTaskRunRequest, workflowID string, executionRunID string) taskrun.TaskRunView {
	suspiciousAfter := request.DispatchRequestedAt.Add(15 * time.Minute)
	if request.DispatchRequestedAt.IsZero() {
		suspiciousAfter = time.Time{}
	}

	initialState := taskrun.StateDispatching
	initialReasonCode := "dispatch_started"
	initialStateSummary := "Run is dispatching in an owned checkout."
	initialNextOwner := "backend"
	initialNextExpectedEvent := "Execution worker records the next task-run state update."
	initialAttention := taskrun.AttentionPriority{Level: taskrun.AttentionWatch, Reason: "Run is active and waiting for the next backend update.", SortKey: "50-dispatching"}
	initialLastProgressSummary := "Captured task docs and provisioned an owned checkout."

	if request.RepoLane.CurrentCommit != "" {
		initialState = taskrun.StateRunning
		initialReasonCode = "owned_lane_bootstrapped"
		initialStateSummary = "Run bootstrapped the owned checkout and is ready for backend execution."
		initialNextOwner = "backend_worker"
		initialNextExpectedEvent = "Execution worker records the next progress checkpoint."
		initialAttention = taskrun.AttentionPriority{Level: taskrun.AttentionWatch, Reason: "Run is active after owned-lane bootstrap.", SortKey: "45-owned_lane_bootstrapped"}
		initialLastProgressSummary = "Bootstrapped the owned checkout and recorded its current commit."
	}

	actions := map[string]taskrun.ActionAvailability{
		taskrun.ActionDispatch: {
			Allowed: false,
			BlockReasons: []taskrun.ActionBlockReason{{
				Code:    "active_run_exists",
				Summary: "Dispatch is blocked while this run owns the current live story.",
			}},
		},
		taskrun.ActionPoke: {
			Allowed: false,
			BlockReasons: []taskrun.ActionBlockReason{{
				Code:    "poke_not_implemented",
				Summary: "Poke is not implemented yet for task runs.",
			}},
		},
		taskrun.ActionInterrupt: {
			Allowed: false,
			BlockReasons: []taskrun.ActionBlockReason{{
				Code:    "interrupt_not_implemented",
				Summary: "Interrupt is not implemented yet for task runs.",
			}},
		},
	}
	if initialState == taskrun.StateRunning {
		actions[taskrun.ActionPoke] = taskrun.ActionAvailability{
			Allowed: false,
			BlockReasons: []taskrun.ActionBlockReason{{
				Code:    "run_not_suspicious_yet",
				Summary: "Poke stays blocked until the run misses its next expected progress deadline.",
			}},
		}
		actions[taskrun.ActionInterrupt] = taskrun.ActionAvailability{Allowed: true}
	}

	return taskrun.TaskRunView{
		RunID:                  request.RunID,
		TaskID:                 request.TaskID,
		WorkflowID:             workflowID,
		TemporalExecutionRunID: executionRunID,
		Status:                 "active",
		StateEnvelope: taskrun.StateEnvelope{
			State:              initialState,
			ReasonCode:         initialReasonCode,
			StateSummary:       initialStateSummary,
			NextOwner:          initialNextOwner,
			NextExpectedEvent:  initialNextExpectedEvent,
			SuspiciousAfter:    suspiciousAfter,
			ActionBlockReasons: collectActionBlockReasons(actions),
		},
		MeaningSummary:      request.MeaningSummary,
		Attention:           initialAttention,
		Actions:             actions,
		RepoLane:            request.RepoLane,
		LastProgressAt:      request.DispatchRequestedAt,
		LastProgressSummary: initialLastProgressSummary,
		CapturedTaskSnapshot: taskrun.TaskDefinitionSnapshot{
			DeclaredWorktreeRoot: request.CapturedTaskSnapshot.DeclaredWorktreeRoot,
			DeclaredTaskRoot:     request.CapturedTaskSnapshot.DeclaredTaskRoot,
			DeclaredTaskRevision: request.CapturedTaskSnapshot.DeclaredTaskRevision,
			DeclaredGitRevision:  request.CapturedTaskSnapshot.DeclaredGitRevision,
			CapturedAt:           request.CapturedTaskSnapshot.CapturedAt,
			Files:                append([]taskrun.TaskArtifactDigest(nil), request.CapturedTaskSnapshot.Files...),
		},
		DocRuntimeDivergenceStatus:  "in_sync",
		DocRuntimeDivergenceSummary: "Runtime task snapshot matches the declared task docs captured at dispatch.",
	}
}

type executionPreflightResult struct {
	CurrentCommit         string          `json:"current_commit"`
	PreflightArtifactPath string          `json:"preflight_artifact_path"`
	OwnedTaskRoot         string          `json:"owned_task_root"`
	GitStatusShort        string          `json:"git_status_short,omitempty"`
	DocPresence           map[string]bool `json:"doc_presence,omitempty"`
	ProgressSummary       string          `json:"progress_summary"`
}

type executionPreflightArtifact struct {
	TaskID               string          `json:"task_id"`
	RunID                string          `json:"run_id"`
	OwnedRepoRoot        string          `json:"owned_repo_root"`
	OwnedTaskRoot        string          `json:"owned_task_root"`
	DeclaredWorktreeRoot string          `json:"declared_worktree_root"`
	DeclaredTaskRoot     string          `json:"declared_task_root"`
	DeclaredTaskRevision string          `json:"declared_task_revision"`
	DeclaredGitRevision  string          `json:"declared_git_revision,omitempty"`
	CurrentCommit        string          `json:"current_commit"`
	GitStatusShort       string          `json:"git_status_short,omitempty"`
	DocPresence          map[string]bool `json:"doc_presence,omitempty"`
	RecordedAt           time.Time       `json:"recorded_at"`
}

func runExecutionPreflight(ctx context.Context, request taskrun.StartTaskRunRequest) (executionPreflightResult, error) {
	if request.RepoLane.OwnedRepoRoot == "" {
		return executionPreflightResult{}, fmt.Errorf("owned repo root is missing")
	}
	if request.RepoLane.RunArtifactRoot == "" {
		return executionPreflightResult{}, fmt.Errorf("run artifact root is missing")
	}

	ownedTaskRoot, err := ownedTaskRoot(request.CapturedTaskSnapshot, request.RepoLane)
	if err != nil {
		return executionPreflightResult{}, err
	}
	currentCommit, err := gitRevParse(request.RepoLane.OwnedRepoRoot, "HEAD")
	if err != nil {
		return executionPreflightResult{}, err
	}
	gitStatusShort, err := gitStatusShort(request.RepoLane.OwnedRepoRoot)
	if err != nil {
		return executionPreflightResult{}, err
	}
	docPresence := map[string]bool{
		"TASK.md":         pathExists(filepath.Join(ownedTaskRoot, "TASK.md")),
		"PLAN.md":         pathExists(filepath.Join(ownedTaskRoot, "PLAN.md")),
		"HANDOFF.md":      pathExists(filepath.Join(ownedTaskRoot, "HANDOFF.md")),
		"TASK-STATE.json": pathExists(filepath.Join(ownedTaskRoot, "TASK-STATE.json")),
	}

	if err := os.MkdirAll(request.RepoLane.RunArtifactRoot, 0o755); err != nil {
		return executionPreflightResult{}, fmt.Errorf("create run artifact root: %w", err)
	}
	artifactPath := filepath.Join(request.RepoLane.RunArtifactRoot, "execution-preflight.json")
	artifact := executionPreflightArtifact{
		TaskID:               request.TaskID,
		RunID:                request.RunID,
		OwnedRepoRoot:        request.RepoLane.OwnedRepoRoot,
		OwnedTaskRoot:        ownedTaskRoot,
		DeclaredWorktreeRoot: request.CapturedTaskSnapshot.DeclaredWorktreeRoot,
		DeclaredTaskRoot:     request.CapturedTaskSnapshot.DeclaredTaskRoot,
		DeclaredTaskRevision: request.CapturedTaskSnapshot.DeclaredTaskRevision,
		DeclaredGitRevision:  request.CapturedTaskSnapshot.DeclaredGitRevision,
		CurrentCommit:        currentCommit,
		GitStatusShort:       gitStatusShort,
		DocPresence:          docPresence,
		RecordedAt:           time.Now().UTC(),
	}
	if err := writeJSONArtifact(artifactPath, artifact); err != nil {
		return executionPreflightResult{}, err
	}

	return executionPreflightResult{
		CurrentCommit:         currentCommit,
		PreflightArtifactPath: artifactPath,
		OwnedTaskRoot:         ownedTaskRoot,
		GitStatusShort:        gitStatusShort,
		DocPresence:           docPresence,
		ProgressSummary:       "Execution preflight inspected the owned task docs and recorded owned-lane readiness.",
	}, nil
}

func collectActionBlockReasons(actions map[string]taskrun.ActionAvailability) map[string][]taskrun.ActionBlockReason {
	blockReasons := map[string][]taskrun.ActionBlockReason{}
	for action, availability := range actions {
		blockReasons[action] = append([]taskrun.ActionBlockReason(nil), availability.BlockReasons...)
	}
	return blockReasons
}

func applyUpdate(view *taskrun.TaskRunView, update taskrun.TaskRunUpdate, now time.Time) {
	if update.State != "" {
		view.StateEnvelope.State = update.State
	}
	if update.ReasonCode != "" {
		view.StateEnvelope.ReasonCode = update.ReasonCode
	}
	if update.StateSummary != "" {
		view.StateEnvelope.StateSummary = update.StateSummary
	}
	if update.NextOwner != "" {
		view.StateEnvelope.NextOwner = update.NextOwner
	}
	if update.NextExpectedEvent != "" {
		view.StateEnvelope.NextExpectedEvent = update.NextExpectedEvent
	}
	if !update.SuspiciousAfter.IsZero() {
		view.StateEnvelope.SuspiciousAfter = update.SuspiciousAfter
	}
	if update.LastProgressSummary != "" {
		view.LastProgressSummary = update.LastProgressSummary
		view.LastProgressAt = now
	}
	if update.WaitContract != nil {
		view.WaitContract = update.WaitContract
	} else if update.State != "" && update.State != taskrun.StateWaitingForHuman {
		view.WaitContract = nil
	}
	if update.Attention != nil {
		view.Attention = *update.Attention
	}
	if update.FollowUp != nil {
		if isEmptyFollowUp(update.FollowUp) {
			view.FollowUp = nil
		} else {
			view.FollowUp = update.FollowUp
		}
	}
	if update.Resolution != nil {
		view.Resolution = update.Resolution
	}
	if update.RepoLane != nil {
		view.RepoLane = *update.RepoLane
	}
	if update.Actions != nil {
		view.Actions = update.Actions
		view.StateEnvelope.ActionBlockReasons = collectActionBlockReasons(update.Actions)
	}
	if update.CompletedAt.IsZero() {
		view.Status = "active"
	} else {
		view.LastProgressAt = update.CompletedAt
	}
	switch view.StateEnvelope.State {
	case taskrun.StateCompleted:
		view.Status = "completed"
	case taskrun.StateFailed:
		view.Status = "failed"
	case taskrun.StateInterrupted:
		view.Status = "interrupted"
	}
	if update.FailureSummary != "" {
		view.FailureSummary = update.FailureSummary
	} else if update.State != "" && update.State != taskrun.StateBlocked && update.State != taskrun.StateFailed {
		view.FailureSummary = ""
	}
}

func isTerminalStatus(status string) bool {
	return status == "completed" || status == "failed" || status == "interrupted"
}

func shouldExit(view taskrun.TaskRunView) bool {
	if view.Status == "interrupted" && hasPendingInterruptReview(view) {
		return false
	}
	return isTerminalStatus(view.Status)
}

func actionsForState(state string) map[string]taskrun.ActionAvailability {
	switch state {
	case taskrun.StateRunning, taskrun.StateDispatching:
		return map[string]taskrun.ActionAvailability{
			taskrun.ActionDispatch: {
				Allowed: false,
				BlockReasons: []taskrun.ActionBlockReason{{
					Code:    "active_run_exists",
					Summary: "Dispatch is blocked while this run owns the current live story.",
				}},
			},
			taskrun.ActionPoke: {
				Allowed: false,
				BlockReasons: []taskrun.ActionBlockReason{{
					Code:    "run_not_suspicious_yet",
					Summary: "Poke stays blocked until the run misses its next expected progress deadline.",
				}},
			},
			taskrun.ActionInterrupt: {Allowed: true},
		}
	case taskrun.StateBlocked:
		return map[string]taskrun.ActionAvailability{
			taskrun.ActionDispatch: {
				Allowed: false,
				BlockReasons: []taskrun.ActionBlockReason{{
					Code:    "active_run_exists",
					Summary: "Dispatch is blocked while this run owns the current live story.",
				}},
			},
			taskrun.ActionPoke: {
				Allowed: false,
				BlockReasons: []taskrun.ActionBlockReason{{
					Code:    "poke_not_allowed_for_state",
					Summary: "Poke is not allowed in the current run state.",
				}},
			},
			taskrun.ActionInterrupt: {Allowed: true},
		}
	default:
		return map[string]taskrun.ActionAvailability{
			taskrun.ActionDispatch: {
				Allowed: false,
				BlockReasons: []taskrun.ActionBlockReason{{
					Code:    "active_run_exists",
					Summary: "Dispatch is blocked while this run owns the current live story.",
				}},
			},
			taskrun.ActionPoke: {
				Allowed: false,
				BlockReasons: []taskrun.ActionBlockReason{{
					Code:    "poke_not_implemented",
					Summary: "Poke is not implemented yet for task runs.",
				}},
			},
			taskrun.ActionInterrupt: {
				Allowed: false,
				BlockReasons: []taskrun.ActionBlockReason{{
					Code:    "interrupt_not_implemented",
					Summary: "Interrupt is not implemented yet for task runs.",
				}},
			},
		}
	}
}

func hasPendingInterruptReview(view taskrun.TaskRunView) bool {
	return view.FollowUp != nil &&
		view.FollowUp.Kind == "interrupt_review" &&
		(view.FollowUp.Status == "pending" || view.FollowUp.Status == "overdue")
}

func isEmptyFollowUp(followUp *taskrun.RunFollowUp) bool {
	return followUp != nil &&
		followUp.Kind == "" &&
		followUp.Owner == "" &&
		followUp.Status == "" &&
		followUp.Summary == "" &&
		followUp.RequestedAt.IsZero() &&
		followUp.DueAt.IsZero() &&
		followUp.CompletedAt.IsZero()
}

func ownedTaskRoot(snapshot taskrun.TaskDefinitionSnapshot, repoLane taskrun.RepoLane) (string, error) {
	rel, err := filepath.Rel(snapshot.DeclaredWorktreeRoot, snapshot.DeclaredTaskRoot)
	if err != nil {
		return "", fmt.Errorf("resolve task root relative path: %w", err)
	}
	if rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
		return "", fmt.Errorf("declared task root %q is outside declared worktree root %q", snapshot.DeclaredTaskRoot, snapshot.DeclaredWorktreeRoot)
	}
	return filepath.Join(repoLane.OwnedRepoRoot, rel), nil
}

func gitRevParse(worktreeRoot string, ref string) (string, error) {
	out, err := exec.Command("git", "-C", worktreeRoot, "rev-parse", ref).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("git rev-parse %s: %w: %s", ref, err, strings.TrimSpace(string(out)))
	}
	return strings.TrimSpace(string(out)), nil
}

func gitStatusShort(worktreeRoot string) (string, error) {
	out, err := exec.Command("git", "-C", worktreeRoot, "status", "--short").CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("git status --short: %w: %s", err, strings.TrimSpace(string(out)))
	}
	return strings.TrimSpace(string(out)), nil
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func writeJSONArtifact(path string, value any) error {
	data, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return err
	}
	data = append(data, '\n')
	if err := os.WriteFile(path, data, 0o644); err != nil {
		return fmt.Errorf("write %s: %w", path, err)
	}
	return nil
}
