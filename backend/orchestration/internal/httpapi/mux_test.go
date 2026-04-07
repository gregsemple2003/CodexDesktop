package httpapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/gregsemple2003/CodexDesktop/backend/orchestration/internal/config"
	"github.com/gregsemple2003/CodexDesktop/backend/orchestration/internal/controlplane"
	"github.com/gregsemple2003/CodexDesktop/backend/orchestration/internal/jobs"
)

type fakeBackend struct {
	schedules map[string]controlplane.RuntimeSchedule
}

func newFakeBackend() *fakeBackend {
	return &fakeBackend{schedules: map[string]controlplane.RuntimeSchedule{}}
}

func (b *fakeBackend) ListManagedSchedules(context.Context) ([]controlplane.RuntimeSchedule, error) {
	schedules := make([]controlplane.RuntimeSchedule, 0, len(b.schedules))
	for _, schedule := range b.schedules {
		schedules = append(schedules, schedule)
	}
	return schedules, nil
}

func (b *fakeBackend) CreateSchedule(_ context.Context, desired controlplane.DesiredSchedule) error {
	b.schedules[desired.ScheduleID] = controlplane.RuntimeSchedule{
		ScheduleID:      desired.ScheduleID,
		JobID:           desired.JobID,
		TriggerIndex:    desired.TriggerIndex,
		CronExpressions: []string{desired.Cron},
		TimeZoneName:    desired.Timezone,
		ManagedCron:     desired.Cron,
		ManagedTimezone: desired.Timezone,
		WorkflowType:    desired.WorkflowType,
		TaskQueue:       desired.TaskQueue,
	}
	return nil
}

func (b *fakeBackend) UpdateSchedule(_ context.Context, desired controlplane.DesiredSchedule) error {
	return b.CreateSchedule(context.Background(), desired)
}

func (b *fakeBackend) DeleteSchedule(_ context.Context, scheduleID string) error {
	delete(b.schedules, scheduleID)
	return nil
}

func (b *fakeBackend) Close() error {
	return nil
}

func TestMuxExposesHealthJobsAndSync(t *testing.T) {
	root := writeJobsRoot(t, []jobs.Spec{
		{
			APIVersion:   jobs.APIVersion,
			JobID:        "codex-daily-ue-determinism-digest",
			Label:        "UE Determinism Digest",
			Description:  "Daily digest",
			DesiredState: jobs.DesiredStateEnabled,
			Triggers: []jobs.Trigger{
				{Type: jobs.TriggerTypeSchedule, Cron: "0 4 * * *", Timezone: "America/Toronto"},
				{Type: jobs.TriggerTypeManual},
				{Type: jobs.TriggerTypeWebhook, Path: "digests/ue-determinism"},
			},
			Executor: jobs.Executor{
				Type:       jobs.ExecutorTypeCodexExec,
				Cwd:        `C:\Users\gregs\.codex`,
				Entrypoint: "ue-determinism-digest",
			},
			Runtime: jobs.RuntimeConfig{
				WorkflowType: "codex.exec.job",
				TaskQueue:    "codex-orchestration",
			},
		},
	})

	service := controlplane.NewService(root, newFakeBackend())
	mux := NewMux(config.Config{
		BindAddress:     "127.0.0.1:4318",
		JobsRoot:        root,
		Namespace:       "default",
		TaskQueue:       "codex-orchestration",
		TemporalAddress: "127.0.0.1:7233",
	}, service)

	syncRequest := httptest.NewRequest(http.MethodPost, "/sync", nil)
	syncResponse := httptest.NewRecorder()
	mux.ServeHTTP(syncResponse, syncRequest)
	if syncResponse.Code != http.StatusOK {
		t.Fatalf("POST /sync status = %d, want 200", syncResponse.Code)
	}

	jobsRequest := httptest.NewRequest(http.MethodGet, "/jobs", nil)
	jobsResponse := httptest.NewRecorder()
	mux.ServeHTTP(jobsResponse, jobsRequest)
	if jobsResponse.Code != http.StatusOK {
		t.Fatalf("GET /jobs status = %d, want 200", jobsResponse.Code)
	}
	var jobsPayload struct {
		Jobs []controlplane.JobView `json:"jobs"`
	}
	if err := json.Unmarshal(jobsResponse.Body.Bytes(), &jobsPayload); err != nil {
		t.Fatalf("decode /jobs: %v", err)
	}
	if len(jobsPayload.Jobs) != 1 {
		t.Fatalf("jobs payload count = %d, want 1", len(jobsPayload.Jobs))
	}

	jobDetailRequest := httptest.NewRequest(http.MethodGet, "/api/v1/jobs/codex-daily-ue-determinism-digest", nil)
	jobDetailResponse := httptest.NewRecorder()
	mux.ServeHTTP(jobDetailResponse, jobDetailRequest)
	if jobDetailResponse.Code != http.StatusOK {
		t.Fatalf("GET /api/v1/jobs/{id} status = %d, want 200", jobDetailResponse.Code)
	}

	jobsAPIRequest := httptest.NewRequest(http.MethodGet, "/api/v1/jobs", nil)
	jobsAPIResponse := httptest.NewRecorder()
	mux.ServeHTTP(jobsAPIResponse, jobsAPIRequest)
	if jobsAPIResponse.Code != http.StatusOK {
		t.Fatalf("GET /api/v1/jobs status = %d, want 200", jobsAPIResponse.Code)
	}

	healthRequest := httptest.NewRequest(http.MethodGet, "/health", nil)
	healthResponse := httptest.NewRecorder()
	mux.ServeHTTP(healthResponse, healthRequest)
	if healthResponse.Code != http.StatusOK {
		t.Fatalf("GET /health status = %d, want 200", healthResponse.Code)
	}

	runsRequest := httptest.NewRequest(http.MethodGet, "/runs?job_id=codex-daily-ue-determinism-digest", nil)
	runsResponse := httptest.NewRecorder()
	mux.ServeHTTP(runsResponse, runsRequest)
	if runsResponse.Code != http.StatusOK {
		t.Fatalf("GET /runs status = %d, want 200", runsResponse.Code)
	}

	runsAPIRequest := httptest.NewRequest(http.MethodGet, "/api/v1/jobs/codex-daily-ue-determinism-digest/runs", nil)
	runsAPIResponse := httptest.NewRecorder()
	mux.ServeHTTP(runsAPIResponse, runsAPIRequest)
	if runsAPIResponse.Code != http.StatusOK {
		t.Fatalf("GET /api/v1/jobs/{id}/runs status = %d, want 200", runsAPIResponse.Code)
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
