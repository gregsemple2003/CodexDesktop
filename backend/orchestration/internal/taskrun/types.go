package taskrun

import (
	"context"
	"errors"
	"time"
)

const ExecutionFailureModeTask0008WorkloadFailureOnce = "task_0008_workload_execution_failure_once"

var ErrRunNotFound = errors.New("task run not found")

const (
	StateReady             = "ready"
	StateQueued            = "queued"
	StateDispatching       = "dispatching"
	StateRunning           = "running"
	StateWaitingForHuman   = "waiting_for_human"
	StateBlocked           = "blocked"
	StateSleepingOrStalled = "sleeping_or_stalled"
	StateInterrupted       = "interrupted"
	StateCompleted         = "completed"
	StateFailed            = "failed"
)

const (
	AttentionNone           = "none"
	AttentionWatch          = "watch"
	AttentionNeedsAttention = "needs_attention"
	AttentionUrgent         = "urgent"
)

const (
	ActionDispatch  = "dispatch"
	ActionPoke      = "poke"
	ActionInterrupt = "interrupt"
)

type StateEnvelope struct {
	State              string                         `json:"state"`
	ReasonCode         string                         `json:"reason_code"`
	StateSummary       string                         `json:"state_summary"`
	EvidenceRefs       []EvidenceRef                  `json:"evidence_refs,omitempty"`
	NextOwner          string                         `json:"next_owner,omitempty"`
	NextExpectedEvent  string                         `json:"next_expected_event,omitempty"`
	SuspiciousAfter    time.Time                      `json:"suspicious_after,omitempty"`
	ActionBlockReasons map[string][]ActionBlockReason `json:"action_block_reasons,omitempty"`
}

type EvidenceRef struct {
	Type      string `json:"type"`
	Label     string `json:"label"`
	URI       string `json:"uri"`
	LineRange string `json:"line_range,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

type ActionBlockReason struct {
	Code    string `json:"code"`
	Summary string `json:"summary"`
}

type HumanActionTarget struct {
	Kind      string `json:"kind"`
	Label     string `json:"label"`
	URI       string `json:"uri"`
	LineRange string `json:"line_range,omitempty"`
}

type WaitContract struct {
	WaitingOn           string             `json:"waiting_on,omitempty"`
	WhyBlocked          string             `json:"why_blocked,omitempty"`
	ResumeWhen          string             `json:"resume_when,omitempty"`
	HumanActionRequired bool               `json:"human_action_required,omitempty"`
	HumanActionTarget   *HumanActionTarget `json:"human_action_target,omitempty"`
	NextOwner           string             `json:"next_owner,omitempty"`
	StaleAfter          time.Time          `json:"stale_after,omitempty"`
	EvidenceRefs        []EvidenceRef      `json:"evidence_refs,omitempty"`
}

type DispatchReadiness struct {
	Ready                bool                `json:"ready"`
	BlockReasons         []ActionBlockReason `json:"block_reasons,omitempty"`
	ExpectedFirstSignal  string              `json:"expected_first_signal,omitempty"`
	FirstSuspiciousAfter time.Time           `json:"first_suspicious_after,omitempty"`
}

type AttentionPriority struct {
	Level   string `json:"attention_level"`
	Reason  string `json:"attention_reason,omitempty"`
	SortKey string `json:"attention_sort_key,omitempty"`
}

type ActionAvailability struct {
	Allowed      bool                `json:"allowed"`
	BlockReasons []ActionBlockReason `json:"block_reasons,omitempty"`
}

type StoryOwnership struct {
	OwnerRunID string `json:"owner_run_id,omitempty"`
	Status     string `json:"status"`
	Reason     string `json:"reason,omitempty"`
}

type RepoLane struct {
	OwnedRepoRoot         string    `json:"owned_repo_root,omitempty"`
	CheckoutMode          string    `json:"checkout_mode,omitempty"`
	BaselineCommit        string    `json:"baseline_commit,omitempty"`
	CurrentCommit         string    `json:"current_commit,omitempty"`
	ApprovedRestoreCommit string    `json:"approved_restore_commit,omitempty"`
	RunArtifactRoot       string    `json:"run_artifact_root,omitempty"`
	BootstrapArtifactPath string    `json:"bootstrap_artifact_path,omitempty"`
	PreflightArtifactPath string    `json:"preflight_artifact_path,omitempty"`
	WorkloadStepPath      string    `json:"workload_step_path,omitempty"`
	WorkloadResultPath    string    `json:"workload_result_path,omitempty"`
	WorkloadOutputPath    string    `json:"workload_output_path,omitempty"`
	WorkloadCodePath      string    `json:"workload_code_path,omitempty"`
	ResetStatus           string    `json:"reset_status,omitempty"`
	LastResetTargetCommit string    `json:"last_reset_target_commit,omitempty"`
	ResetFailureSummary   string    `json:"reset_failure_summary,omitempty"`
	LastResetAt           time.Time `json:"last_reset_at,omitempty"`
}

type RunFollowUp struct {
	Kind        string    `json:"kind"`
	Owner       string    `json:"owner"`
	Status      string    `json:"status"`
	Summary     string    `json:"summary"`
	RequestedAt time.Time `json:"requested_at,omitempty"`
	DueAt       time.Time `json:"due_at,omitempty"`
	CompletedAt time.Time `json:"completed_at,omitempty"`
}

type RunResolution struct {
	Kind       string    `json:"kind"`
	Decision   string    `json:"decision"`
	Summary    string    `json:"summary"`
	ResolvedBy string    `json:"resolved_by,omitempty"`
	ResolvedAt time.Time `json:"resolved_at,omitempty"`
}

type TaskDefinitionSnapshot struct {
	DeclaredWorktreeRoot string               `json:"declared_worktree_root"`
	DeclaredTaskRoot     string               `json:"declared_task_root"`
	DeclaredTaskRevision string               `json:"declared_task_revision"`
	DeclaredGitRevision  string               `json:"declared_git_revision,omitempty"`
	CapturedAt           time.Time            `json:"captured_at"`
	Files                []TaskArtifactDigest `json:"files,omitempty"`
}

type TaskArtifactDigest struct {
	RelativePath string `json:"relative_path"`
	SHA256       string `json:"sha256"`
}

type TaskRunView struct {
	RunID                       string                        `json:"run_id"`
	TaskID                      string                        `json:"task_id"`
	WorkflowID                  string                        `json:"workflow_id,omitempty"`
	TemporalExecutionRunID      string                        `json:"temporal_execution_run_id,omitempty"`
	Status                      string                        `json:"status"`
	StateEnvelope               StateEnvelope                 `json:"state_envelope"`
	MeaningSummary              string                        `json:"meaning_summary,omitempty"`
	WaitContract                *WaitContract                 `json:"wait_contract,omitempty"`
	Attention                   AttentionPriority             `json:"attention"`
	Actions                     map[string]ActionAvailability `json:"actions,omitempty"`
	FollowUp                    *RunFollowUp                  `json:"follow_up,omitempty"`
	Resolution                  *RunResolution                `json:"resolution,omitempty"`
	RepoLane                    RepoLane                      `json:"repo_lane"`
	LastProgressAt              time.Time                     `json:"last_progress_at,omitempty"`
	LastProgressSummary         string                        `json:"last_progress_summary,omitempty"`
	FailureSummary              string                        `json:"failure_summary,omitempty"`
	CapturedTaskSnapshot        TaskDefinitionSnapshot        `json:"captured_task_snapshot"`
	DocRuntimeDivergenceStatus  string                        `json:"doc_runtime_divergence_status,omitempty"`
	DocRuntimeDivergenceSummary string                        `json:"doc_runtime_divergence_summary,omitempty"`
}

type TaskView struct {
	TaskID               string                        `json:"task_id"`
	Title                string                        `json:"title"`
	MeaningSummary       string                        `json:"meaning_summary"`
	StateEnvelope        StateEnvelope                 `json:"state_envelope"`
	DispatchReadiness    DispatchReadiness             `json:"dispatch_readiness"`
	Attention            AttentionPriority             `json:"attention"`
	Actions              map[string]ActionAvailability `json:"actions"`
	DeclaredWorktreeRoot string                        `json:"declared_worktree_root"`
	DeclaredTaskRoot     string                        `json:"declared_task_root"`
	DeclaredTaskRevision string                        `json:"declared_task_revision"`
	DeclaredGitRevision  string                        `json:"declared_git_revision,omitempty"`
	CurrentStory         StoryOwnership                `json:"current_story"`
	LatestRun            *TaskRunView                  `json:"latest_run,omitempty"`
	CurrentGate          string                        `json:"current_gate,omitempty"`
	CurrentPass          string                        `json:"current_pass,omitempty"`
	Phase                string                        `json:"phase,omitempty"`
	PlanApproved         bool                          `json:"plan_approved"`
	Blockers             []string                      `json:"blockers,omitempty"`
	UpdatedAt            string                        `json:"updated_at,omitempty"`
}

type Runtime interface {
	StartTaskRun(ctx context.Context, request StartTaskRunRequest) (TaskRunView, error)
	GetTaskRun(ctx context.Context, runID string) (TaskRunView, error)
	GetActiveTaskRun(ctx context.Context, taskID string) (TaskRunView, error)
	ReconcileTaskSnapshot(ctx context.Context, runID string, snapshot TaskDefinitionSnapshot) (TaskRunView, error)
	UpdateTaskRun(ctx context.Context, runID string, update TaskRunUpdate) (TaskRunView, error)
	RetryTaskRunWorkload(ctx context.Context, runID string, request WorkloadRetryRequest) (TaskRunView, error)
}

type StartTaskRunRequest struct {
	RunID                string                 `json:"run_id"`
	TaskID               string                 `json:"task_id"`
	Title                string                 `json:"title"`
	MeaningSummary       string                 `json:"meaning_summary"`
	CapturedTaskSnapshot TaskDefinitionSnapshot `json:"captured_task_snapshot"`
	ExecutionDirective   *ExecutionDirective    `json:"execution_directive,omitempty"`
	RepoLane             RepoLane               `json:"repo_lane"`
	DispatchRequestedAt  time.Time              `json:"dispatch_requested_at"`
}

type ExecutionDirective struct {
	FailureMode string `json:"failure_mode,omitempty"`
}

type InterruptReviewResolution struct {
	Decision   string `json:"decision"`
	Summary    string `json:"summary,omitempty"`
	ResolvedBy string `json:"resolved_by,omitempty"`
}

type WorkloadRetryRequest struct {
	CapturedTaskSnapshot TaskDefinitionSnapshot `json:"captured_task_snapshot"`
	RepoLane             RepoLane               `json:"repo_lane"`
	RetryRequestedAt     time.Time              `json:"retry_requested_at"`
}

type TaskRunUpdate struct {
	State               string                        `json:"state"`
	ReasonCode          string                        `json:"reason_code"`
	StateSummary        string                        `json:"state_summary"`
	NextOwner           string                        `json:"next_owner,omitempty"`
	NextExpectedEvent   string                        `json:"next_expected_event,omitempty"`
	SuspiciousAfter     time.Time                     `json:"suspicious_after,omitempty"`
	LastProgressSummary string                        `json:"last_progress_summary,omitempty"`
	WaitContract        *WaitContract                 `json:"wait_contract,omitempty"`
	Attention           *AttentionPriority            `json:"attention,omitempty"`
	RepoLane            *RepoLane                     `json:"repo_lane,omitempty"`
	Actions             map[string]ActionAvailability `json:"actions,omitempty"`
	FollowUp            *RunFollowUp                  `json:"follow_up,omitempty"`
	Resolution          *RunResolution                `json:"resolution,omitempty"`
	CompletedAt         time.Time                     `json:"completed_at,omitempty"`
	FailureSummary      string                        `json:"failure_summary,omitempty"`
}
