package taskexec

import (
	"time"

	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"

	"github.com/gregsemple2003/CodexDesktop/backend/orchestration/internal/taskrun"
)

const (
	TaskRunWorkflowName         = "codex.task.run"
	TaskRunStateQueryName       = "taskrun.current_state"
	ReconcileSnapshotSignalName = "taskrun.reconcile_snapshot"
	UpdateRunSignalName         = "taskrun.update_state"
)

func Register(w worker.Worker) {
	w.RegisterWorkflowWithOptions(TaskRunWorkflow, workflow.RegisterOptions{Name: TaskRunWorkflowName})
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
