package taskexec

import (
	"testing"
	"time"

	"github.com/gregsemple2003/CodexDesktop/backend/orchestration/internal/taskrun"
)

func TestApplyUpdateMovesRunIntoWaitingForHuman(t *testing.T) {
	dispatchAt := time.Date(2026, time.April, 24, 21, 0, 0, 0, time.UTC)
	view := InitialView(taskrun.StartTaskRunRequest{
		RunID:          "taskrun--Task-0008--active",
		TaskID:         "Task-0008",
		MeaningSummary: "Create the durable backend task-run contract.",
		CapturedTaskSnapshot: taskrun.TaskDefinitionSnapshot{
			DeclaredWorktreeRoot: `C:\Agent\CodexDashboard`,
			DeclaredTaskRoot:     `C:\Agent\CodexDashboard\Tracking\Task-0008`,
			DeclaredTaskRevision: "revision-1",
			DeclaredGitRevision:  "abc123",
			CapturedAt:           dispatchAt,
		},
		RepoLane: taskrun.RepoLane{
			OwnedRepoRoot:         `C:\Temp\owned`,
			CheckoutMode:          "git_worktree_detached",
			BaselineCommit:        "abc123",
			ApprovedRestoreCommit: "abc123",
			ResetStatus:           "not_run",
		},
		DispatchRequestedAt: dispatchAt,
	}, "workflow-id", "run-id")

	applyUpdate(&view, taskrun.TaskRunUpdate{
		State:               taskrun.StateWaitingForHuman,
		ReasonCode:          "review_required",
		StateSummary:        "Run is waiting for human review.",
		NextOwner:           "human",
		NextExpectedEvent:   "Approve the next backend action.",
		LastProgressSummary: "Run recorded a review checkpoint.",
		WaitContract: &taskrun.WaitContract{
			WaitingOn:           "human_review",
			WhyBlocked:          "The next backend action needs human approval.",
			ResumeWhen:          "The human approves the next backend action.",
			HumanActionRequired: true,
			HumanActionTarget: &taskrun.HumanActionTarget{
				Kind:  "approval_action",
				Label: "Approve backend review step",
				URI:   "approval://taskrun/Task-0008",
			},
		},
		Attention: &taskrun.AttentionPriority{
			Level:   taskrun.AttentionNeedsAttention,
			Reason:  "Run is waiting on human review.",
			SortKey: "20-waiting_for_human",
		},
	}, dispatchAt.Add(5*time.Minute))

	if view.StateEnvelope.State != taskrun.StateWaitingForHuman {
		t.Fatalf("state = %q", view.StateEnvelope.State)
	}
	if view.WaitContract == nil || view.WaitContract.HumanActionTarget == nil {
		t.Fatal("expected explicit wait contract and human action target")
	}
	if view.Attention.Level != taskrun.AttentionNeedsAttention {
		t.Fatalf("attention level = %q", view.Attention.Level)
	}
	if view.LastProgressSummary != "Run recorded a review checkpoint." {
		t.Fatalf("last progress summary = %q", view.LastProgressSummary)
	}
}

func TestApplyUpdateMarksInterruptedRunTerminal(t *testing.T) {
	dispatchAt := time.Date(2026, time.April, 24, 21, 0, 0, 0, time.UTC)
	view := InitialView(taskrun.StartTaskRunRequest{
		RunID:          "taskrun--Task-0008--active",
		TaskID:         "Task-0008",
		MeaningSummary: "Create the durable backend task-run contract.",
		CapturedTaskSnapshot: taskrun.TaskDefinitionSnapshot{
			DeclaredWorktreeRoot: `C:\Agent\CodexDashboard`,
			DeclaredTaskRoot:     `C:\Agent\CodexDashboard\Tracking\Task-0008`,
			DeclaredTaskRevision: "revision-1",
			DeclaredGitRevision:  "abc123",
			CapturedAt:           dispatchAt,
		},
		RepoLane: taskrun.RepoLane{
			OwnedRepoRoot:         `C:\Temp\owned`,
			CheckoutMode:          "git_worktree_detached",
			BaselineCommit:        "abc123",
			ApprovedRestoreCommit: "abc123",
			ResetStatus:           "not_run",
		},
		DispatchRequestedAt: dispatchAt,
	}, "workflow-id", "run-id")

	applyUpdate(&view, taskrun.TaskRunUpdate{
		State:        taskrun.StateInterrupted,
		ReasonCode:   "interrupt_requested",
		StateSummary: "Run was interrupted and the owned checkout was restored.",
		CompletedAt:  dispatchAt.Add(7 * time.Minute),
	}, dispatchAt.Add(7*time.Minute))

	if view.Status != "interrupted" {
		t.Fatalf("status = %q", view.Status)
	}
	if !isTerminalStatus(view.Status) {
		t.Fatal("interrupted run should be terminal")
	}
}

func TestApplyUpdateClearsWaitContractWhenRunLeavesHumanWait(t *testing.T) {
	dispatchAt := time.Date(2026, time.April, 24, 21, 0, 0, 0, time.UTC)
	view := InitialView(taskrun.StartTaskRunRequest{
		RunID:          "taskrun--Task-0008--active",
		TaskID:         "Task-0008",
		MeaningSummary: "Create the durable backend task-run contract.",
		CapturedTaskSnapshot: taskrun.TaskDefinitionSnapshot{
			DeclaredWorktreeRoot: `C:\Agent\CodexDashboard`,
			DeclaredTaskRoot:     `C:\Agent\CodexDashboard\Tracking\Task-0008`,
			DeclaredTaskRevision: "revision-1",
			DeclaredGitRevision:  "abc123",
			CapturedAt:           dispatchAt,
		},
		RepoLane: taskrun.RepoLane{
			OwnedRepoRoot:         `C:\Temp\owned`,
			CheckoutMode:          "git_worktree_detached",
			BaselineCommit:        "abc123",
			ApprovedRestoreCommit: "abc123",
			ResetStatus:           "not_run",
		},
		DispatchRequestedAt: dispatchAt,
	}, "workflow-id", "run-id")

	applyUpdate(&view, taskrun.TaskRunUpdate{
		State: taskrun.StateWaitingForHuman,
		WaitContract: &taskrun.WaitContract{
			WaitingOn: "human_review",
		},
	}, dispatchAt.Add(2*time.Minute))
	applyUpdate(&view, taskrun.TaskRunUpdate{
		State:        taskrun.StateBlocked,
		ReasonCode:   "interrupt_cleanup_blocked",
		StateSummary: "Run interrupt could not restore the owned checkout.",
	}, dispatchAt.Add(3*time.Minute))

	if view.WaitContract != nil {
		t.Fatal("wait contract should clear once the run leaves waiting_for_human")
	}
}
