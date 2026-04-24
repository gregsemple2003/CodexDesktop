package temporalbackend

import (
	"errors"
	"testing"

	"github.com/gregsemple2003/CodexDesktop/backend/orchestration/internal/controlplane"
	"github.com/gregsemple2003/CodexDesktop/backend/orchestration/internal/jobs"
	"github.com/gregsemple2003/CodexDesktop/backend/orchestration/internal/taskrun"
)

func TestBuildScheduleOptionsUseManagedCatchupWindow(t *testing.T) {
	desired := testDesiredSchedule()

	options := buildScheduleOptions(desired)

	if options.CatchupWindow != controlplane.ManagedScheduleCatchupWindow {
		t.Fatalf("catchup window = %v, want %v", options.CatchupWindow, controlplane.ManagedScheduleCatchupWindow)
	}
}

func TestBuildScheduleUseManagedCatchupWindow(t *testing.T) {
	desired := testDesiredSchedule()

	schedule := buildSchedule(desired)

	if schedule.Policy == nil {
		t.Fatal("expected schedule policy")
	}
	if schedule.Policy.CatchupWindow != controlplane.ManagedScheduleCatchupWindow {
		t.Fatalf("catchup window = %v, want %v", schedule.Policy.CatchupWindow, controlplane.ManagedScheduleCatchupWindow)
	}
}

func TestReadUpdatedTaskRunReturnsOpenViewWhenQuerySucceeds(t *testing.T) {
	want := taskrun.TaskRunView{RunID: "taskrun--Task-0008--active", Status: "active"}

	got, err := readUpdatedTaskRun(func() (taskrun.TaskRunView, error) {
		return want, nil
	}, func() (taskrun.TaskRunView, error) {
		t.Fatal("closed fallback should not run when current query succeeds")
		return taskrun.TaskRunView{}, nil
	})
	if err != nil {
		t.Fatalf("readUpdatedTaskRun returned error: %v", err)
	}
	if got.RunID != want.RunID || got.Status != want.Status {
		t.Fatalf("got %#v, want %#v", got, want)
	}
}

func TestReadUpdatedTaskRunFallsBackToClosedResultOnNotFound(t *testing.T) {
	want := taskrun.TaskRunView{
		RunID:      "taskrun--Task-0008--active",
		Status:     "interrupted",
		Resolution: &taskrun.RunResolution{Kind: "interrupt_review", Decision: "redispatch_ready"},
	}

	got, err := readUpdatedTaskRun(func() (taskrun.TaskRunView, error) {
		return taskrun.TaskRunView{}, taskrun.ErrRunNotFound
	}, func() (taskrun.TaskRunView, error) {
		return want, nil
	})
	if err != nil {
		t.Fatalf("readUpdatedTaskRun returned error: %v", err)
	}
	if got.RunID != want.RunID || got.Status != want.Status {
		t.Fatalf("got %#v, want %#v", got, want)
	}
	if got.Resolution == nil || got.Resolution.Decision != "redispatch_ready" {
		t.Fatalf("resolution = %#v", got.Resolution)
	}
}

func TestReadUpdatedTaskRunPreservesNonNotFoundError(t *testing.T) {
	wantErr := errors.New("query failed")

	_, err := readUpdatedTaskRun(func() (taskrun.TaskRunView, error) {
		return taskrun.TaskRunView{}, wantErr
	}, func() (taskrun.TaskRunView, error) {
		t.Fatal("closed fallback should not run on non-not-found errors")
		return taskrun.TaskRunView{}, nil
	})
	if !errors.Is(err, wantErr) {
		t.Fatalf("error = %v, want %v", err, wantErr)
	}
}

func testDesiredSchedule() controlplane.DesiredSchedule {
	return controlplane.DesiredSchedule{
		ScheduleID:    "codex-job--codex-daily-agentic-swe-digest--00",
		JobID:         "codex-daily-agentic-swe-digest",
		TriggerIndex:  0,
		Enabled:       true,
		Cron:          "0 4 * * *",
		Timezone:      "America/Toronto",
		CatchupWindow: controlplane.ManagedScheduleCatchupWindow,
		WorkflowType:  "codex.exec.job",
		TaskQueue:     "codex-orchestration",
		WorkflowID:    "codex-daily-agentic-swe-digest/schedule/00",
		SpecHash:      "spec-hash",
		Spec: jobs.Spec{
			APIVersion:   jobs.APIVersion,
			JobID:        "codex-daily-agentic-swe-digest",
			Label:        "Codex Daily Agentic SWE Digest",
			DesiredState: jobs.DesiredStateEnabled,
			Triggers: []jobs.Trigger{
				{Type: jobs.TriggerTypeSchedule, Cron: "0 4 * * *", Timezone: "America/Toronto"},
				{Type: jobs.TriggerTypeManual},
			},
			Executor: jobs.Executor{
				Type:       jobs.ExecutorTypeCodexExec,
				Cwd:        `C:\Users\gregs\.codex`,
				Entrypoint: "agentic-swe-digest",
			},
			Runtime: jobs.RuntimeConfig{
				WorkflowType: "codex.exec.job",
				TaskQueue:    "codex-orchestration",
			},
		},
	}
}
