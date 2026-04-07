package temporalbackend

import (
	"testing"

	"github.com/gregsemple2003/CodexDesktop/backend/orchestration/internal/controlplane"
	"github.com/gregsemple2003/CodexDesktop/backend/orchestration/internal/jobs"
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
