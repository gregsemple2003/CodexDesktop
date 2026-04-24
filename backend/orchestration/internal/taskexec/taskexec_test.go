package taskexec

import (
	"context"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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

func TestInitialViewUsesBootstrappedOwnedLaneAsRunning(t *testing.T) {
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
			CurrentCommit:         "abc123",
			ApprovedRestoreCommit: "abc123",
			RunArtifactRoot:       `C:\Temp\artifacts`,
			BootstrapArtifactPath: `C:\Temp\artifacts\owned-lane-bootstrap.json`,
			ResetStatus:           "not_run",
		},
		DispatchRequestedAt: dispatchAt,
	}, "workflow-id", "run-id")

	if view.StateEnvelope.State != taskrun.StateRunning {
		t.Fatalf("state = %q, want %q", view.StateEnvelope.State, taskrun.StateRunning)
	}
	if view.StateEnvelope.ReasonCode != "owned_lane_bootstrapped" {
		t.Fatalf("reason code = %q", view.StateEnvelope.ReasonCode)
	}
	if !view.Actions[taskrun.ActionInterrupt].Allowed {
		t.Fatal("interrupt should be allowed after owned-lane bootstrap")
	}
	if view.RepoLane.CurrentCommit != "abc123" {
		t.Fatalf("current commit = %q", view.RepoLane.CurrentCommit)
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

func TestApplyUpdateStoresRunFollowUp(t *testing.T) {
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
		State: taskrun.StateSleepingOrStalled,
		FollowUp: &taskrun.RunFollowUp{
			Kind:        "poke_worker_check",
			Owner:       "backend_worker",
			Status:      "pending",
			Summary:     "Execution worker should acknowledge the poke.",
			RequestedAt: dispatchAt.Add(2 * time.Minute),
			DueAt:       dispatchAt.Add(7 * time.Minute),
		},
	}, dispatchAt.Add(2*time.Minute))

	if view.FollowUp == nil || view.FollowUp.Kind != "poke_worker_check" {
		t.Fatalf("follow-up = %#v", view.FollowUp)
	}
}

func TestApplyUpdateClearsExplicitEmptyFollowUp(t *testing.T) {
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
		FollowUp: &taskrun.RunFollowUp{
			Kind:   "cleanup_repair",
			Owner:  "human_or_supervisor",
			Status: "pending",
		},
	}, dispatchAt.Add(1*time.Minute))
	applyUpdate(&view, taskrun.TaskRunUpdate{
		FollowUp: &taskrun.RunFollowUp{},
	}, dispatchAt.Add(2*time.Minute))

	if view.FollowUp != nil {
		t.Fatalf("follow-up should be cleared, got %#v", view.FollowUp)
	}
}

func TestApplyUpdateStoresRunResolution(t *testing.T) {
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
		State: taskrun.StateInterrupted,
		Resolution: &taskrun.RunResolution{
			Kind:       "interrupt_review",
			Decision:   "redispatch_ready",
			Summary:    "Human review approved another dispatch attempt.",
			ResolvedBy: "human",
			ResolvedAt: dispatchAt.Add(9 * time.Minute),
		},
	}, dispatchAt.Add(9*time.Minute))

	if view.Resolution == nil || view.Resolution.Decision != "redispatch_ready" {
		t.Fatalf("resolution = %#v", view.Resolution)
	}
}

func TestShouldExitKeepsInterruptReviewPendingRunAlive(t *testing.T) {
	view := taskrun.TaskRunView{
		Status: "interrupted",
		FollowUp: &taskrun.RunFollowUp{
			Kind:   "interrupt_review",
			Status: "pending",
		},
	}

	if shouldExit(view) {
		t.Fatal("shouldExit returned true for interrupted run with pending interrupt review")
	}

	view.FollowUp.Status = "completed"
	if !shouldExit(view) {
		t.Fatal("shouldExit returned false after interrupt review completed")
	}
}

func TestRunExecutionPreflightWritesArtifactFromOwnedLane(t *testing.T) {
	worktreeRoot := t.TempDir()
	writeTaskFile(t, filepath.Join(worktreeRoot, "Tracking", "Task-0008", "TASK.md"), "# task\n")
	writeTaskFile(t, filepath.Join(worktreeRoot, "Tracking", "Task-0008", "PLAN.md"), "# plan\n")
	writeTaskFile(t, filepath.Join(worktreeRoot, "Tracking", "Task-0008", "HANDOFF.md"), "# handoff\n")
	writeTaskFile(t, filepath.Join(worktreeRoot, "Tracking", "Task-0008", "TASK-STATE.json"), "{\"task_id\":\"Task-0008\"}\n")
	runTaskExecCommand(t, worktreeRoot, "git", "init")
	runTaskExecCommand(t, worktreeRoot, "git", "config", "user.email", "taskexec-tests@example.com")
	runTaskExecCommand(t, worktreeRoot, "git", "config", "user.name", "TaskExec Tests")
	runTaskExecCommand(t, worktreeRoot, "git", "add", ".")
	runTaskExecCommand(t, worktreeRoot, "git", "commit", "-m", "initial")

	ownedRoot := filepath.Join(t.TempDir(), "owned")
	runTaskExecCommand(t, worktreeRoot, "git", "worktree", "add", "--detach", ownedRoot, "HEAD")

	runArtifactRoot := filepath.Join(t.TempDir(), "artifacts")
	request := taskrun.StartTaskRunRequest{
		RunID:          "taskrun--Task-0008--active",
		TaskID:         "Task-0008",
		MeaningSummary: "Create the durable backend task-run contract.",
		CapturedTaskSnapshot: taskrun.TaskDefinitionSnapshot{
			DeclaredWorktreeRoot: worktreeRoot,
			DeclaredTaskRoot:     filepath.Join(worktreeRoot, "Tracking", "Task-0008"),
			DeclaredTaskRevision: "revision-1",
			DeclaredGitRevision:  stringsTrim(runTaskExecOutput(t, worktreeRoot, "git", "rev-parse", "HEAD")),
			CapturedAt:           time.Date(2026, time.April, 24, 21, 0, 0, 0, time.UTC),
		},
		RepoLane: taskrun.RepoLane{
			OwnedRepoRoot:         ownedRoot,
			CheckoutMode:          "git_worktree_detached",
			BaselineCommit:        stringsTrim(runTaskExecOutput(t, worktreeRoot, "git", "rev-parse", "HEAD")),
			CurrentCommit:         stringsTrim(runTaskExecOutput(t, ownedRoot, "git", "rev-parse", "HEAD")),
			ApprovedRestoreCommit: stringsTrim(runTaskExecOutput(t, worktreeRoot, "git", "rev-parse", "HEAD")),
			RunArtifactRoot:       runArtifactRoot,
			BootstrapArtifactPath: filepath.Join(runArtifactRoot, "owned-lane-bootstrap.json"),
		},
		DispatchRequestedAt: time.Date(2026, time.April, 24, 21, 0, 0, 0, time.UTC),
	}

	result, err := runExecutionPreflight(context.Background(), request)
	if err != nil {
		t.Fatalf("runExecutionPreflight: %v", err)
	}
	if result.PreflightArtifactPath == "" {
		t.Fatal("expected preflight artifact path")
	}
	if !result.DocPresence["TASK.md"] || !result.DocPresence["PLAN.md"] || !result.DocPresence["HANDOFF.md"] || !result.DocPresence["TASK-STATE.json"] {
		t.Fatalf("doc presence = %#v", result.DocPresence)
	}
	raw, err := os.ReadFile(result.PreflightArtifactPath)
	if err != nil {
		t.Fatalf("read preflight artifact: %v", err)
	}
	var artifact executionPreflightArtifact
	if err := json.Unmarshal(raw, &artifact); err != nil {
		t.Fatalf("decode preflight artifact: %v", err)
	}
	if artifact.OwnedTaskRoot != filepath.Join(ownedRoot, "Tracking", "Task-0008") {
		t.Fatalf("owned task root = %q", artifact.OwnedTaskRoot)
	}
	if artifact.CurrentCommit != request.RepoLane.CurrentCommit {
		t.Fatalf("current commit = %q, want %q", artifact.CurrentCommit, request.RepoLane.CurrentCommit)
	}
}

func TestRunWorkloadStepWritesOwnedLaneExecutionPacket(t *testing.T) {
	worktreeRoot := t.TempDir()
	writeTaskFile(t, filepath.Join(worktreeRoot, "Tracking", "Task-0008", "TASK.md"), "# task\n")
	writeTaskFile(t, filepath.Join(worktreeRoot, "Tracking", "Task-0008", "PLAN.md"), "# plan\n")
	writeTaskFile(t, filepath.Join(worktreeRoot, "Tracking", "Task-0008", "HANDOFF.md"), "# handoff\n")
	writeTaskFile(t, filepath.Join(worktreeRoot, "Tracking", "Task-0008", "TASK-STATE.json"), "{\"task_id\":\"Task-0008\"}\n")
	runTaskExecCommand(t, worktreeRoot, "git", "init")
	runTaskExecCommand(t, worktreeRoot, "git", "config", "user.email", "taskexec-tests@example.com")
	runTaskExecCommand(t, worktreeRoot, "git", "config", "user.name", "TaskExec Tests")
	runTaskExecCommand(t, worktreeRoot, "git", "add", ".")
	runTaskExecCommand(t, worktreeRoot, "git", "commit", "-m", "initial")

	ownedRoot := filepath.Join(t.TempDir(), "owned")
	runTaskExecCommand(t, worktreeRoot, "git", "worktree", "add", "--detach", ownedRoot, "HEAD")

	runArtifactRoot := filepath.Join(t.TempDir(), "artifacts")
	headCommit := stringsTrim(runTaskExecOutput(t, worktreeRoot, "git", "rev-parse", "HEAD"))
	request := taskrun.StartTaskRunRequest{
		RunID:          "taskrun--Task-0008--active",
		TaskID:         "Task-0008",
		MeaningSummary: "Create the durable backend task-run contract.",
		CapturedTaskSnapshot: taskrun.TaskDefinitionSnapshot{
			DeclaredWorktreeRoot: worktreeRoot,
			DeclaredTaskRoot:     filepath.Join(worktreeRoot, "Tracking", "Task-0008"),
			DeclaredTaskRevision: "revision-1",
			DeclaredGitRevision:  headCommit,
			CapturedAt:           time.Date(2026, time.April, 24, 21, 0, 0, 0, time.UTC),
		},
		RepoLane: taskrun.RepoLane{
			OwnedRepoRoot:         ownedRoot,
			CheckoutMode:          "git_worktree_detached",
			BaselineCommit:        headCommit,
			CurrentCommit:         headCommit,
			ApprovedRestoreCommit: headCommit,
			RunArtifactRoot:       runArtifactRoot,
			BootstrapArtifactPath: filepath.Join(runArtifactRoot, "owned-lane-bootstrap.json"),
			PreflightArtifactPath: filepath.Join(runArtifactRoot, "execution-preflight.json"),
		},
		DispatchRequestedAt: time.Date(2026, time.April, 24, 21, 0, 0, 0, time.UTC),
	}
	writeTaskFile(t, request.RepoLane.PreflightArtifactPath, "{}\n")

	result, err := runWorkloadStep(context.Background(), request, request.RepoLane)
	if err != nil {
		t.Fatalf("runWorkloadStep: %v", err)
	}
	if result.WorkloadStepPath == "" {
		t.Fatal("expected workload step path")
	}
	if !strings.Contains(result.WorkloadStepPath, ".codex-taskrun") {
		t.Fatalf("workload step path = %q", result.WorkloadStepPath)
	}
	raw, err := os.ReadFile(result.WorkloadStepPath)
	if err != nil {
		t.Fatalf("read workload step artifact: %v", err)
	}
	var artifact workloadStepArtifact
	if err := json.Unmarshal(raw, &artifact); err != nil {
		t.Fatalf("decode workload step artifact: %v", err)
	}
	if artifact.WorkloadInstruction == "" {
		t.Fatal("expected workload instruction")
	}
	if artifact.CurrentCommit != headCommit {
		t.Fatalf("current commit = %q, want %q", artifact.CurrentCommit, headCommit)
	}
	if artifact.PreflightArtifactPath != request.RepoLane.PreflightArtifactPath {
		t.Fatalf("preflight artifact path = %q, want %q", artifact.PreflightArtifactPath, request.RepoLane.PreflightArtifactPath)
	}
	if artifact.ExecutionKind != "task_0008_backend_validation" {
		t.Fatalf("execution kind = %q", artifact.ExecutionKind)
	}
	wantWorkingDir := filepath.Join(ownedRoot, "backend", "orchestration")
	if artifact.ExecutionWorkingDir != wantWorkingDir {
		t.Fatalf("execution working dir = %q, want %q", artifact.ExecutionWorkingDir, wantWorkingDir)
	}
	if len(artifact.ExecutionCommand) != 4 || artifact.ExecutionCommand[0] != "go" {
		t.Fatalf("execution command = %#v", artifact.ExecutionCommand)
	}
}

func TestRunExecuteWorkloadStepWritesResultFromPreparedPacket(t *testing.T) {
	ownedRoot := filepath.Join(t.TempDir(), "owned")
	if err := os.MkdirAll(filepath.Join(ownedRoot, ".codex-taskrun", "taskrun--Task-0008--active"), 0o755); err != nil {
		t.Fatalf("mkdir owned root: %v", err)
	}
	runTaskExecCommand(t, ownedRoot, "git", "init")
	runTaskExecCommand(t, ownedRoot, "git", "config", "user.email", "taskexec-tests@example.com")
	runTaskExecCommand(t, ownedRoot, "git", "config", "user.name", "TaskExec Tests")
	writeTaskFile(t, filepath.Join(ownedRoot, "README.txt"), "owned lane\n")
	runTaskExecCommand(t, ownedRoot, "git", "add", ".")
	runTaskExecCommand(t, ownedRoot, "git", "commit", "-m", "initial")

	stepPath := filepath.Join(ownedRoot, ".codex-taskrun", "taskrun--Task-0008--active", "workload-step-0001.json")
	step := workloadStepArtifact{
		TaskID:              "Task-0008",
		RunID:               "taskrun--Task-0008--active",
		OwnedRepoRoot:       ownedRoot,
		CurrentCommit:       stringsTrim(runTaskExecOutput(t, ownedRoot, "git", "rev-parse", "HEAD")),
		WorkloadInstruction: "Use the owned task root and captured task snapshot to execute the next backend-owned task step from inside this owned lane.",
		GeneratedAt:         time.Now().UTC(),
	}
	rawStep, err := json.Marshal(step)
	if err != nil {
		t.Fatalf("marshal step: %v", err)
	}
	if err := os.WriteFile(stepPath, append(rawStep, '\n'), 0o644); err != nil {
		t.Fatalf("write step: %v", err)
	}

	request := taskrun.StartTaskRunRequest{
		RunID:  "taskrun--Task-0008--active",
		TaskID: "Task-0008",
	}
	repoLane := taskrun.RepoLane{
		OwnedRepoRoot:    ownedRoot,
		WorkloadStepPath: stepPath,
	}

	result, err := runExecuteWorkloadStep(context.Background(), request, repoLane)
	if err != nil {
		t.Fatalf("runExecuteWorkloadStep: %v", err)
	}
	if result.WorkloadResultPath == "" {
		t.Fatal("expected workload result path")
	}
	rawResult, err := os.ReadFile(result.WorkloadResultPath)
	if err != nil {
		t.Fatalf("read workload result: %v", err)
	}
	var artifact workloadExecutionArtifact
	if err := json.Unmarshal(rawResult, &artifact); err != nil {
		t.Fatalf("decode workload result: %v", err)
	}
	if artifact.ExecutionSummary == "" {
		t.Fatal("expected execution summary")
	}
	if artifact.WorkloadStepPath != stepPath {
		t.Fatalf("workload step path = %q, want %q", artifact.WorkloadStepPath, stepPath)
	}
}

func TestRunExecuteWorkloadStepRunsTaskSpecificValidation(t *testing.T) {
	ownedRoot := t.TempDir()
	runTaskExecCommand(t, ownedRoot, "git", "init")
	runTaskExecCommand(t, ownedRoot, "git", "config", "user.email", "taskexec-tests@example.com")
	runTaskExecCommand(t, ownedRoot, "git", "config", "user.name", "TaskExec Tests")
	moduleRoot := filepath.Join(ownedRoot, "backend", "orchestration")
	ownedTaskRoot := filepath.Join(ownedRoot, "Tracking", "Task-0008")
	writeTaskFile(t, filepath.Join(moduleRoot, "go.mod"), "module example.com/task0008owned/backend/orchestration\n\ngo 1.25.0\n")
	writeTaskFile(t, filepath.Join(moduleRoot, "internal", "taskexec", "taskexec.go"), "package taskexec\n\nfunc Name() string { return \"taskexec\" }\n")
	writeTaskFile(t, filepath.Join(moduleRoot, "internal", "taskrun", "taskrun.go"), "package taskrun\n\nfunc Name() string { return \"taskrun\" }\n")
	writeTaskFile(t, filepath.Join(ownedTaskRoot, "TASK.md"), "## Summary\n\nTask-specific owned lane execution for Task-0008.\n")
	writeTaskFile(t, filepath.Join(ownedTaskRoot, "HANDOFF.md"), "## Next Recommended Step\n\n- Mutate one bounded backend-owned file in the owned lane.\n")
	writeTaskFile(t, filepath.Join(ownedTaskRoot, "CONSTRAINTS.md"), "## Active Constraints\n\n- Keep the slice bounded.\n")
	writeTaskFile(t, filepath.Join(ownedRoot, "README.txt"), "owned lane\n")
	runTaskExecCommand(t, ownedRoot, "git", "add", ".")
	runTaskExecCommand(t, ownedRoot, "git", "commit", "-m", "initial")

	stepRoot := filepath.Join(ownedRoot, ".codex-taskrun", "taskrun--Task-0008--active")
	if err := os.MkdirAll(stepRoot, 0o755); err != nil {
		t.Fatalf("mkdir step root: %v", err)
	}
	runArtifactRoot := filepath.Join(t.TempDir(), "artifacts")
	stepPath := filepath.Join(stepRoot, "workload-step-0001.json")
	step := workloadStepArtifact{
		TaskID:              "Task-0008",
		RunID:               "taskrun--Task-0008--active",
		OwnedRepoRoot:       ownedRoot,
		OwnedTaskRoot:       ownedTaskRoot,
		CurrentCommit:       stringsTrim(runTaskExecOutput(t, ownedRoot, "git", "rev-parse", "HEAD")),
		WorkloadInstruction: "Run focused Task-0008 backend validation from the owned checkout.",
		ExecutionKind:       "task_0008_backend_validation",
		ExecutionWorkingDir: moduleRoot,
		ExecutionCommand: []string{
			"go",
			"test",
			"./internal/taskexec",
			"./internal/taskrun",
		},
		GeneratedAt: time.Now().UTC(),
	}
	rawStep, err := json.Marshal(step)
	if err != nil {
		t.Fatalf("marshal step: %v", err)
	}
	if err := os.WriteFile(stepPath, append(rawStep, '\n'), 0o644); err != nil {
		t.Fatalf("write step: %v", err)
	}

	result, err := runExecuteWorkloadStep(context.Background(), taskrun.StartTaskRunRequest{
		RunID:  "taskrun--Task-0008--active",
		TaskID: "Task-0008",
	}, taskrun.RepoLane{
		OwnedRepoRoot:    ownedRoot,
		RunArtifactRoot:  runArtifactRoot,
		WorkloadStepPath: stepPath,
	})
	if err != nil {
		t.Fatalf("runExecuteWorkloadStep task-specific validation: %v", err)
	}
	rawResult, err := os.ReadFile(result.WorkloadResultPath)
	if err != nil {
		t.Fatalf("read workload result: %v", err)
	}
	var artifact workloadExecutionArtifact
	if err := json.Unmarshal(rawResult, &artifact); err != nil {
		t.Fatalf("decode workload result: %v", err)
	}
	if artifact.ExecutionKind != "task_0008_backend_validation" {
		t.Fatalf("execution kind = %q", artifact.ExecutionKind)
	}
	if artifact.ExecutionSummary != "Executed Task-0008 backend validation and wrote an owned-lane implementation brief." {
		t.Fatalf("execution summary = %q", artifact.ExecutionSummary)
	}
	if artifact.StdoutPath == "" || artifact.StderrPath == "" {
		t.Fatalf("expected stdout/stderr paths, got %#v", artifact)
	}
	if artifact.WorkloadOutputPath == "" {
		t.Fatalf("expected workload output path, got %#v", artifact)
	}
	if _, err := os.Stat(artifact.StdoutPath); err != nil {
		t.Fatalf("stat stdout path: %v", err)
	}
	if _, err := os.Stat(artifact.StderrPath); err != nil {
		t.Fatalf("stat stderr path: %v", err)
	}
	if _, err := os.Stat(artifact.WorkloadOutputPath); err != nil {
		t.Fatalf("stat workload output path: %v", err)
	}
	rawBrief, err := os.ReadFile(artifact.WorkloadOutputPath)
	if err != nil {
		t.Fatalf("read workload output path: %v", err)
	}
	if !strings.Contains(string(rawBrief), "Task-0008 Owned-Lane Implementation Brief") {
		t.Fatalf("brief contents = %q", string(rawBrief))
	}
	if !strings.Contains(artifact.GitStatusShortAfter, "OwnedLane") {
		t.Fatalf("git status after = %q", artifact.GitStatusShortAfter)
	}
	if artifact.ExitCode != 0 {
		t.Fatalf("exit code = %d", artifact.ExitCode)
	}
}

func writeTaskFile(t *testing.T, path string, contents string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("mkdir %s: %v", filepath.Dir(path), err)
	}
	if err := os.WriteFile(path, []byte(contents), 0o644); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
}

func runTaskExecCommand(t *testing.T, dir string, exe string, args ...string) {
	t.Helper()
	cmd := exec.Command(exe, args...)
	cmd.Dir = dir
	if output, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("%s %v failed: %v\n%s", exe, args, err, string(output))
	}
}

func runTaskExecOutput(t *testing.T, dir string, exe string, args ...string) string {
	t.Helper()
	cmd := exec.Command(exe, args...)
	cmd.Dir = dir
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("%s %v failed: %v\n%s", exe, args, err, string(output))
	}
	return string(output)
}

func stringsTrim(value string) string {
	return strings.TrimSpace(value)
}
