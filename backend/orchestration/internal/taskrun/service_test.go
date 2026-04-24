package taskrun

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

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

	service := NewService(worktreeRoot)
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

	service := NewService(worktreeRoot)
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

func writeFile(t *testing.T, path string, contents string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(contents), 0o644); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
}
