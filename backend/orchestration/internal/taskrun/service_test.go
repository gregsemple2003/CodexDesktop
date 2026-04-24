package taskrun

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"
)

type fakeRuntime struct {
	activeByTask map[string]TaskRunView
	byRunID      map[string]TaskRunView
	started      []StartTaskRunRequest
}

func newFakeRuntime() *fakeRuntime {
	return &fakeRuntime{
		activeByTask: map[string]TaskRunView{},
		byRunID:      map[string]TaskRunView{},
	}
}

func (f *fakeRuntime) StartTaskRun(_ context.Context, request StartTaskRunRequest) (TaskRunView, error) {
	f.started = append(f.started, request)
	run := TaskRunView{
		RunID:                  request.RunID,
		TaskID:                 request.TaskID,
		WorkflowID:             request.RunID,
		TemporalExecutionRunID: "temporal-run-id",
		Status:                 "active",
		StateEnvelope: StateEnvelope{
			State:             StateDispatching,
			ReasonCode:        "dispatch_started",
			StateSummary:      "Run is dispatching in an owned checkout.",
			NextOwner:         "backend",
			NextExpectedEvent: "Execution worker records the next task-run state update.",
			SuspiciousAfter:   request.DispatchRequestedAt.Add(15 * time.Minute),
		},
		MeaningSummary:             request.MeaningSummary,
		Attention:                  AttentionPriority{Level: AttentionWatch, Reason: "Run is active.", SortKey: "50-dispatching"},
		Actions:                    map[string]ActionAvailability{},
		RepoLane:                   request.RepoLane,
		LastProgressAt:             request.DispatchRequestedAt,
		LastProgressSummary:        "Captured task docs and provisioned an owned checkout.",
		CapturedTaskSnapshot:       request.CapturedTaskSnapshot,
		DocRuntimeDivergenceStatus: "in_sync",
	}
	f.activeByTask[request.TaskID] = run
	f.byRunID[request.RunID] = run
	return run, nil
}

func (f *fakeRuntime) GetTaskRun(_ context.Context, runID string) (TaskRunView, error) {
	run, ok := f.byRunID[runID]
	if !ok {
		return TaskRunView{}, ErrRunNotFound
	}
	return run, nil
}

func (f *fakeRuntime) GetActiveTaskRun(_ context.Context, taskID string) (TaskRunView, error) {
	run, ok := f.activeByTask[taskID]
	if !ok {
		return TaskRunView{}, ErrRunNotFound
	}
	return run, nil
}

func (f *fakeRuntime) ReconcileTaskSnapshot(_ context.Context, runID string, snapshot TaskDefinitionSnapshot) (TaskRunView, error) {
	run, ok := f.byRunID[runID]
	if !ok {
		return TaskRunView{}, ErrRunNotFound
	}
	run.CapturedTaskSnapshot = snapshot
	run.DocRuntimeDivergenceStatus = "reconciled"
	run.DocRuntimeDivergenceSummary = "Runtime captured newer task docs during task readback."
	f.byRunID[runID] = run
	f.activeByTask[run.TaskID] = run
	return run, nil
}

func (f *fakeRuntime) UpdateTaskRun(_ context.Context, runID string, update TaskRunUpdate) (TaskRunView, error) {
	run, ok := f.byRunID[runID]
	if !ok {
		return TaskRunView{}, ErrRunNotFound
	}
	if update.State != "" {
		run.StateEnvelope.State = update.State
	}
	if update.ReasonCode != "" {
		run.StateEnvelope.ReasonCode = update.ReasonCode
	}
	if update.StateSummary != "" {
		run.StateEnvelope.StateSummary = update.StateSummary
	}
	if update.NextOwner != "" {
		run.StateEnvelope.NextOwner = update.NextOwner
	}
	if update.NextExpectedEvent != "" {
		run.StateEnvelope.NextExpectedEvent = update.NextExpectedEvent
	}
	if !update.SuspiciousAfter.IsZero() {
		run.StateEnvelope.SuspiciousAfter = update.SuspiciousAfter
	}
	if update.WaitContract != nil {
		run.WaitContract = update.WaitContract
	}
	if update.Attention != nil {
		run.Attention = *update.Attention
	}
	if update.Actions != nil {
		run.Actions = update.Actions
		run.StateEnvelope.ActionBlockReasons = collectActionBlockReasons(update.Actions)
	}
	if update.RepoLane != nil {
		run.RepoLane = *update.RepoLane
	}
	if update.LastProgressSummary != "" {
		run.LastProgressSummary = update.LastProgressSummary
		run.LastProgressAt = time.Now().UTC()
	}
	if !update.CompletedAt.IsZero() {
		run.LastProgressAt = update.CompletedAt
	}
	switch run.StateEnvelope.State {
	case StateCompleted:
		run.Status = "completed"
	case StateFailed:
		run.Status = "failed"
	case StateInterrupted:
		run.Status = "interrupted"
	default:
		run.Status = "active"
	}
	f.byRunID[runID] = run
	f.activeByTask[run.TaskID] = run
	return run, nil
}

func TestListTasksParsesMeaningAndReadyState(t *testing.T) {
	worktreeRoot := writeTaskTrackingRoot(t, map[string]taskFixture{
		"Task-0008": {
			taskMD: `# Task 0008

## Title

Build the backend task dispatch layer.

## Summary

Create the durable backend task-run contract so later clients do not guess state.
`,
			taskState: `{
  "task_id": "Task-0008",
  "status": "in_progress",
  "phase": "implementation",
  "plan_approved": true,
  "current_pass": "PASS-0000",
  "current_gate": "implementation",
  "blockers": [],
  "updated_at": "2026-04-24T16:27:00-04:00"
}`,
			planMD:        "# approved plan\n",
			handoffMD:     "# handoff\n",
			constraintsMD: "# constraints\n",
		},
	})

	service := NewService(worktreeRoot, filepath.Join(worktreeRoot, ".runs"), nil)
	tasks, err := service.ListTasks(context.Background())
	if err != nil {
		t.Fatalf("list tasks: %v", err)
	}

	if len(tasks) != 1 {
		t.Fatalf("task count = %d, want 1", len(tasks))
	}
	task := tasks[0]
	if task.TaskID != "Task-0008" {
		t.Fatalf("task id = %q", task.TaskID)
	}
	if task.MeaningSummary != "Create the durable backend task-run contract so later clients do not guess state." {
		t.Fatalf("meaning summary = %q", task.MeaningSummary)
	}
	if task.StateEnvelope.State != StateReady {
		t.Fatalf("state = %q, want %q", task.StateEnvelope.State, StateReady)
	}
	if task.DispatchReadiness.Ready {
		t.Fatal("dispatch readiness should stay false until the durable dispatch lane exists")
	}
	if task.Actions[ActionDispatch].Allowed {
		t.Fatal("dispatch action should stay blocked until the durable dispatch lane exists")
	}
	if task.CurrentStory.Status != "no_active_run" {
		t.Fatalf("current story status = %q", task.CurrentStory.Status)
	}
}

func TestTaskUsesWaitingForHumanWhenPlanIsNotApproved(t *testing.T) {
	worktreeRoot := writeTaskTrackingRoot(t, map[string]taskFixture{
		"Task-0012": {
			taskMD: `# Task 0012

## Title

Ship the next task.

## Summary

Need approval before implementation starts.
`,
			taskState: `{
  "task_id": "Task-0012",
  "status": "in_progress",
  "phase": "planning",
  "plan_approved": false,
  "current_pass": "",
  "current_gate": "planning",
  "blockers": [],
  "updated_at": "2026-04-24T16:27:00-04:00"
}`,
			planMD: "# plan\n",
		},
	})

	service := NewService(worktreeRoot, filepath.Join(worktreeRoot, ".runs"), nil)
	task, err := service.Task(context.Background(), "Task-0012")
	if err != nil {
		t.Fatalf("task detail: %v", err)
	}

	if task.StateEnvelope.State != StateWaitingForHuman {
		t.Fatalf("state = %q, want %q", task.StateEnvelope.State, StateWaitingForHuman)
	}
	if task.StateEnvelope.ReasonCode != "plan_approval_required" {
		t.Fatalf("reason code = %q", task.StateEnvelope.ReasonCode)
	}
	if task.Actions[ActionDispatch].Allowed {
		t.Fatal("dispatch should not be allowed while plan approval is missing")
	}
}

func TestDispatchProvisionOwnedLaneAndCaptureBaselineCommit(t *testing.T) {
	worktreeRoot := writeGitTaskTrackingRoot(t, map[string]taskFixture{
		"Task-0008": {
			taskMD: `# Task 0008

## Title

Build the backend task dispatch layer.

## Summary

Create the durable backend task-run contract so later clients do not guess state.
`,
			taskState: `{
  "task_id": "Task-0008",
  "status": "in_progress",
  "phase": "implementation",
  "plan_approved": true,
  "current_pass": "PASS-0001",
  "current_gate": "implementation",
  "blockers": [],
  "updated_at": "2026-04-24T16:44:31-04:00"
}`,
			planMD: "# approved plan\n",
		},
	})

	runtime := newFakeRuntime()
	runsRoot := filepath.Join(worktreeRoot, ".runs")
	service := NewService(worktreeRoot, runsRoot, runtime)

	run, err := service.Dispatch(context.Background(), "Task-0008")
	if err != nil {
		t.Fatalf("dispatch: %v", err)
	}

	if len(runtime.started) != 1 {
		t.Fatalf("started requests = %d, want 1", len(runtime.started))
	}
	request := runtime.started[0]
	if request.TaskID != "Task-0008" {
		t.Fatalf("task id = %q", request.TaskID)
	}
	if request.RepoLane.BaselineCommit == "" {
		t.Fatal("expected baseline commit to be captured")
	}
	if request.RepoLane.ApprovedRestoreCommit != request.RepoLane.BaselineCommit {
		t.Fatalf("approved restore commit = %q, want baseline %q", request.RepoLane.ApprovedRestoreCommit, request.RepoLane.BaselineCommit)
	}
	if request.RepoLane.OwnedRepoRoot == "" {
		t.Fatal("expected owned repo root to be set")
	}
	if _, err := os.Stat(request.RepoLane.OwnedRepoRoot); err != nil {
		t.Fatalf("owned repo root missing: %v", err)
	}
	if run.RunID != ActiveRunID("Task-0008") {
		t.Fatalf("run id = %q", run.RunID)
	}
}

func TestUpdateRunAppliesRicherStateContract(t *testing.T) {
	worktreeRoot := writeGitTaskTrackingRoot(t, map[string]taskFixture{
		"Task-0008": {
			taskMD: `# Task 0008

## Title

Build the backend task dispatch layer.

## Summary

Create the durable backend task-run contract so later clients do not guess state.
`,
			taskState: `{
  "task_id": "Task-0008",
  "status": "in_progress",
  "phase": "implementation",
  "plan_approved": true,
  "current_pass": "PASS-0001",
  "current_gate": "implementation",
  "blockers": [],
  "updated_at": "2026-04-24T16:44:31-04:00"
}`,
			planMD: "# approved plan\n",
		},
	})

	runtime := newFakeRuntime()
	runsRoot := filepath.Join(worktreeRoot, ".runs")
	service := NewService(worktreeRoot, runsRoot, runtime)
	run, err := service.Dispatch(context.Background(), "Task-0008")
	if err != nil {
		t.Fatalf("dispatch: %v", err)
	}

	updated, err := service.UpdateRun(context.Background(), run.RunID, TaskRunUpdate{
		State:               StateWaitingForHuman,
		ReasonCode:          "review_required",
		StateSummary:        "Run is waiting for human review.",
		NextOwner:           "human",
		NextExpectedEvent:   "Approve the next backend action.",
		LastProgressSummary: "Run recorded a review checkpoint.",
		WaitContract: &WaitContract{
			WaitingOn:           "human_review",
			WhyBlocked:          "The next backend action needs human approval.",
			ResumeWhen:          "The human approves the next backend action.",
			HumanActionRequired: true,
			HumanActionTarget: &HumanActionTarget{
				Kind:  "approval_action",
				Label: "Approve backend review step",
				URI:   "approval://taskrun/Task-0008",
			},
		},
	})
	if err != nil {
		t.Fatalf("update run: %v", err)
	}

	if updated.StateEnvelope.State != StateWaitingForHuman {
		t.Fatalf("updated state = %q", updated.StateEnvelope.State)
	}
	if updated.WaitContract == nil || updated.WaitContract.HumanActionTarget == nil {
		t.Fatal("expected wait contract with explicit human action target")
	}
	if updated.Attention.Level != AttentionNeedsAttention {
		t.Fatalf("attention level = %q, want %q", updated.Attention.Level, AttentionNeedsAttention)
	}
}

func TestRunReadSupervisesStaleProgressIntoSleepingState(t *testing.T) {
	worktreeRoot := writeGitTaskTrackingRoot(t, map[string]taskFixture{
		"Task-0008": {
			taskMD: `# Task 0008

## Title

Build the backend task dispatch layer.

## Summary

Create the durable backend task-run contract so later clients do not guess state.
`,
			taskState: `{
  "task_id": "Task-0008",
  "status": "in_progress",
  "phase": "implementation",
  "plan_approved": true,
  "current_pass": "PASS-0002",
  "current_gate": "implementation",
  "blockers": [],
  "updated_at": "2026-04-24T17:10:00-04:00"
}`,
			planMD: "# approved plan\n",
		},
	})

	runtime := newFakeRuntime()
	service := NewService(worktreeRoot, filepath.Join(worktreeRoot, ".runs"), runtime)
	run, err := service.Dispatch(context.Background(), "Task-0008")
	if err != nil {
		t.Fatalf("dispatch: %v", err)
	}

	if _, err := service.UpdateRun(context.Background(), run.RunID, TaskRunUpdate{
		State:             StateRunning,
		ReasonCode:        "worker_started",
		StateSummary:      "Run is actively executing.",
		SuspiciousAfter:   time.Now().UTC().Add(-1 * time.Minute),
		NextOwner:         "backend",
		NextExpectedEvent: "Execution worker records the next progress checkpoint.",
	}); err != nil {
		t.Fatalf("prime run: %v", err)
	}

	supervised, err := service.Run(context.Background(), run.RunID)
	if err != nil {
		t.Fatalf("run detail: %v", err)
	}
	if supervised.StateEnvelope.State != StateSleepingOrStalled {
		t.Fatalf("state = %q, want %q", supervised.StateEnvelope.State, StateSleepingOrStalled)
	}
	if !supervised.Actions[ActionPoke].Allowed {
		t.Fatal("poke should be allowed after supervision marks the run stalled")
	}
	if !supervised.Actions[ActionInterrupt].Allowed {
		t.Fatal("interrupt should stay allowed for a stalled run")
	}
}

func TestInterruptRunRestoresOwnedLaneAndMarksInterrupted(t *testing.T) {
	worktreeRoot := writeGitTaskTrackingRoot(t, map[string]taskFixture{
		"Task-0008": {
			taskMD: `# Task 0008

## Title

Build the backend task dispatch layer.

## Summary

Create the durable backend task-run contract so later clients do not guess state.
`,
			taskState: `{
  "task_id": "Task-0008",
  "status": "in_progress",
  "phase": "implementation",
  "plan_approved": true,
  "current_pass": "PASS-0002",
  "current_gate": "implementation",
  "blockers": [],
  "updated_at": "2026-04-24T17:10:00-04:00"
}`,
			planMD: "# approved plan\n",
		},
	})

	runtime := newFakeRuntime()
	service := NewService(worktreeRoot, filepath.Join(worktreeRoot, ".runs"), runtime)
	run, err := service.Dispatch(context.Background(), "Task-0008")
	if err != nil {
		t.Fatalf("dispatch: %v", err)
	}

	writeFile(t, filepath.Join(run.RepoLane.OwnedRepoRoot, "scratch.txt"), "temporary\n")
	writeFile(t, filepath.Join(run.RepoLane.OwnedRepoRoot, "Tracking", "Task-0008", "TASK.md"), "# changed\n")

	interrupted, err := service.InterruptRun(context.Background(), run.RunID)
	if err != nil {
		t.Fatalf("interrupt run: %v", err)
	}
	if interrupted.StateEnvelope.State != StateInterrupted {
		t.Fatalf("state = %q, want %q", interrupted.StateEnvelope.State, StateInterrupted)
	}
	if interrupted.RepoLane.ResetStatus != "restored" {
		t.Fatalf("reset status = %q, want restored", interrupted.RepoLane.ResetStatus)
	}
	if interrupted.Status != "interrupted" {
		t.Fatalf("status = %q, want interrupted", interrupted.Status)
	}

	cmd := exec.Command("git", "-C", interrupted.RepoLane.OwnedRepoRoot, "status", "--short")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("git status in owned lane: %v\n%s", err, string(output))
	}
	if got := string(output); got != "" {
		t.Fatalf("owned lane should be clean after interrupt reset, got %q", got)
	}

	task, err := service.Task(context.Background(), "Task-0008")
	if err != nil {
		t.Fatalf("task detail after interrupt: %v", err)
	}
	if task.CurrentStory.Status != "no_active_run" {
		t.Fatalf("current story after interrupt = %q, want no_active_run", task.CurrentStory.Status)
	}
	if !task.DispatchReadiness.Ready {
		t.Fatal("dispatch should be ready again after the latest run is terminal")
	}
}

type taskFixture struct {
	taskMD        string
	taskState     string
	planMD        string
	handoffMD     string
	constraintsMD string
}

func writeTaskTrackingRoot(t *testing.T, tasks map[string]taskFixture) string {
	t.Helper()
	worktreeRoot := t.TempDir()
	trackingRoot := filepath.Join(worktreeRoot, "Tracking")
	if err := os.MkdirAll(trackingRoot, 0o755); err != nil {
		t.Fatalf("mkdir tracking root: %v", err)
	}

	for taskID, fixture := range tasks {
		taskRoot := filepath.Join(trackingRoot, taskID)
		if err := os.MkdirAll(taskRoot, 0o755); err != nil {
			t.Fatalf("mkdir task root: %v", err)
		}
		writeFile(t, filepath.Join(taskRoot, "TASK.md"), fixture.taskMD)
		writeFile(t, filepath.Join(taskRoot, "TASK-STATE.json"), fixture.taskState)
		if fixture.planMD != "" {
			writeFile(t, filepath.Join(taskRoot, "PLAN.md"), fixture.planMD)
		}
		if fixture.handoffMD != "" {
			writeFile(t, filepath.Join(taskRoot, "HANDOFF.md"), fixture.handoffMD)
		}
		if fixture.constraintsMD != "" {
			writeFile(t, filepath.Join(taskRoot, "CONSTRAINTS.md"), fixture.constraintsMD)
		}
	}

	return worktreeRoot
}

func writeGitTaskTrackingRoot(t *testing.T, tasks map[string]taskFixture) string {
	t.Helper()
	worktreeRoot := writeTaskTrackingRoot(t, tasks)
	runCommand(t, worktreeRoot, "git", "init")
	runCommand(t, worktreeRoot, "git", "config", "user.email", "taskrun-tests@example.com")
	runCommand(t, worktreeRoot, "git", "config", "user.name", "TaskRun Tests")
	runCommand(t, worktreeRoot, "git", "add", ".")
	runCommand(t, worktreeRoot, "git", "commit", "-m", "initial task fixtures")
	return worktreeRoot
}

func writeFile(t *testing.T, path string, contents string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(contents), 0o644); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
}

func runCommand(t *testing.T, dir string, exe string, args ...string) {
	t.Helper()
	cmd := exec.Command(exe, args...)
	cmd.Dir = dir
	if output, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("%s %v failed: %v\n%s", exe, args, err, string(output))
	}
}
