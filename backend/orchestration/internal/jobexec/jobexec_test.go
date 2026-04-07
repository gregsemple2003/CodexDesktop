package jobexec

import (
	"strings"
	"testing"
	"time"

	"github.com/gregsemple2003/CodexDesktop/backend/orchestration/internal/config"
	"github.com/gregsemple2003/CodexDesktop/backend/orchestration/internal/controlplane"
	"github.com/gregsemple2003/CodexDesktop/backend/orchestration/internal/jobs"
)

func TestBuildCommandPlanUsesCodexExecAndArtifacts(t *testing.T) {
	request := controlplane.JobRunRequest{
		JobID:           "codex-daily-agentic-swe-digest",
		TriggerType:     jobs.TriggerTypeManual,
		DesiredSpecHash: "spec-hash-123",
		RequestedAt:     time.Date(2026, time.April, 6, 22, 0, 0, 0, time.UTC),
		WorkflowID:      "workflow-id",
		RunID:           "run-id",
		Spec: jobs.Spec{
			JobID: "codex-daily-agentic-swe-digest",
			Executor: jobs.Executor{
				Type:       jobs.ExecutorTypeCodexExec,
				Cwd:        `C:\Users\gregs\.codex`,
				Entrypoint: "agentic-swe-digest",
				Args:       []string{"--days", "1", "--email"},
			},
			Runtime: jobs.RuntimeConfig{
				WorkflowType: "codex.exec.job",
				TaskQueue:    "codex-orchestration",
			},
		},
	}

	plan, err := BuildCommandPlan(config.Config{
		CodexExecutable: `C:\tools\codex.exe`,
		RunsRoot:        `C:\temp\codex-runs`,
	}, request)
	if err != nil {
		t.Fatalf("build command plan: %v", err)
	}

	if plan.Executable != `C:\tools\codex.exe` {
		t.Fatalf("executable = %q", plan.Executable)
	}
	if len(plan.Args) < 8 || plan.Args[0] != "exec" {
		t.Fatalf("unexpected command args: %v", plan.Args)
	}
	if !strings.Contains(strings.Join(plan.Args, " "), "--full-auto") {
		t.Fatalf("expected --full-auto in command args: %v", plan.Args)
	}
	if !strings.Contains(strings.Join(plan.Args, " "), "agentic-swe-digest") {
		t.Fatalf("expected prompt to mention entrypoint: %v", plan.Args)
	}
	if !strings.Contains(plan.EventLogPath, "workflow-id") || !strings.Contains(plan.FinalMessagePath, "workflow-id") {
		t.Fatalf("expected workflow id in artifact paths: %+v", plan)
	}
}
