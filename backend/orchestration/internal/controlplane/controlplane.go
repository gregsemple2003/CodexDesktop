package controlplane

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gregsemple2003/CodexDesktop/backend/orchestration/internal/jobs"
)

const (
	managedSchedulePrefix = "codex-job--"
)

type Backend interface {
	ListManagedSchedules(ctx context.Context) ([]RuntimeSchedule, error)
	CreateSchedule(ctx context.Context, desired DesiredSchedule) error
	UpdateSchedule(ctx context.Context, desired DesiredSchedule) error
	DeleteSchedule(ctx context.Context, scheduleID string) error
	Close() error
}

type Service struct {
	jobsRoot string
	backend  Backend

	mu       sync.RWMutex
	lastSync SyncMeta
}

type SyncMeta struct {
	LastAttemptAt time.Time `json:"last_attempt_at,omitempty"`
	LastSuccessAt time.Time `json:"last_success_at,omitempty"`
	LastError     string    `json:"last_error,omitempty"`
}

type SyncReport struct {
	StartedAt   time.Time `json:"started_at"`
	CompletedAt time.Time `json:"completed_at"`
	Created     []string  `json:"created_schedule_ids,omitempty"`
	Updated     []string  `json:"updated_schedule_ids,omitempty"`
	Deleted     []string  `json:"deleted_schedule_ids,omitempty"`
	State       StateView `json:"state"`
}

type StateView struct {
	GeneratedAt time.Time `json:"generated_at"`
	LastSync    SyncMeta  `json:"last_sync"`
	Jobs        []JobView `json:"jobs"`
}

type JobView struct {
	JobID        string             `json:"job_id"`
	Label        string             `json:"label"`
	Description  string             `json:"description"`
	DesiredState string             `json:"desired_state"`
	Status       string             `json:"status"`
	Drift        []string           `json:"drift,omitempty"`
	Triggers     []jobs.Trigger     `json:"triggers"`
	Executor     jobs.Executor      `json:"executor"`
	Runtime      jobs.RuntimeConfig `json:"runtime"`
	Schedules    []ScheduleView     `json:"schedules"`
	RecentRuns   []RunRecord        `json:"recent_runs,omitempty"`
}

type ScheduleView struct {
	ScheduleID             string      `json:"schedule_id"`
	TriggerIndex           int         `json:"trigger_index"`
	Status                 string      `json:"status"`
	Drift                  []string    `json:"drift,omitempty"`
	DesiredCron            string      `json:"desired_cron,omitempty"`
	DesiredTimezone        string      `json:"desired_timezone,omitempty"`
	RuntimeCronExpressions []string    `json:"runtime_cron_expressions,omitempty"`
	RuntimeTimezone        string      `json:"runtime_timezone,omitempty"`
	WorkflowType           string      `json:"workflow_type,omitempty"`
	TaskQueue              string      `json:"task_queue,omitempty"`
	Paused                 bool        `json:"paused,omitempty"`
	Note                   string      `json:"note,omitempty"`
	NextActionTimes        []time.Time `json:"next_action_times,omitempty"`
	RecentRuns             []RunRecord `json:"recent_runs,omitempty"`
}

type RunRecord struct {
	ScheduleID          string    `json:"schedule_id"`
	ScheduleTime        time.Time `json:"schedule_time"`
	ActualTime          time.Time `json:"actual_time"`
	WorkflowID          string    `json:"workflow_id,omitempty"`
	FirstExecutionRunID string    `json:"first_execution_run_id,omitempty"`
}

type DesiredSchedule struct {
	ScheduleID   string
	JobID        string
	TriggerIndex int
	Enabled      bool
	Cron         string
	Timezone     string
	WorkflowType string
	TaskQueue    string
	WorkflowID   string
}

type RuntimeSchedule struct {
	ScheduleID      string
	JobID           string
	TriggerIndex    int
	CronExpressions []string
	TimeZoneName    string
	ManagedCron     string
	ManagedTimezone string
	WorkflowType    string
	TaskQueue       string
	Paused          bool
	Note            string
	RecentRuns      []RunRecord
	NextActionTimes []time.Time
}

type compiledJob struct {
	Spec      jobs.Spec
	Schedules []DesiredSchedule
}

func NewService(jobsRoot string, backend Backend) *Service {
	return &Service{
		jobsRoot: jobsRoot,
		backend:  backend,
	}
}

func (s *Service) Close() error {
	if s.backend == nil {
		return nil
	}
	return s.backend.Close()
}

func (s *Service) Snapshot(ctx context.Context) (StateView, error) {
	compiled, err := s.loadCompiledJobs()
	if err != nil {
		return StateView{}, err
	}
	runtimeSchedules, err := s.backend.ListManagedSchedules(ctx)
	if err != nil {
		return StateView{}, err
	}
	return buildState(compiled, runtimeSchedules, time.Now().UTC(), s.syncMeta()), nil
}

func (s *Service) Reconcile(ctx context.Context) (SyncReport, error) {
	startedAt := time.Now().UTC()
	s.recordAttempt(startedAt, "")

	compiled, err := s.loadCompiledJobs()
	if err != nil {
		s.recordAttempt(startedAt, err.Error())
		return SyncReport{}, err
	}

	runtimeSchedules, err := s.backend.ListManagedSchedules(ctx)
	if err != nil {
		s.recordAttempt(startedAt, err.Error())
		return SyncReport{}, err
	}

	desiredByID := map[string]DesiredSchedule{}
	for _, job := range compiled {
		for _, schedule := range job.Schedules {
			if schedule.Enabled {
				desiredByID[schedule.ScheduleID] = schedule
			}
		}
	}

	runtimeByID := map[string]RuntimeSchedule{}
	for _, schedule := range runtimeSchedules {
		runtimeByID[schedule.ScheduleID] = schedule
	}

	report := SyncReport{
		StartedAt: startedAt,
	}

	for _, job := range compiled {
		for _, desired := range job.Schedules {
			if !desired.Enabled {
				continue
			}
			runtime, exists := runtimeByID[desired.ScheduleID]
			if !exists {
				if err := s.backend.CreateSchedule(ctx, desired); err != nil {
					s.recordAttempt(startedAt, err.Error())
					return SyncReport{}, err
				}
				report.Created = append(report.Created, desired.ScheduleID)
				continue
			}
			if drift := diffSchedule(desired, runtime); len(drift) > 0 {
				if err := s.backend.UpdateSchedule(ctx, desired); err != nil {
					s.recordAttempt(startedAt, err.Error())
					return SyncReport{}, err
				}
				report.Updated = append(report.Updated, desired.ScheduleID)
			}
		}
	}

	for scheduleID := range runtimeByID {
		if _, exists := desiredByID[scheduleID]; exists {
			continue
		}
		if err := s.backend.DeleteSchedule(ctx, scheduleID); err != nil {
			s.recordAttempt(startedAt, err.Error())
			return SyncReport{}, err
		}
		report.Deleted = append(report.Deleted, scheduleID)
	}

	runtimeAfter, err := s.backend.ListManagedSchedules(ctx)
	if err != nil {
		s.recordAttempt(startedAt, err.Error())
		return SyncReport{}, err
	}

	completedAt := time.Now().UTC()
	s.recordAttempt(completedAt, "")
	state := buildState(compiled, runtimeAfter, completedAt, s.syncMeta())

	report.CompletedAt = completedAt
	report.State = state
	return report, nil
}

func (s *Service) Job(ctx context.Context, jobID string) (JobView, error) {
	state, err := s.Snapshot(ctx)
	if err != nil {
		return JobView{}, err
	}
	for _, job := range state.Jobs {
		if job.JobID == jobID {
			return job, nil
		}
	}
	return JobView{}, fmt.Errorf("job %q not found", jobID)
}

func (s *Service) Runs(ctx context.Context, jobID string) ([]RunRecord, error) {
	job, err := s.Job(ctx, jobID)
	if err != nil {
		return nil, err
	}
	return append([]RunRecord{}, job.RecentRuns...), nil
}

func (s *Service) loadCompiledJobs() ([]compiledJob, error) {
	specs, err := jobs.LoadSpecs(s.jobsRoot)
	if err != nil {
		return nil, err
	}

	compiled := make([]compiledJob, 0, len(specs))
	for _, spec := range specs {
		compiled = append(compiled, compileJob(spec))
	}
	return compiled, nil
}

func compileJob(spec jobs.Spec) compiledJob {
	job := compiledJob{
		Spec:      spec,
		Schedules: make([]DesiredSchedule, 0),
	}

	scheduleIndex := 0
	for _, trigger := range spec.Triggers {
		if trigger.Type != jobs.TriggerTypeSchedule {
			continue
		}
		job.Schedules = append(job.Schedules, DesiredSchedule{
			ScheduleID:   ManagedScheduleID(spec.JobID, scheduleIndex),
			JobID:        spec.JobID,
			TriggerIndex: scheduleIndex,
			Enabled:      spec.DesiredState == jobs.DesiredStateEnabled,
			Cron:         trigger.Cron,
			Timezone:     trigger.Timezone,
			WorkflowType: spec.Runtime.WorkflowType,
			TaskQueue:    spec.Runtime.TaskQueue,
			WorkflowID:   fmt.Sprintf("%s/schedule/%02d", spec.JobID, scheduleIndex),
		})
		scheduleIndex++
	}
	return job
}

func ManagedScheduleID(jobID string, scheduleIndex int) string {
	return fmt.Sprintf("%s%s--%02d", managedSchedulePrefix, jobID, scheduleIndex)
}

func ParseManagedScheduleID(scheduleID string) (string, int, bool) {
	if !strings.HasPrefix(scheduleID, managedSchedulePrefix) {
		return "", 0, false
	}
	raw := strings.TrimPrefix(scheduleID, managedSchedulePrefix)
	lastSep := strings.LastIndex(raw, "--")
	if lastSep < 0 {
		return "", 0, false
	}
	index, err := strconv.Atoi(raw[lastSep+2:])
	if err != nil {
		return "", 0, false
	}
	return raw[:lastSep], index, true
}

func buildState(compiled []compiledJob, runtime []RuntimeSchedule, generatedAt time.Time, syncMeta SyncMeta) StateView {
	runtimeByID := map[string]RuntimeSchedule{}
	for _, schedule := range runtime {
		runtimeByID[schedule.ScheduleID] = schedule
	}

	jobsByID := map[string]*JobView{}
	seenRuntime := map[string]struct{}{}

	for _, compiledJob := range compiled {
		job := JobView{
			JobID:        compiledJob.Spec.JobID,
			Label:        compiledJob.Spec.Label,
			Description:  compiledJob.Spec.Description,
			DesiredState: compiledJob.Spec.DesiredState,
			Status:       "in_sync",
			Triggers:     append([]jobs.Trigger(nil), compiledJob.Spec.Triggers...),
			Executor:     compiledJob.Spec.Executor,
			Runtime:      compiledJob.Spec.Runtime,
			Schedules:    make([]ScheduleView, 0, len(compiledJob.Schedules)),
		}

		for _, desired := range compiledJob.Schedules {
			runtimeSchedule, exists := runtimeByID[desired.ScheduleID]
			if exists {
				seenRuntime[desired.ScheduleID] = struct{}{}
			}

			view := ScheduleView{
				ScheduleID:      desired.ScheduleID,
				TriggerIndex:    desired.TriggerIndex,
				DesiredCron:     desired.Cron,
				DesiredTimezone: desired.Timezone,
				WorkflowType:    desired.WorkflowType,
				TaskQueue:       desired.TaskQueue,
			}

			switch {
			case !desired.Enabled && exists:
				view.Status = "drifted"
				view.Drift = []string{"unexpected_runtime_schedule"}
				job.Drift = append(job.Drift, fmt.Sprintf("%s:unexpected_runtime_schedule", desired.ScheduleID))
				job.Status = "drifted"
				populateRuntimeScheduleView(&view, runtimeSchedule)
			case !desired.Enabled:
				view.Status = "disabled"
			case !exists:
				view.Status = "missing"
				view.Drift = []string{"missing_runtime_schedule"}
				job.Drift = append(job.Drift, fmt.Sprintf("%s:missing_runtime_schedule", desired.ScheduleID))
				job.Status = "drifted"
			default:
				populateRuntimeScheduleView(&view, runtimeSchedule)
				drift := diffSchedule(desired, runtimeSchedule)
				if len(drift) > 0 {
					view.Status = "drifted"
					view.Drift = drift
					job.Drift = append(job.Drift, fmt.Sprintf("%s:%s", desired.ScheduleID, strings.Join(drift, ",")))
					job.Status = "drifted"
				} else {
					view.Status = "in_sync"
				}
			}

			job.Schedules = append(job.Schedules, view)
		}

		if compiledJob.Spec.DesiredState == jobs.DesiredStateDisabled && job.Status != "drifted" {
			job.Status = "disabled"
		}

		job.RecentRuns = aggregateRecentRuns(job.Schedules)
		sortSchedules(job.Schedules)
		jobsByID[job.JobID] = &job
	}

	for _, runtimeSchedule := range runtime {
		if _, alreadyMatched := seenRuntime[runtimeSchedule.ScheduleID]; alreadyMatched {
			continue
		}
		jobID := runtimeSchedule.JobID
		if jobID == "" {
			if parsedJobID, _, ok := ParseManagedScheduleID(runtimeSchedule.ScheduleID); ok {
				jobID = parsedJobID
			} else {
				jobID = runtimeSchedule.ScheduleID
			}
		}

		job, exists := jobsByID[jobID]
		if !exists {
			job = &JobView{
				JobID:        jobID,
				Label:        jobID,
				Description:  "Runtime schedule exists without a tracked desired-state job spec.",
				DesiredState: "missing",
				Status:       "drifted",
			}
			jobsByID[jobID] = job
		}

		view := ScheduleView{
			ScheduleID:   runtimeSchedule.ScheduleID,
			TriggerIndex: runtimeSchedule.TriggerIndex,
			Status:       "extra",
			Drift:        []string{"unexpected_runtime_schedule"},
		}
		populateRuntimeScheduleView(&view, runtimeSchedule)
		job.Schedules = append(job.Schedules, view)
		job.Drift = append(job.Drift, fmt.Sprintf("%s:unexpected_runtime_schedule", runtimeSchedule.ScheduleID))
		job.Status = "drifted"
		job.RecentRuns = aggregateRecentRuns(job.Schedules)
	}

	state := StateView{
		GeneratedAt: generatedAt,
		LastSync:    syncMeta,
		Jobs:        make([]JobView, 0, len(jobsByID)),
	}

	for _, job := range jobsByID {
		sortSchedules(job.Schedules)
		job.RecentRuns = aggregateRecentRuns(job.Schedules)
		sort.Strings(job.Drift)
		state.Jobs = append(state.Jobs, *job)
	}

	sort.Slice(state.Jobs, func(i, j int) bool {
		return state.Jobs[i].JobID < state.Jobs[j].JobID
	})

	return state
}

func diffSchedule(desired DesiredSchedule, runtime RuntimeSchedule) []string {
	drift := make([]string, 0, 4)

	runtimeCron := runtime.ManagedCron
	if runtimeCron == "" && len(runtime.CronExpressions) > 0 {
		runtimeCron = runtime.CronExpressions[0]
	}
	if runtimeCron != desired.Cron {
		drift = append(drift, "cron")
	}

	runtimeTimezone := runtime.ManagedTimezone
	if runtimeTimezone == "" {
		runtimeTimezone = runtime.TimeZoneName
	}
	if runtimeTimezone != desired.Timezone {
		drift = append(drift, "timezone")
	}

	if runtime.WorkflowType != desired.WorkflowType {
		drift = append(drift, "workflow_type")
	}
	if runtime.TaskQueue != desired.TaskQueue {
		drift = append(drift, "task_queue")
	}
	if runtime.Paused {
		drift = append(drift, "paused")
	}

	return drift
}

func populateRuntimeScheduleView(view *ScheduleView, runtime RuntimeSchedule) {
	view.RuntimeCronExpressions = append([]string(nil), runtime.CronExpressions...)
	view.RuntimeTimezone = runtime.TimeZoneName
	if runtime.ManagedTimezone != "" {
		view.RuntimeTimezone = runtime.ManagedTimezone
	}
	if view.WorkflowType == "" {
		view.WorkflowType = runtime.WorkflowType
	}
	if view.TaskQueue == "" {
		view.TaskQueue = runtime.TaskQueue
	}
	view.Paused = runtime.Paused
	view.Note = runtime.Note
	view.NextActionTimes = append([]time.Time(nil), runtime.NextActionTimes...)
	view.RecentRuns = append([]RunRecord(nil), runtime.RecentRuns...)
}

func aggregateRecentRuns(schedules []ScheduleView) []RunRecord {
	runs := make([]RunRecord, 0)
	for _, schedule := range schedules {
		runs = append(runs, schedule.RecentRuns...)
	}
	sort.Slice(runs, func(i, j int) bool {
		left := runs[i].ActualTime
		if left.IsZero() {
			left = runs[i].ScheduleTime
		}
		right := runs[j].ActualTime
		if right.IsZero() {
			right = runs[j].ScheduleTime
		}
		return left.After(right)
	})
	if len(runs) > 10 {
		runs = runs[:10]
	}
	return runs
}

func sortSchedules(schedules []ScheduleView) {
	sort.Slice(schedules, func(i, j int) bool {
		return schedules[i].ScheduleID < schedules[j].ScheduleID
	})
}

func (s *Service) recordAttempt(at time.Time, errText string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.lastSync.LastAttemptAt = at
	s.lastSync.LastError = errText
	if errText == "" {
		s.lastSync.LastSuccessAt = at
	}
}

func (s *Service) syncMeta() SyncMeta {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.lastSync
}
