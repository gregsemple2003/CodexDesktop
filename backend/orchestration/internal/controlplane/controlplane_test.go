package controlplane

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/gregsemple2003/CodexDesktop/backend/orchestration/internal/jobs"
)

type fakeBackend struct {
	schedules map[string]RuntimeSchedule
	started   []JobRunRequest
}

func newFakeBackend() *fakeBackend {
	return &fakeBackend{schedules: map[string]RuntimeSchedule{}}
}

func (b *fakeBackend) ListManagedSchedules(context.Context) ([]RuntimeSchedule, error) {
	schedules := make([]RuntimeSchedule, 0, len(b.schedules))
	for _, schedule := range b.schedules {
		schedules = append(schedules, schedule)
	}
	return schedules, nil
}

func (b *fakeBackend) CreateSchedule(_ context.Context, desired DesiredSchedule) error {
	b.schedules[desired.ScheduleID] = runtimeFromDesired(desired)
	return nil
}

func (b *fakeBackend) UpdateSchedule(_ context.Context, desired DesiredSchedule) error {
	b.schedules[desired.ScheduleID] = runtimeFromDesired(desired)
	return nil
}

func (b *fakeBackend) DeleteSchedule(_ context.Context, scheduleID string) error {
	delete(b.schedules, scheduleID)
	return nil
}

func (b *fakeBackend) Close() error {
	return nil
}

func (b *fakeBackend) StartJobRun(_ context.Context, request JobRunRequest) (StartedRun, error) {
	b.started = append(b.started, request)
	return StartedRun{
		JobID:           request.JobID,
		TriggerType:     request.TriggerType,
		TriggerPath:     request.TriggerPath,
		DesiredSpecHash: request.DesiredSpecHash,
		RequestedAt:     request.RequestedAt,
		WorkflowID:      "workflow-id",
		RunID:           "run-id",
	}, nil
}

func TestReconcileCreatesSchedulesAndReportsInSync(t *testing.T) {
	root := writeJobsRoot(t, []jobs.Spec{
		{
			APIVersion:   jobs.APIVersion,
			JobID:        "codex-daily-agentic-swe-digest",
			Label:        "Agentic SWE Digest",
			Description:  "Daily digest",
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
	})

	service := NewService(root, newFakeBackend())
	report, err := service.Reconcile(context.Background())
	if err != nil {
		t.Fatalf("reconcile: %v", err)
	}

	if len(report.Created) != 1 {
		t.Fatalf("created schedules = %d, want 1", len(report.Created))
	}
	if len(report.State.Jobs) != 1 {
		t.Fatalf("jobs = %d, want 1", len(report.State.Jobs))
	}
	job := report.State.Jobs[0]
	if job.Status != "in_sync" {
		t.Fatalf("job status = %q, want in_sync", job.Status)
	}
	if len(job.Schedules) != 1 {
		t.Fatalf("schedule count = %d, want 1", len(job.Schedules))
	}
	if job.Schedules[0].Status != "in_sync" {
		t.Fatalf("schedule status = %q, want in_sync", job.Schedules[0].Status)
	}
	if report.State.LastSync.LastSuccessAt.IsZero() {
		t.Fatal("expected last sync success timestamp to be recorded")
	}
}

func TestCompileJobUsesScheduleTriggerOrderForManagedIDs(t *testing.T) {
	spec := jobs.Spec{
		APIVersion:   jobs.APIVersion,
		JobID:        "codex-daily-physical-agents-digest",
		Label:        "Physical Agents Digest",
		Description:  "Daily digest",
		DesiredState: jobs.DesiredStateEnabled,
		Triggers: []jobs.Trigger{
			{Type: jobs.TriggerTypeManual},
			{Type: jobs.TriggerTypeSchedule, Cron: "0 4 * * *", Timezone: "America/Toronto"},
			{Type: jobs.TriggerTypeWebhook, Path: "digests/physical-agents"},
			{Type: jobs.TriggerTypeSchedule, Cron: "0 5 * * *", Timezone: "America/Toronto"},
		},
		Executor: jobs.Executor{
			Type:       jobs.ExecutorTypeCodexExec,
			Cwd:        `C:\Users\gregs\.codex`,
			Entrypoint: "physical-agents-digest",
		},
		Runtime: jobs.RuntimeConfig{
			WorkflowType: "codex.exec.job",
			TaskQueue:    "codex-orchestration",
		},
	}

	compiled := compileJob(spec)

	if len(compiled.Schedules) != 2 {
		t.Fatalf("schedule count = %d, want 2", len(compiled.Schedules))
	}
	if compiled.Schedules[0].ScheduleID != "codex-job--codex-daily-physical-agents-digest--00" {
		t.Fatalf("first schedule id = %q", compiled.Schedules[0].ScheduleID)
	}
	if compiled.Schedules[1].ScheduleID != "codex-job--codex-daily-physical-agents-digest--01" {
		t.Fatalf("second schedule id = %q", compiled.Schedules[1].ScheduleID)
	}
	if compiled.Schedules[0].TriggerIndex != 0 || compiled.Schedules[1].TriggerIndex != 1 {
		t.Fatalf("unexpected trigger indices: %+v", compiled.Schedules)
	}
}

func TestSnapshotFlagsDisabledJobWithRuntimeSchedule(t *testing.T) {
	spec := jobs.Spec{
		APIVersion:   jobs.APIVersion,
		JobID:        "codex-daily-physical-agents-digest",
		Label:        "Physical Agents Digest",
		Description:  "Daily digest",
		DesiredState: jobs.DesiredStateDisabled,
		Triggers: []jobs.Trigger{
			{Type: jobs.TriggerTypeSchedule, Cron: "0 4 * * *", Timezone: "America/Toronto"},
			{Type: jobs.TriggerTypeWebhook, Path: "digests/physical-agents"},
		},
		Executor: jobs.Executor{
			Type:       jobs.ExecutorTypeCodexExec,
			Cwd:        `C:\Users\gregs\.codex`,
			Entrypoint: "physical-agents-digest",
		},
		Runtime: jobs.RuntimeConfig{
			WorkflowType: "codex.exec.job",
			TaskQueue:    "codex-orchestration",
		},
	}
	root := writeJobsRoot(t, []jobs.Spec{spec})

	compiled := compileJob(spec)
	backend := newFakeBackend()
	backend.schedules[compiled.Schedules[0].ScheduleID] = runtimeFromDesired(compiled.Schedules[0])

	service := NewService(root, backend)
	state, err := service.Snapshot(context.Background())
	if err != nil {
		t.Fatalf("snapshot: %v", err)
	}

	if len(state.Jobs) != 1 {
		t.Fatalf("jobs = %d, want 1", len(state.Jobs))
	}
	job := state.Jobs[0]
	if job.Status != "drifted" {
		t.Fatalf("job status = %q, want drifted", job.Status)
	}
	if len(job.Schedules) != 1 {
		t.Fatalf("schedule count = %d, want 1", len(job.Schedules))
	}
	if job.Schedules[0].Status != "drifted" {
		t.Fatalf("schedule status = %q, want drifted", job.Schedules[0].Status)
	}
	if len(job.Schedules[0].Drift) == 0 || job.Schedules[0].Drift[0] != "unexpected_runtime_schedule" {
		t.Fatalf("schedule drift = %v, want unexpected_runtime_schedule", job.Schedules[0].Drift)
	}
}

func TestRunNowStartsManualWorkflowForManualJobs(t *testing.T) {
	root := writeJobsRoot(t, []jobs.Spec{
		{
			APIVersion:   jobs.APIVersion,
			JobID:        "codex-daily-agentic-swe-digest",
			Label:        "Agentic SWE Digest",
			Description:  "Daily digest",
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
	})

	backend := newFakeBackend()
	service := NewService(root, backend)
	started, err := service.RunNow(context.Background(), "codex-daily-agentic-swe-digest")
	if err != nil {
		t.Fatalf("run now: %v", err)
	}

	if started.TriggerType != jobs.TriggerTypeManual {
		t.Fatalf("trigger type = %q, want manual", started.TriggerType)
	}
	if len(backend.started) != 1 || backend.started[0].TriggerType != jobs.TriggerTypeManual {
		t.Fatalf("unexpected started requests: %+v", backend.started)
	}
	if backend.started[0].DesiredSpecHash == "" {
		t.Fatal("expected desired spec hash to be populated")
	}
}

func TestTriggerWebhookStartsWebhookWorkflowForMatchingPath(t *testing.T) {
	root := writeJobsRoot(t, []jobs.Spec{
		{
			APIVersion:   jobs.APIVersion,
			JobID:        "codex-daily-physical-agents-digest",
			Label:        "Physical Agents Digest",
			Description:  "Daily digest",
			DesiredState: jobs.DesiredStateEnabled,
			Triggers: []jobs.Trigger{
				{Type: jobs.TriggerTypeWebhook, Path: "digests/physical-agents"},
			},
			Executor: jobs.Executor{
				Type:       jobs.ExecutorTypeCodexExec,
				Cwd:        `C:\Users\gregs\.codex`,
				Entrypoint: "physical-agents-digest",
			},
			Runtime: jobs.RuntimeConfig{
				WorkflowType: "codex.exec.job",
				TaskQueue:    "codex-orchestration",
			},
		},
	})

	backend := newFakeBackend()
	service := NewService(root, backend)
	started, err := service.TriggerWebhook(context.Background(), "digests/physical-agents")
	if err != nil {
		t.Fatalf("trigger webhook: %v", err)
	}

	if started.TriggerType != jobs.TriggerTypeWebhook {
		t.Fatalf("trigger type = %q, want webhook", started.TriggerType)
	}
	if len(backend.started) != 1 || backend.started[0].TriggerPath != "digests/physical-agents" {
		t.Fatalf("unexpected started requests: %+v", backend.started)
	}
}

func runtimeFromDesired(desired DesiredSchedule) RuntimeSchedule {
	return RuntimeSchedule{
		ScheduleID:      desired.ScheduleID,
		JobID:           desired.JobID,
		TriggerIndex:    desired.TriggerIndex,
		CronExpressions: []string{desired.Cron},
		TimeZoneName:    desired.Timezone,
		ManagedCron:     desired.Cron,
		ManagedTimezone: desired.Timezone,
		ManagedSpecHash: desired.SpecHash,
		WorkflowType:    desired.WorkflowType,
		TaskQueue:       desired.TaskQueue,
		RecentRuns: []RunRecord{
			{
				ScheduleID:   desired.ScheduleID,
				ScheduleTime: time.Date(2026, time.April, 6, 4, 0, 0, 0, time.UTC),
				ActualTime:   time.Date(2026, time.April, 6, 4, 0, 1, 0, time.UTC),
				WorkflowID:   desired.WorkflowID,
			},
		},
	}
}

func writeJobsRoot(t *testing.T, specs []jobs.Spec) string {
	t.Helper()
	root := t.TempDir()
	specsDir := filepath.Join(root, "specs")
	if err := os.MkdirAll(specsDir, 0o755); err != nil {
		t.Fatalf("mkdir specs: %v", err)
	}
	for _, spec := range specs {
		path := filepath.Join(specsDir, spec.JobID+".json")
		raw, err := json.Marshal(spec)
		if err != nil {
			t.Fatalf("marshal spec: %v", err)
		}
		if err := os.WriteFile(path, raw, 0o644); err != nil {
			t.Fatalf("write spec %s: %v", path, err)
		}
	}
	return root
}
