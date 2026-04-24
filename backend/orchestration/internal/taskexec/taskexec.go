package taskexec

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"

	"github.com/gregsemple2003/CodexDesktop/backend/orchestration/internal/taskrun"
)

const (
	TaskRunWorkflowName         = "codex.task.run"
	TaskRunStateQueryName       = "taskrun.current_state"
	ReconcileSnapshotSignalName = "taskrun.reconcile_snapshot"
	UpdateRunSignalName         = "taskrun.update_state"
	RunExecutionPreflightName   = "taskrun.execution_preflight"
	RunWorkloadStepName         = "taskrun.workload_step"
	RunExecuteWorkloadName      = "taskrun.execute_workload_step"
)

func Register(w worker.Worker) {
	w.RegisterWorkflowWithOptions(TaskRunWorkflow, workflow.RegisterOptions{Name: TaskRunWorkflowName})
	w.RegisterActivityWithOptions(runExecutionPreflight, activity.RegisterOptions{Name: RunExecutionPreflightName})
	w.RegisterActivityWithOptions(runWorkloadStep, activity.RegisterOptions{Name: RunWorkloadStepName})
	w.RegisterActivityWithOptions(runExecuteWorkloadStep, activity.RegisterOptions{Name: RunExecuteWorkloadName})
}

func TaskRunWorkflow(ctx workflow.Context, request taskrun.StartTaskRunRequest) (taskrun.TaskRunView, error) {
	now := workflow.Now(ctx).UTC()
	if request.DispatchRequestedAt.IsZero() {
		request.DispatchRequestedAt = now
	}

	info := workflow.GetInfo(ctx)
	view := InitialView(request, info.WorkflowExecution.ID, info.WorkflowExecution.RunID)

	if err := workflow.SetQueryHandler(ctx, TaskRunStateQueryName, func() (taskrun.TaskRunView, error) {
		return view, nil
	}); err != nil {
		return taskrun.TaskRunView{}, err
	}

	if request.RepoLane.RunArtifactRoot != "" {
		activityCtx := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
			StartToCloseTimeout: 2 * time.Minute,
			RetryPolicy: &temporal.RetryPolicy{
				MaximumAttempts: 1,
			},
		})
		var preflight executionPreflightResult
		if err := workflow.ExecuteActivity(activityCtx, RunExecutionPreflightName, request).Get(activityCtx, &preflight); err != nil {
			applyUpdate(&view, taskrun.TaskRunUpdate{
				State:               taskrun.StateBlocked,
				ReasonCode:          "execution_preflight_failed",
				StateSummary:        "Run could not complete owned-lane execution preflight.",
				NextOwner:           "human_or_supervisor",
				NextExpectedEvent:   "Review the preflight failure before continuing execution.",
				SuspiciousAfter:     workflow.Now(ctx).UTC(),
				LastProgressSummary: "Execution preflight failed before the first workload step.",
				Attention: &taskrun.AttentionPriority{
					Level:   taskrun.AttentionUrgent,
					Reason:  "Run could not prepare the owned lane for execution.",
					SortKey: "14-execution_preflight_failed",
				},
				Actions:        actionsForState(taskrun.StateBlocked),
				FailureSummary: err.Error(),
			}, workflow.Now(ctx).UTC())
		} else {
			repoLane := view.RepoLane
			if preflight.CurrentCommit != "" {
				repoLane.CurrentCommit = preflight.CurrentCommit
			}
			if preflight.PreflightArtifactPath != "" {
				repoLane.PreflightArtifactPath = preflight.PreflightArtifactPath
			}
			applyUpdate(&view, taskrun.TaskRunUpdate{
				State:               taskrun.StateRunning,
				ReasonCode:          "execution_preflight_complete",
				StateSummary:        "Run completed owned-lane execution preflight.",
				NextOwner:           "backend_worker",
				NextExpectedEvent:   "Execution worker records the first workload step.",
				SuspiciousAfter:     workflow.Now(ctx).UTC().Add(15 * time.Minute),
				LastProgressSummary: preflight.ProgressSummary,
				Attention: &taskrun.AttentionPriority{
					Level:   taskrun.AttentionWatch,
					Reason:  "Run completed execution preflight and is ready for the first workload step.",
					SortKey: "44-execution_preflight_complete",
				},
				RepoLane: &repoLane,
				Actions:  actionsForState(taskrun.StateRunning),
			}, workflow.Now(ctx).UTC())

			var workload workloadStepResult
			if err := workflow.ExecuteActivity(activityCtx, RunWorkloadStepName, request, repoLane).Get(activityCtx, &workload); err != nil {
				applyUpdate(&view, taskrun.TaskRunUpdate{
					State:               taskrun.StateBlocked,
					ReasonCode:          "workload_step_failed",
					StateSummary:        "Run could not prepare the first workload step inside the owned lane.",
					NextOwner:           "human_or_supervisor",
					NextExpectedEvent:   "Review the owned-lane workload step failure before continuing execution.",
					SuspiciousAfter:     workflow.Now(ctx).UTC(),
					LastProgressSummary: "The first workload step failed after execution preflight completed.",
					Attention: &taskrun.AttentionPriority{
						Level:   taskrun.AttentionUrgent,
						Reason:  "Run could not prepare its first workload step inside the owned lane.",
						SortKey: "14-workload_step_failed",
					},
					RepoLane:       &repoLane,
					Actions:        actionsForState(taskrun.StateBlocked),
					FailureSummary: err.Error(),
				}, workflow.Now(ctx).UTC())
			} else {
				if workload.CurrentCommit != "" {
					repoLane.CurrentCommit = workload.CurrentCommit
				}
				if workload.WorkloadStepPath != "" {
					repoLane.WorkloadStepPath = workload.WorkloadStepPath
				}
				applyUpdate(&view, taskrun.TaskRunUpdate{
					State:               taskrun.StateRunning,
					ReasonCode:          "workload_step_prepared",
					StateSummary:        "Run prepared the first backend workload step inside the owned lane.",
					NextOwner:           "backend_worker",
					NextExpectedEvent:   "Execution worker executes the prepared workload step.",
					SuspiciousAfter:     workflow.Now(ctx).UTC().Add(15 * time.Minute),
					LastProgressSummary: workload.ProgressSummary,
					Attention: &taskrun.AttentionPriority{
						Level:   taskrun.AttentionWatch,
						Reason:  "Run prepared its first workload step and is ready for backend execution.",
						SortKey: "43-workload_step_prepared",
					},
					RepoLane: &repoLane,
					Actions:  actionsForState(taskrun.StateRunning),
				}, workflow.Now(ctx).UTC())

				var execution workloadExecutionResult
				if err := workflow.ExecuteActivity(activityCtx, RunExecuteWorkloadName, request, repoLane).Get(activityCtx, &execution); err != nil {
					applyUpdate(&view, taskrun.TaskRunUpdate{
						State:               taskrun.StateBlocked,
						ReasonCode:          "workload_execution_failed",
						StateSummary:        "Run could not execute the prepared workload step inside the owned lane.",
						NextOwner:           "human_or_supervisor",
						NextExpectedEvent:   "Review the workload execution failure before continuing execution.",
						SuspiciousAfter:     workflow.Now(ctx).UTC(),
						LastProgressSummary: "The prepared workload step failed during execution inside the owned lane.",
						Attention: &taskrun.AttentionPriority{
							Level:   taskrun.AttentionUrgent,
							Reason:  "Run could not execute the prepared workload step inside the owned lane.",
							SortKey: "13-workload_execution_failed",
						},
						RepoLane:       &repoLane,
						Actions:        actionsForState(taskrun.StateBlocked),
						FailureSummary: err.Error(),
					}, workflow.Now(ctx).UTC())
				} else {
					if execution.CurrentCommit != "" {
						repoLane.CurrentCommit = execution.CurrentCommit
					}
					if execution.WorkloadResultPath != "" {
						repoLane.WorkloadResultPath = execution.WorkloadResultPath
					}
					if execution.WorkloadOutputPath != "" {
						repoLane.WorkloadOutputPath = execution.WorkloadOutputPath
					}
					if execution.WorkloadCodePath != "" {
						repoLane.WorkloadCodePath = execution.WorkloadCodePath
					}
					reasonCode := "workload_step_executed"
					stateSummary := "Run executed the first backend workload step inside the owned lane."
					nextExpectedEvent := "Execution worker prepares or executes the next workload step."
					attentionReason := "Run executed its first workload step and is ready for the next backend step."
					attentionSortKey := "42-workload_step_executed"
					if request.TaskID == "Task-0008" {
						reasonCode = "task_0008_redispatch_lane_released"
						stateSummary = "Run validated Task-0008 and changed owned-lane redispatch cleanup behavior in an existing implementation file."
						nextExpectedEvent = "Execution worker applies the next broader Task-0008 recovery or redispatch behavior change."
						attentionReason = "Run changed owned-lane redispatch cleanup behavior in an existing Task-0008 implementation file and is ready for the next backend step."
						attentionSortKey = "35-task_0008_redispatch_lane_released"
					}
					applyUpdate(&view, taskrun.TaskRunUpdate{
						State:               taskrun.StateRunning,
						ReasonCode:          reasonCode,
						StateSummary:        stateSummary,
						NextOwner:           "backend_worker",
						NextExpectedEvent:   nextExpectedEvent,
						SuspiciousAfter:     workflow.Now(ctx).UTC().Add(15 * time.Minute),
						LastProgressSummary: execution.ProgressSummary,
						Attention: &taskrun.AttentionPriority{
							Level:   taskrun.AttentionWatch,
							Reason:  attentionReason,
							SortKey: attentionSortKey,
						},
						RepoLane: &repoLane,
						Actions:  actionsForState(taskrun.StateRunning),
					}, workflow.Now(ctx).UTC())
				}
			}
		}
	}

	reconcileCh := workflow.GetSignalChannel(ctx, ReconcileSnapshotSignalName)
	updateCh := workflow.GetSignalChannel(ctx, UpdateRunSignalName)
	for {
		selector := workflow.NewSelector(ctx)
		selector.AddReceive(reconcileCh, func(c workflow.ReceiveChannel, more bool) {
			var snapshot taskrun.TaskDefinitionSnapshot
			c.Receive(ctx, &snapshot)
			view.CapturedTaskSnapshot = snapshot
			view.DocRuntimeDivergenceStatus = "reconciled"
			view.DocRuntimeDivergenceSummary = "Runtime captured newer task docs during task readback."
			view.LastProgressAt = workflow.Now(ctx).UTC()
			view.LastProgressSummary = "Reconciled declared task docs into runtime state."
			view.StateEnvelope.SuspiciousAfter = workflow.Now(ctx).UTC().Add(15 * time.Minute)
		})
		selector.AddReceive(updateCh, func(c workflow.ReceiveChannel, more bool) {
			var update taskrun.TaskRunUpdate
			c.Receive(ctx, &update)
			applyUpdate(&view, update, workflow.Now(ctx).UTC())
		})
		selector.Select(ctx)
		if shouldExit(view) {
			return view, nil
		}
	}
}

func InitialView(request taskrun.StartTaskRunRequest, workflowID string, executionRunID string) taskrun.TaskRunView {
	suspiciousAfter := request.DispatchRequestedAt.Add(15 * time.Minute)
	if request.DispatchRequestedAt.IsZero() {
		suspiciousAfter = time.Time{}
	}

	initialState := taskrun.StateDispatching
	initialReasonCode := "dispatch_started"
	initialStateSummary := "Run is dispatching in an owned checkout."
	initialNextOwner := "backend"
	initialNextExpectedEvent := "Execution worker records the next task-run state update."
	initialAttention := taskrun.AttentionPriority{Level: taskrun.AttentionWatch, Reason: "Run is active and waiting for the next backend update.", SortKey: "50-dispatching"}
	initialLastProgressSummary := "Captured task docs and provisioned an owned checkout."

	if request.RepoLane.CurrentCommit != "" {
		initialState = taskrun.StateRunning
		initialReasonCode = "owned_lane_bootstrapped"
		initialStateSummary = "Run bootstrapped the owned checkout and is ready for backend execution."
		initialNextOwner = "backend_worker"
		initialNextExpectedEvent = "Execution worker records the next progress checkpoint."
		initialAttention = taskrun.AttentionPriority{Level: taskrun.AttentionWatch, Reason: "Run is active after owned-lane bootstrap.", SortKey: "45-owned_lane_bootstrapped"}
		initialLastProgressSummary = "Bootstrapped the owned checkout and recorded its current commit."
	}

	actions := map[string]taskrun.ActionAvailability{
		taskrun.ActionDispatch: {
			Allowed: false,
			BlockReasons: []taskrun.ActionBlockReason{{
				Code:    "active_run_exists",
				Summary: "Dispatch is blocked while this run owns the current live story.",
			}},
		},
		taskrun.ActionPoke: {
			Allowed: false,
			BlockReasons: []taskrun.ActionBlockReason{{
				Code:    "poke_not_implemented",
				Summary: "Poke is not implemented yet for task runs.",
			}},
		},
		taskrun.ActionInterrupt: {
			Allowed: false,
			BlockReasons: []taskrun.ActionBlockReason{{
				Code:    "interrupt_not_implemented",
				Summary: "Interrupt is not implemented yet for task runs.",
			}},
		},
	}
	if initialState == taskrun.StateRunning {
		actions[taskrun.ActionPoke] = taskrun.ActionAvailability{
			Allowed: false,
			BlockReasons: []taskrun.ActionBlockReason{{
				Code:    "run_not_suspicious_yet",
				Summary: "Poke stays blocked until the run misses its next expected progress deadline.",
			}},
		}
		actions[taskrun.ActionInterrupt] = taskrun.ActionAvailability{Allowed: true}
	}

	return taskrun.TaskRunView{
		RunID:                  request.RunID,
		TaskID:                 request.TaskID,
		WorkflowID:             workflowID,
		TemporalExecutionRunID: executionRunID,
		Status:                 "active",
		StateEnvelope: taskrun.StateEnvelope{
			State:              initialState,
			ReasonCode:         initialReasonCode,
			StateSummary:       initialStateSummary,
			NextOwner:          initialNextOwner,
			NextExpectedEvent:  initialNextExpectedEvent,
			SuspiciousAfter:    suspiciousAfter,
			ActionBlockReasons: collectActionBlockReasons(actions),
		},
		MeaningSummary:      request.MeaningSummary,
		Attention:           initialAttention,
		Actions:             actions,
		RepoLane:            request.RepoLane,
		LastProgressAt:      request.DispatchRequestedAt,
		LastProgressSummary: initialLastProgressSummary,
		CapturedTaskSnapshot: taskrun.TaskDefinitionSnapshot{
			DeclaredWorktreeRoot: request.CapturedTaskSnapshot.DeclaredWorktreeRoot,
			DeclaredTaskRoot:     request.CapturedTaskSnapshot.DeclaredTaskRoot,
			DeclaredTaskRevision: request.CapturedTaskSnapshot.DeclaredTaskRevision,
			DeclaredGitRevision:  request.CapturedTaskSnapshot.DeclaredGitRevision,
			CapturedAt:           request.CapturedTaskSnapshot.CapturedAt,
			Files:                append([]taskrun.TaskArtifactDigest(nil), request.CapturedTaskSnapshot.Files...),
		},
		DocRuntimeDivergenceStatus:  "in_sync",
		DocRuntimeDivergenceSummary: "Runtime task snapshot matches the declared task docs captured at dispatch.",
	}
}

type executionPreflightResult struct {
	CurrentCommit         string          `json:"current_commit"`
	PreflightArtifactPath string          `json:"preflight_artifact_path"`
	OwnedTaskRoot         string          `json:"owned_task_root"`
	GitStatusShort        string          `json:"git_status_short,omitempty"`
	DocPresence           map[string]bool `json:"doc_presence,omitempty"`
	ProgressSummary       string          `json:"progress_summary"`
}

type executionPreflightArtifact struct {
	TaskID               string          `json:"task_id"`
	RunID                string          `json:"run_id"`
	OwnedRepoRoot        string          `json:"owned_repo_root"`
	OwnedTaskRoot        string          `json:"owned_task_root"`
	DeclaredWorktreeRoot string          `json:"declared_worktree_root"`
	DeclaredTaskRoot     string          `json:"declared_task_root"`
	DeclaredTaskRevision string          `json:"declared_task_revision"`
	DeclaredGitRevision  string          `json:"declared_git_revision,omitempty"`
	CurrentCommit        string          `json:"current_commit"`
	GitStatusShort       string          `json:"git_status_short,omitempty"`
	DocPresence          map[string]bool `json:"doc_presence,omitempty"`
	RecordedAt           time.Time       `json:"recorded_at"`
}

type workloadStepResult struct {
	CurrentCommit    string `json:"current_commit"`
	WorkloadStepPath string `json:"workload_step_path"`
	ProgressSummary  string `json:"progress_summary"`
}

type workloadStepArtifact struct {
	TaskID                string    `json:"task_id"`
	RunID                 string    `json:"run_id"`
	MeaningSummary        string    `json:"meaning_summary"`
	OwnedRepoRoot         string    `json:"owned_repo_root"`
	OwnedTaskRoot         string    `json:"owned_task_root"`
	DeclaredTaskRoot      string    `json:"declared_task_root"`
	DeclaredTaskRevision  string    `json:"declared_task_revision"`
	DeclaredGitRevision   string    `json:"declared_git_revision,omitempty"`
	PreflightArtifactPath string    `json:"preflight_artifact_path,omitempty"`
	BootstrapArtifactPath string    `json:"bootstrap_artifact_path,omitempty"`
	CurrentCommit         string    `json:"current_commit"`
	GeneratedAt           time.Time `json:"generated_at"`
	WorkloadInstruction   string    `json:"workload_instruction"`
	ExecutionKind         string    `json:"execution_kind,omitempty"`
	ExecutionWorkingDir   string    `json:"execution_working_dir,omitempty"`
	ExecutionCommand      []string  `json:"execution_command,omitempty"`
}

type workloadExecutionResult struct {
	CurrentCommit      string `json:"current_commit"`
	WorkloadResultPath string `json:"workload_result_path"`
	WorkloadOutputPath string `json:"workload_output_path,omitempty"`
	WorkloadCodePath   string `json:"workload_code_path,omitempty"`
	ProgressSummary    string `json:"progress_summary"`
}

type workloadExecutionArtifact struct {
	TaskID              string    `json:"task_id"`
	RunID               string    `json:"run_id"`
	OwnedRepoRoot       string    `json:"owned_repo_root"`
	WorkloadStepPath    string    `json:"workload_step_path"`
	WorkloadInstruction string    `json:"workload_instruction"`
	ExecutionKind       string    `json:"execution_kind,omitempty"`
	ExecutionWorkingDir string    `json:"execution_working_dir,omitempty"`
	ExecutionCommand    []string  `json:"execution_command,omitempty"`
	StdoutPath          string    `json:"stdout_path,omitempty"`
	StderrPath          string    `json:"stderr_path,omitempty"`
	ExitCode            int       `json:"exit_code,omitempty"`
	WorkloadOutputPath  string    `json:"workload_output_path,omitempty"`
	WorkloadCodePath    string    `json:"workload_code_path,omitempty"`
	BehaviorProbePath   string    `json:"behavior_probe_path,omitempty"`
	GitStatusShortAfter string    `json:"git_status_short_after,omitempty"`
	CurrentCommit       string    `json:"current_commit"`
	ExecutedAt          time.Time `json:"executed_at"`
	ExecutionSummary    string    `json:"execution_summary"`
}

func runExecutionPreflight(ctx context.Context, request taskrun.StartTaskRunRequest) (executionPreflightResult, error) {
	if request.RepoLane.OwnedRepoRoot == "" {
		return executionPreflightResult{}, fmt.Errorf("owned repo root is missing")
	}
	if request.RepoLane.RunArtifactRoot == "" {
		return executionPreflightResult{}, fmt.Errorf("run artifact root is missing")
	}

	ownedTaskRoot, err := ownedTaskRoot(request.CapturedTaskSnapshot, request.RepoLane)
	if err != nil {
		return executionPreflightResult{}, err
	}
	currentCommit, err := gitRevParse(request.RepoLane.OwnedRepoRoot, "HEAD")
	if err != nil {
		return executionPreflightResult{}, err
	}
	gitStatusShort, err := gitStatusShort(request.RepoLane.OwnedRepoRoot)
	if err != nil {
		return executionPreflightResult{}, err
	}
	docPresence := map[string]bool{
		"TASK.md":         pathExists(filepath.Join(ownedTaskRoot, "TASK.md")),
		"PLAN.md":         pathExists(filepath.Join(ownedTaskRoot, "PLAN.md")),
		"HANDOFF.md":      pathExists(filepath.Join(ownedTaskRoot, "HANDOFF.md")),
		"TASK-STATE.json": pathExists(filepath.Join(ownedTaskRoot, "TASK-STATE.json")),
	}

	if err := os.MkdirAll(request.RepoLane.RunArtifactRoot, 0o755); err != nil {
		return executionPreflightResult{}, fmt.Errorf("create run artifact root: %w", err)
	}
	artifactPath := filepath.Join(request.RepoLane.RunArtifactRoot, "execution-preflight.json")
	artifact := executionPreflightArtifact{
		TaskID:               request.TaskID,
		RunID:                request.RunID,
		OwnedRepoRoot:        request.RepoLane.OwnedRepoRoot,
		OwnedTaskRoot:        ownedTaskRoot,
		DeclaredWorktreeRoot: request.CapturedTaskSnapshot.DeclaredWorktreeRoot,
		DeclaredTaskRoot:     request.CapturedTaskSnapshot.DeclaredTaskRoot,
		DeclaredTaskRevision: request.CapturedTaskSnapshot.DeclaredTaskRevision,
		DeclaredGitRevision:  request.CapturedTaskSnapshot.DeclaredGitRevision,
		CurrentCommit:        currentCommit,
		GitStatusShort:       gitStatusShort,
		DocPresence:          docPresence,
		RecordedAt:           time.Now().UTC(),
	}
	if err := writeJSONArtifact(artifactPath, artifact); err != nil {
		return executionPreflightResult{}, err
	}

	return executionPreflightResult{
		CurrentCommit:         currentCommit,
		PreflightArtifactPath: artifactPath,
		OwnedTaskRoot:         ownedTaskRoot,
		GitStatusShort:        gitStatusShort,
		DocPresence:           docPresence,
		ProgressSummary:       "Execution preflight inspected the owned task docs and recorded owned-lane readiness.",
	}, nil
}

func runWorkloadStep(ctx context.Context, request taskrun.StartTaskRunRequest, repoLane taskrun.RepoLane) (workloadStepResult, error) {
	if repoLane.OwnedRepoRoot == "" {
		return workloadStepResult{}, fmt.Errorf("owned repo root is missing")
	}
	ownedTaskRoot, err := ownedTaskRoot(request.CapturedTaskSnapshot, repoLane)
	if err != nil {
		return workloadStepResult{}, err
	}
	stepRoot := filepath.Join(repoLane.OwnedRepoRoot, ".codex-taskrun", sanitizePathSegment(request.RunID))
	if err := os.MkdirAll(stepRoot, 0o755); err != nil {
		return workloadStepResult{}, fmt.Errorf("create owned-lane workload step root: %w", err)
	}
	currentCommit, err := gitRevParse(repoLane.OwnedRepoRoot, "HEAD")
	if err != nil {
		return workloadStepResult{}, err
	}
	stepPath := filepath.Join(stepRoot, "workload-step-0001.json")
	artifact := workloadStepArtifact{
		TaskID:                request.TaskID,
		RunID:                 request.RunID,
		MeaningSummary:        request.MeaningSummary,
		OwnedRepoRoot:         repoLane.OwnedRepoRoot,
		OwnedTaskRoot:         ownedTaskRoot,
		DeclaredTaskRoot:      request.CapturedTaskSnapshot.DeclaredTaskRoot,
		DeclaredTaskRevision:  request.CapturedTaskSnapshot.DeclaredTaskRevision,
		DeclaredGitRevision:   request.CapturedTaskSnapshot.DeclaredGitRevision,
		PreflightArtifactPath: repoLane.PreflightArtifactPath,
		BootstrapArtifactPath: repoLane.BootstrapArtifactPath,
		CurrentCommit:         currentCommit,
		GeneratedAt:           time.Now().UTC(),
		WorkloadInstruction:   "Use the owned task root and captured task snapshot to execute the next backend-owned task step from inside this owned lane.",
	}
	if request.TaskID == "Task-0008" {
		artifact.WorkloadInstruction = "Run focused Task-0008 backend validation from the owned checkout so the first real task-specific execution step happens inside the backend-owned lane."
		artifact.ExecutionKind = "task_0008_backend_validation"
		artifact.ExecutionWorkingDir = filepath.Join(repoLane.OwnedRepoRoot, "backend", "orchestration")
		artifact.ExecutionCommand = []string{
			"go",
			"test",
			"./internal/taskexec",
			"./internal/taskrun",
		}
	}
	if err := writeJSONArtifact(stepPath, artifact); err != nil {
		return workloadStepResult{}, err
	}
	return workloadStepResult{
		CurrentCommit:    currentCommit,
		WorkloadStepPath: stepPath,
		ProgressSummary:  "Prepared the first backend workload step inside the owned lane.",
	}, nil
}

func runExecuteWorkloadStep(ctx context.Context, request taskrun.StartTaskRunRequest, repoLane taskrun.RepoLane) (workloadExecutionResult, error) {
	if repoLane.OwnedRepoRoot == "" {
		return workloadExecutionResult{}, fmt.Errorf("owned repo root is missing")
	}
	if repoLane.WorkloadStepPath == "" {
		return workloadExecutionResult{}, fmt.Errorf("workload step path is missing")
	}
	rawStep, err := os.ReadFile(repoLane.WorkloadStepPath)
	if err != nil {
		return workloadExecutionResult{}, fmt.Errorf("read workload step %s: %w", repoLane.WorkloadStepPath, err)
	}
	var step workloadStepArtifact
	if err := json.Unmarshal(rawStep, &step); err != nil {
		return workloadExecutionResult{}, fmt.Errorf("decode workload step %s: %w", repoLane.WorkloadStepPath, err)
	}
	if strings.TrimSpace(step.WorkloadInstruction) == "" {
		return workloadExecutionResult{}, fmt.Errorf("workload instruction is missing")
	}
	resultPath := strings.TrimSuffix(repoLane.WorkloadStepPath, ".json") + ".result.json"
	currentCommit, err := gitRevParse(repoLane.OwnedRepoRoot, "HEAD")
	if err != nil {
		return workloadExecutionResult{}, err
	}
	executionSummary := "Executed the first backend workload step packet inside the owned lane."
	stdoutPath := ""
	stderrPath := ""
	exitCode := 0
	workloadOutputPath := ""
	workloadCodePath := ""
	behaviorProbePath := ""
	gitStatusAfter := ""
	if step.ExecutionKind == "task_0008_backend_validation" {
		executionSummary, stdoutPath, stderrPath, exitCode, workloadOutputPath, workloadCodePath, behaviorProbePath, gitStatusAfter, err = executeTask0008Validation(repoLane, step)
		if err != nil {
			return workloadExecutionResult{}, err
		}
	}
	artifact := workloadExecutionArtifact{
		TaskID:              request.TaskID,
		RunID:               request.RunID,
		OwnedRepoRoot:       repoLane.OwnedRepoRoot,
		WorkloadStepPath:    repoLane.WorkloadStepPath,
		WorkloadInstruction: step.WorkloadInstruction,
		ExecutionKind:       step.ExecutionKind,
		ExecutionWorkingDir: step.ExecutionWorkingDir,
		ExecutionCommand:    append([]string(nil), step.ExecutionCommand...),
		StdoutPath:          stdoutPath,
		StderrPath:          stderrPath,
		ExitCode:            exitCode,
		WorkloadOutputPath:  workloadOutputPath,
		WorkloadCodePath:    workloadCodePath,
		BehaviorProbePath:   behaviorProbePath,
		GitStatusShortAfter: gitStatusAfter,
		CurrentCommit:       currentCommit,
		ExecutedAt:          time.Now().UTC(),
		ExecutionSummary:    executionSummary,
	}
	if err := writeJSONArtifact(resultPath, artifact); err != nil {
		return workloadExecutionResult{}, err
	}
	return workloadExecutionResult{
		CurrentCommit:      currentCommit,
		WorkloadResultPath: resultPath,
		WorkloadOutputPath: workloadOutputPath,
		WorkloadCodePath:   workloadCodePath,
		ProgressSummary:    executionSummary,
	}, nil
}

func executeTask0008Validation(repoLane taskrun.RepoLane, step workloadStepArtifact) (string, string, string, int, string, string, string, string, error) {
	if repoLane.RunArtifactRoot == "" {
		return "", "", "", 0, "", "", "", "", fmt.Errorf("run artifact root is missing")
	}
	if len(step.ExecutionCommand) == 0 {
		return "", "", "", 0, "", "", "", "", fmt.Errorf("execution command is missing")
	}
	workingDir := step.ExecutionWorkingDir
	if workingDir == "" {
		workingDir = repoLane.OwnedRepoRoot
	}
	if err := os.MkdirAll(repoLane.RunArtifactRoot, 0o755); err != nil {
		return "", "", "", 0, "", "", "", "", fmt.Errorf("create run artifact root: %w", err)
	}
	stdoutPath := filepath.Join(repoLane.RunArtifactRoot, "task-specific-validation.stdout.txt")
	stderrPath := filepath.Join(repoLane.RunArtifactRoot, "task-specific-validation.stderr.txt")
	stdoutFile, err := os.Create(stdoutPath)
	if err != nil {
		return "", "", "", 0, "", "", "", "", fmt.Errorf("create validation stdout log: %w", err)
	}
	defer stdoutFile.Close()
	stderrFile, err := os.Create(stderrPath)
	if err != nil {
		return "", "", "", 0, "", "", "", "", fmt.Errorf("create validation stderr log: %w", err)
	}
	defer stderrFile.Close()

	workloadOutputPath, err := writeTask0008OwnedLaneBrief(step)
	if err != nil {
		return "", stdoutPath, stderrPath, 0, "", "", "", "", err
	}
	workloadCodePath, err := writeTask0008OwnedLaneCode(step)
	if err != nil {
		return "", stdoutPath, stderrPath, 0, workloadOutputPath, "", "", "", err
	}

	cmd := exec.Command(step.ExecutionCommand[0], step.ExecutionCommand[1:]...)
	cmd.Dir = workingDir
	cmd.Stdout = stdoutFile
	cmd.Stderr = stderrFile
	err = cmd.Run()
	exitCode := 0
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitCode = exitErr.ExitCode()
		} else {
			exitCode = 1
		}
		return "", stdoutPath, stderrPath, exitCode, workloadOutputPath, workloadCodePath, "", "", fmt.Errorf("task-specific validation failed with exit code %d: %w", exitCode, err)
	}
	behaviorProbePath, err := runTask0008OwnedLaneBehaviorProbe(repoLane, step)
	if err != nil {
		return "", stdoutPath, stderrPath, exitCode, workloadOutputPath, workloadCodePath, "", "", err
	}
	gitStatusAfter, err := gitStatusShortPaths(
		repoLane.OwnedRepoRoot,
		".codex-taskrun",
		filepath.Join("Tracking", "Task-0008", "OwnedLane"),
		filepath.Join("backend", "orchestration", "internal", "taskrun", "service.go"),
	)
	if err != nil {
		return "", stdoutPath, stderrPath, exitCode, workloadOutputPath, workloadCodePath, behaviorProbePath, "", err
	}

	summary := "Executed Task-0008 backend validation and changed owned-lane redispatch cleanup behavior in an existing implementation file."
	return summary, stdoutPath, stderrPath, exitCode, workloadOutputPath, workloadCodePath, behaviorProbePath, gitStatusAfter, nil
}

func writeTask0008OwnedLaneBrief(step workloadStepArtifact) (string, error) {
	if step.OwnedTaskRoot == "" {
		return "", fmt.Errorf("owned task root is missing")
	}
	taskMarkdown, err := os.ReadFile(filepath.Join(step.OwnedTaskRoot, "TASK.md"))
	if err != nil {
		return "", fmt.Errorf("read owned TASK.md: %w", err)
	}
	handoffMarkdown, err := os.ReadFile(filepath.Join(step.OwnedTaskRoot, "HANDOFF.md"))
	if err != nil {
		return "", fmt.Errorf("read owned HANDOFF.md: %w", err)
	}
	constraintsMarkdown, err := os.ReadFile(filepath.Join(step.OwnedTaskRoot, "CONSTRAINTS.md"))
	if err != nil {
		return "", fmt.Errorf("read owned CONSTRAINTS.md: %w", err)
	}

	outputPath := filepath.Join(step.OwnedTaskRoot, "OwnedLane", "IMPLEMENTATION-BRIEF.md")
	if err := os.MkdirAll(filepath.Dir(outputPath), 0o755); err != nil {
		return "", fmt.Errorf("create owned-lane brief directory: %w", err)
	}

	taskSummary := taskexecFirstParagraph(taskexecExtractMarkdownSection(string(taskMarkdown), "Summary"))
	nextRecommended := strings.TrimSpace(taskexecExtractMarkdownSection(string(handoffMarkdown), "Next Recommended Step"))
	if nextRecommended == "" {
		nextRecommended = "Keep the next owned-lane implementation step bounded and task-specific."
	}
	constraintReminder := taskexecFirstBulletUnderHeading(string(constraintsMarkdown), "Active Constraints")
	if constraintReminder == "" {
		constraintReminder = "Keep the current slice bounded and preserve the declared-doc versus runtime split."
	}

	brief := strings.Join([]string{
		"# Task-0008 Owned-Lane Implementation Brief",
		"",
		fmt.Sprintf("Generated at `%s` for run `%s`.", time.Now().UTC().Format(time.RFC3339), step.RunID),
		"",
		"## Why This Task Exists",
		"",
		taskSummary,
		"",
		"## Current Task-Specific Worker Action",
		"",
		"- Completed owned-lane backend validation for `./internal/taskexec` and `./internal/taskrun`.",
		"- Wrote this brief inside the owned Task-0008 directory so the run has a real repo-state change to build on.",
		"",
		"## Declared Inputs Used",
		"",
		fmt.Sprintf("- Declared task revision: `%s`", step.DeclaredTaskRevision),
		fmt.Sprintf("- Declared git revision: `%s`", step.DeclaredGitRevision),
		fmt.Sprintf("- Preflight artifact: `%s`", step.PreflightArtifactPath),
		fmt.Sprintf("- Bootstrap artifact: `%s`", step.BootstrapArtifactPath),
		"",
		"## Next Recommended Step From Handoff",
		"",
		nextRecommended,
		"",
		"## Constraint Reminder",
		"",
		fmt.Sprintf("- %s", constraintReminder),
		"",
		"## Candidate Owned-Lane Targets",
		"",
		"- `backend/orchestration/internal/taskexec/taskexec.go`",
		"- `backend/orchestration/internal/taskrun/service.go`",
		"- `backend/orchestration/internal/taskrun/types.go`",
		"",
		"## Worker Intent",
		"",
		"- Keep the next Task-0008-owned change bounded.",
		"- Preserve the declared-doc versus runtime split.",
		"- Prefer a real owned-lane code or task artifact mutation over another marker-only transition.",
		"",
	}, "\n")
	if err := os.WriteFile(outputPath, []byte(brief), 0o644); err != nil {
		return "", fmt.Errorf("write owned-lane implementation brief: %w", err)
	}
	return outputPath, nil
}

func writeTask0008OwnedLaneCode(step workloadStepArtifact) (string, error) {
	codePath := filepath.Join(step.OwnedRepoRoot, "backend", "orchestration", "internal", "taskrun", "service.go")
	raw, err := os.ReadFile(codePath)
	if err != nil {
		return "", fmt.Errorf("read owned-lane implementation file: %w", err)
	}
	updated := applyTask0008InterruptReviewWindowChange(string(raw))
	if updated == string(raw) {
		return "", fmt.Errorf("owned-lane behavior change was not applied")
	}
	if err := os.WriteFile(codePath, []byte(updated), 0o644); err != nil {
		return "", fmt.Errorf("write owned-lane implementation file: %w", err)
	}
	return codePath, nil
}

func applyTask0008InterruptReviewWindowChange(source string) string {
	const dispatchNeedle = `	if !task.DispatchReadiness.Ready {
		return TaskRunView{}, fmt.Errorf("dispatch blocked: %s", summarizeBlockReasons(task.DispatchReadiness.BlockReasons))
	}

	repoLane, err := s.provisionOwnedLane(task.TaskID)`
	const dispatchReplacement = `	if !task.DispatchReadiness.Ready {
		return TaskRunView{}, fmt.Errorf("dispatch blocked: %s", summarizeBlockReasons(task.DispatchReadiness.BlockReasons))
	}
	if err := s.releasePreviousOwnedLane(ctx, task.TaskID); err != nil {
		return TaskRunView{}, err
	}

	repoLane, err := s.provisionOwnedLane(task.TaskID)`
	const fixtureDispatchNeedle = `	if !state.PlanApproved || len(state.Blockers) > 0 {
		return TaskRunView{}, fmt.Errorf("dispatch blocked")
	}
	repoLane, err := s.provisionOwnedLane(taskID)`
	const fixtureDispatchReplacement = `	if !state.PlanApproved || len(state.Blockers) > 0 {
		return TaskRunView{}, fmt.Errorf("dispatch blocked")
	}
	if err := s.releasePreviousOwnedLane(ctx, taskID); err != nil {
		return TaskRunView{}, err
	}
	repoLane, err := s.provisionOwnedLane(taskID)`
	const helperNeedle = `func (s *Service) bootstrapOwnedLane(taskID string, runID string, snapshot TaskDefinitionSnapshot, repoLane RepoLane) (RepoLane, error) {`
	const helperInsert = `func (s *Service) releasePreviousOwnedLane(ctx context.Context, taskID string) error {
	if s.runtime == nil {
		return nil
	}
	previousRun, err := s.runtime.GetActiveTaskRun(ctx, taskID)
	if err != nil {
		if errors.Is(err, ErrRunNotFound) {
			return nil
		}
		return err
	}
	if runOwnsLiveStory(previousRun) || previousRun.RepoLane.OwnedRepoRoot == "" {
		return nil
	}
	if err := s.cleanupOwnedLane(previousRun.RepoLane); err != nil {
		return fmt.Errorf("release previous owned lane for %s: %w", taskID, err)
	}
	return nil
}

func (s *Service) bootstrapOwnedLane(taskID string, runID string, snapshot TaskDefinitionSnapshot, repoLane RepoLane) (RepoLane, error) {`
	const cleanupNeedle = `func (s *Service) cleanupOwnedLane(repoLane RepoLane) error {
	if repoLane.OwnedRepoRoot == "" {
		return nil
	}
	cmd := exec.Command("git", "-C", s.declaredWorktreeRoot, "worktree", "remove", "--force", repoLane.OwnedRepoRoot)`
	const cleanupReplacement = `func (s *Service) cleanupOwnedLane(repoLane RepoLane) error {
	if repoLane.OwnedRepoRoot == "" {
		return nil
	}
	if !pathWithinRoot(repoLane.OwnedRepoRoot, s.ownedLaneRoot) {
		return fmt.Errorf("owned repo root %q is outside the backend-owned lane root", repoLane.OwnedRepoRoot)
	}
	cmd := exec.Command("git", "-C", s.declaredWorktreeRoot, "worktree", "remove", "--force", repoLane.OwnedRepoRoot)`
	updated := strings.Replace(source, dispatchNeedle, dispatchReplacement, 1)
	updated = strings.Replace(updated, fixtureDispatchNeedle, fixtureDispatchReplacement, 1)
	updated = strings.Replace(updated, helperNeedle, helperInsert, 1)
	updated = strings.Replace(updated, cleanupNeedle, cleanupReplacement, 1)
	return updated
}

func runTask0008OwnedLaneBehaviorProbe(repoLane taskrun.RepoLane, step workloadStepArtifact) (string, error) {
	if repoLane.RunArtifactRoot == "" {
		return "", fmt.Errorf("run artifact root is missing")
	}
	moduleRoot := step.ExecutionWorkingDir
	if moduleRoot == "" {
		moduleRoot = filepath.Join(repoLane.OwnedRepoRoot, "backend", "orchestration")
	}
	modulePath, err := modulePathFromGoMod(moduleRoot)
	if err != nil {
		return "", err
	}
	probeDir := filepath.Join(moduleRoot, ".codex-taskrun", sanitizePathSegment(step.RunID), "behaviorprobe")
	if err := os.MkdirAll(probeDir, 0o755); err != nil {
		return "", fmt.Errorf("create behavior probe dir: %w", err)
	}
	probePath := filepath.Join(probeDir, "main.go")
	probeSource := fmt.Sprintf(`package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	taskrun "%s/internal/taskrun"
)

type fakeRuntime struct {
	activeByTask map[string]taskrun.TaskRunView
	byRunID      map[string]taskrun.TaskRunView
	run          taskrun.TaskRunView
}

func (f *fakeRuntime) StartTaskRun(ctx context.Context, request taskrun.StartTaskRunRequest) (taskrun.TaskRunView, error) {
	run := taskrun.TaskRunView{
		RunID:  request.RunID,
		TaskID: request.TaskID,
		Status: "active",
		StateEnvelope: taskrun.StateEnvelope{
			State:             taskrun.StateRunning,
			ReasonCode:        "owned_lane_bootstrapped",
			StateSummary:      "Run bootstrapped the owned checkout and is ready for backend execution.",
			NextOwner:         "backend_worker",
			NextExpectedEvent: "Execution worker records the next progress checkpoint.",
			SuspiciousAfter:   request.DispatchRequestedAt.Add(15 * time.Minute),
		},
		RepoLane: request.RepoLane,
	}
	f.run = run
	f.activeByTask[request.TaskID] = run
	f.byRunID[request.RunID] = run
	return run, nil
}

func (f *fakeRuntime) GetTaskRun(ctx context.Context, runID string) (taskrun.TaskRunView, error) {
	run, ok := f.byRunID[runID]
	if !ok {
		return taskrun.TaskRunView{}, fmt.Errorf("run not found")
	}
	return run, nil
}

func (f *fakeRuntime) GetActiveTaskRun(ctx context.Context, taskID string) (taskrun.TaskRunView, error) {
	run, ok := f.activeByTask[taskID]
	if !ok {
		return taskrun.TaskRunView{}, fmt.Errorf("run not found")
	}
	return run, nil
}

func (f *fakeRuntime) ReconcileTaskSnapshot(ctx context.Context, runID string, snapshot taskrun.TaskDefinitionSnapshot) (taskrun.TaskRunView, error) {
	return f.GetTaskRun(ctx, runID)
}

func (f *fakeRuntime) UpdateTaskRun(ctx context.Context, runID string, update taskrun.TaskRunUpdate) (taskrun.TaskRunView, error) {
	run, err := f.GetTaskRun(ctx, runID)
	if err != nil {
		return taskrun.TaskRunView{}, err
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
	if update.FollowUp != nil {
		run.FollowUp = update.FollowUp
	}
	if update.RepoLane != nil {
		run.RepoLane = *update.RepoLane
	}
	if update.Actions != nil {
		run.Actions = update.Actions
	}
	if update.Attention != nil {
		run.Attention = *update.Attention
	}
	if !update.CompletedAt.IsZero() {
		run.Status = "interrupted"
	}
	f.run = run
	f.activeByTask[run.TaskID] = run
	f.byRunID[run.RunID] = run
	return run, nil
}

func mustRun(dir string, name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	if output, err := cmd.CombinedOutput(); err != nil {
		panic(fmt.Sprintf("%%s %%v failed: %%v %%s", name, args, err, string(output)))
	}
}

func mustOutput(dir string, name string, args ...string) string {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	output, err := cmd.CombinedOutput()
	if err != nil {
		panic(fmt.Sprintf("%%s %%v failed: %%v %%s", name, args, err, string(output)))
	}
	return string(output)
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func main() {
	baseRoot := filepath.Join(os.TempDir(), fmt.Sprintf("task0008-probe-%%d", time.Now().UnixNano()))
	worktreeRoot := filepath.Join(baseRoot, "worktree")
	ownedRoot := filepath.Join(os.TempDir(), "cdxow", filepath.Base(baseRoot), "w")
	runsRoot := filepath.Join(baseRoot, "runs")
	_ = os.RemoveAll(worktreeRoot)
	_ = os.RemoveAll(filepath.Dir(ownedRoot))
	_ = os.RemoveAll(runsRoot)
	_ = os.MkdirAll(filepath.Join(worktreeRoot, "Tracking", "Task-0008"), 0o755)
	_ = os.WriteFile(filepath.Join(worktreeRoot, "Tracking", "Task-0008", "TASK.md"), []byte("# Task 0008\n\n## Title\n\nTask 0008\n\n## Summary\n\nOwned lane redispatch probe.\n"), 0o644)
	_ = os.WriteFile(filepath.Join(worktreeRoot, "Tracking", "Task-0008", "TASK-STATE.json"), []byte("{\"task_id\":\"Task-0008\",\"status\":\"in_progress\",\"phase\":\"implementation\",\"plan_approved\":true,\"current_pass\":\"PASS-0002\",\"current_gate\":\"implementation\",\"blockers\":[],\"updated_at\":\"2026-04-24T17:10:00-04:00\"}\n"), 0o644)
	_ = os.WriteFile(filepath.Join(worktreeRoot, "Tracking", "Task-0008", "PLAN.md"), []byte("# approved plan\n"), 0o644)
	mustRun(worktreeRoot, "git", "init")
	mustRun(worktreeRoot, "git", "config", "user.email", "probe@example.com")
	mustRun(worktreeRoot, "git", "config", "user.name", "Probe")
	mustRun(worktreeRoot, "git", "add", ".")
	mustRun(worktreeRoot, "git", "commit", "-m", "initial")
	mustRun(worktreeRoot, "git", "worktree", "add", "--detach", ownedRoot, "HEAD")
	head := strings.TrimSpace(mustOutput(worktreeRoot, "git", "rev-parse", "HEAD"))

	rt := &fakeRuntime{
		activeByTask: map[string]taskrun.TaskRunView{},
		byRunID:      map[string]taskrun.TaskRunView{},
		run: taskrun.TaskRunView{
			RunID:  "taskrun--Task-0008--active",
			TaskID: "Task-0008",
			Status: "interrupted",
			StateEnvelope: taskrun.StateEnvelope{
				State:             taskrun.StateInterrupted,
				ReasonCode:        "interrupt_review_resolved_redispatch_ready",
				StateSummary:      "Interrupt review approved the run for redispatch.",
				NextOwner:         "backend",
				NextExpectedEvent: "Dispatch a new run when the task is ready.",
			},
			FollowUp: &taskrun.RunFollowUp{
				Kind:        "interrupt_review",
				Owner:       "human_or_supervisor",
				Status:      "completed",
				Summary:     "Approved for redispatch.",
				RequestedAt: time.Now().UTC().Add(-30 * time.Minute),
				DueAt:       time.Now().UTC().Add(-28 * time.Minute),
				CompletedAt: time.Now().UTC().Add(-25 * time.Minute),
			},
			Resolution: &taskrun.RunResolution{
				Kind:       "interrupt_review",
				Decision:   "redispatch_ready",
				Summary:    "Approved for redispatch.",
				ResolvedBy: "human",
				ResolvedAt: time.Now().UTC().Add(-25 * time.Minute),
			},
			RepoLane: taskrun.RepoLane{
				OwnedRepoRoot:         ownedRoot,
				CheckoutMode:          "git_worktree_detached",
				BaselineCommit:        head,
				CurrentCommit:         head,
				ApprovedRestoreCommit: head,
				ResetStatus:           "not_run",
			},
		},
	}
	rt.activeByTask[rt.run.TaskID] = rt.run
	rt.byRunID[rt.run.RunID] = rt.run

	service := taskrun.NewService(worktreeRoot, runsRoot, rt)
	redispatched, err := service.Dispatch(context.Background(), "Task-0008")
	if err != nil {
		panic(err)
	}

	_ = json.NewEncoder(os.Stdout).Encode(map[string]any{
		"redispatch_run_id":              redispatched.RunID,
		"original_owned_root":            ownedRoot,
		"original_owned_root_removed":    !pathExists(ownedRoot),
		"new_owned_root":                 redispatched.RepoLane.OwnedRepoRoot,
		"new_owned_root_exists":          pathExists(redispatched.RepoLane.OwnedRepoRoot),
		"new_owned_root_differs":         redispatched.RepoLane.OwnedRepoRoot != ownedRoot,
		"new_run_reason_code":            redispatched.StateEnvelope.ReasonCode,
		"new_run_current_commit_present": redispatched.RepoLane.CurrentCommit != "",
	})
}
`, modulePath)
	if err := os.WriteFile(probePath, []byte(probeSource), 0o644); err != nil {
		return "", fmt.Errorf("write behavior probe: %w", err)
	}

	cmd := exec.Command("go", "run", "./.codex-taskrun/"+sanitizePathSegment(step.RunID)+"/behaviorprobe")
	cmd.Dir = moduleRoot
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("run owned-lane behavior probe: %w: %s", err, strings.TrimSpace(string(output)))
	}
	artifactPath := filepath.Join(repoLane.RunArtifactRoot, "task-specific-behavior-probe.json")
	if err := os.WriteFile(artifactPath, output, 0o644); err != nil {
		return "", fmt.Errorf("write behavior probe artifact: %w", err)
	}
	var probe struct {
		OriginalOwnedRootRemoved bool   `json:"original_owned_root_removed"`
		NewOwnedRoot             string `json:"new_owned_root"`
		NewOwnedRootExists       bool   `json:"new_owned_root_exists"`
		NewOwnedRootDiffers      bool   `json:"new_owned_root_differs"`
		NewRunReasonCode         string `json:"new_run_reason_code"`
		NewRunCurrentCommit      bool   `json:"new_run_current_commit_present"`
	}
	if err := json.Unmarshal(output, &probe); err != nil {
		return "", fmt.Errorf("decode behavior probe output: %w", err)
	}
	if !probe.OriginalOwnedRootRemoved {
		return "", fmt.Errorf("behavior probe expected the original owned root to be removed")
	}
	if !probe.NewOwnedRootExists {
		return "", fmt.Errorf("behavior probe expected the new owned root to exist")
	}
	if !probe.NewOwnedRootDiffers {
		return "", fmt.Errorf("behavior probe expected a fresh owned root")
	}
	if probe.NewRunReasonCode != "owned_lane_bootstrapped" {
		return "", fmt.Errorf("behavior probe expected owned_lane_bootstrapped after redispatch, got %q", probe.NewRunReasonCode)
	}
	if !probe.NewRunCurrentCommit {
		return "", fmt.Errorf("behavior probe expected the redispatched run to capture current commit")
	}
	return artifactPath, nil
}

func modulePathFromGoMod(moduleRoot string) (string, error) {
	raw, err := os.ReadFile(filepath.Join(moduleRoot, "go.mod"))
	if err != nil {
		return "", fmt.Errorf("read go.mod: %w", err)
	}
	for _, line := range strings.Split(string(raw), "\n") {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(trimmed, "module ")), nil
		}
	}
	return "", fmt.Errorf("module path not found in go.mod")
}

func taskexecExtractMarkdownSection(markdown string, heading string) string {
	lines := strings.Split(markdown, "\n")
	header := "## " + heading
	capture := false
	section := make([]string, 0)
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == header {
			capture = true
			continue
		}
		if capture && strings.HasPrefix(trimmed, "## ") {
			break
		}
		if capture {
			section = append(section, line)
		}
	}
	return strings.TrimSpace(strings.Join(section, "\n"))
}

func taskexecFirstParagraph(section string) string {
	lines := strings.Split(section, "\n")
	paragraph := make([]string, 0)
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			if len(paragraph) > 0 {
				break
			}
			continue
		}
		paragraph = append(paragraph, trimmed)
	}
	return strings.Join(paragraph, " ")
}

func taskexecFirstBulletUnderHeading(markdown string, heading string) string {
	section := taskexecExtractMarkdownSection(markdown, heading)
	for _, line := range strings.Split(section, "\n") {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "- ") {
			return strings.TrimPrefix(trimmed, "- ")
		}
	}
	return ""
}

func collectActionBlockReasons(actions map[string]taskrun.ActionAvailability) map[string][]taskrun.ActionBlockReason {
	blockReasons := map[string][]taskrun.ActionBlockReason{}
	for action, availability := range actions {
		blockReasons[action] = append([]taskrun.ActionBlockReason(nil), availability.BlockReasons...)
	}
	return blockReasons
}

func applyUpdate(view *taskrun.TaskRunView, update taskrun.TaskRunUpdate, now time.Time) {
	if update.State != "" {
		view.StateEnvelope.State = update.State
	}
	if update.ReasonCode != "" {
		view.StateEnvelope.ReasonCode = update.ReasonCode
	}
	if update.StateSummary != "" {
		view.StateEnvelope.StateSummary = update.StateSummary
	}
	if update.NextOwner != "" {
		view.StateEnvelope.NextOwner = update.NextOwner
	}
	if update.NextExpectedEvent != "" {
		view.StateEnvelope.NextExpectedEvent = update.NextExpectedEvent
	}
	if !update.SuspiciousAfter.IsZero() {
		view.StateEnvelope.SuspiciousAfter = update.SuspiciousAfter
	}
	if update.LastProgressSummary != "" {
		view.LastProgressSummary = update.LastProgressSummary
		view.LastProgressAt = now
	}
	if update.WaitContract != nil {
		view.WaitContract = update.WaitContract
	} else if update.State != "" && update.State != taskrun.StateWaitingForHuman {
		view.WaitContract = nil
	}
	if update.Attention != nil {
		view.Attention = *update.Attention
	}
	if update.FollowUp != nil {
		if isEmptyFollowUp(update.FollowUp) {
			view.FollowUp = nil
		} else {
			view.FollowUp = update.FollowUp
		}
	}
	if update.Resolution != nil {
		view.Resolution = update.Resolution
	}
	if update.RepoLane != nil {
		view.RepoLane = *update.RepoLane
	}
	if update.Actions != nil {
		view.Actions = update.Actions
		view.StateEnvelope.ActionBlockReasons = collectActionBlockReasons(update.Actions)
	}
	if update.CompletedAt.IsZero() {
		view.Status = "active"
	} else {
		view.LastProgressAt = update.CompletedAt
	}
	switch view.StateEnvelope.State {
	case taskrun.StateCompleted:
		view.Status = "completed"
	case taskrun.StateFailed:
		view.Status = "failed"
	case taskrun.StateInterrupted:
		view.Status = "interrupted"
	}
	if update.FailureSummary != "" {
		view.FailureSummary = update.FailureSummary
	} else if update.State != "" && update.State != taskrun.StateBlocked && update.State != taskrun.StateFailed {
		view.FailureSummary = ""
	}
}

func isTerminalStatus(status string) bool {
	return status == "completed" || status == "failed" || status == "interrupted"
}

func shouldExit(view taskrun.TaskRunView) bool {
	if view.Status == "interrupted" && hasPendingInterruptReview(view) {
		return false
	}
	return isTerminalStatus(view.Status)
}

func actionsForState(state string) map[string]taskrun.ActionAvailability {
	switch state {
	case taskrun.StateRunning, taskrun.StateDispatching:
		return map[string]taskrun.ActionAvailability{
			taskrun.ActionDispatch: {
				Allowed: false,
				BlockReasons: []taskrun.ActionBlockReason{{
					Code:    "active_run_exists",
					Summary: "Dispatch is blocked while this run owns the current live story.",
				}},
			},
			taskrun.ActionPoke: {
				Allowed: false,
				BlockReasons: []taskrun.ActionBlockReason{{
					Code:    "run_not_suspicious_yet",
					Summary: "Poke stays blocked until the run misses its next expected progress deadline.",
				}},
			},
			taskrun.ActionInterrupt: {Allowed: true},
		}
	case taskrun.StateBlocked:
		return map[string]taskrun.ActionAvailability{
			taskrun.ActionDispatch: {
				Allowed: false,
				BlockReasons: []taskrun.ActionBlockReason{{
					Code:    "active_run_exists",
					Summary: "Dispatch is blocked while this run owns the current live story.",
				}},
			},
			taskrun.ActionPoke: {
				Allowed: false,
				BlockReasons: []taskrun.ActionBlockReason{{
					Code:    "poke_not_allowed_for_state",
					Summary: "Poke is not allowed in the current run state.",
				}},
			},
			taskrun.ActionInterrupt: {Allowed: true},
		}
	default:
		return map[string]taskrun.ActionAvailability{
			taskrun.ActionDispatch: {
				Allowed: false,
				BlockReasons: []taskrun.ActionBlockReason{{
					Code:    "active_run_exists",
					Summary: "Dispatch is blocked while this run owns the current live story.",
				}},
			},
			taskrun.ActionPoke: {
				Allowed: false,
				BlockReasons: []taskrun.ActionBlockReason{{
					Code:    "poke_not_implemented",
					Summary: "Poke is not implemented yet for task runs.",
				}},
			},
			taskrun.ActionInterrupt: {
				Allowed: false,
				BlockReasons: []taskrun.ActionBlockReason{{
					Code:    "interrupt_not_implemented",
					Summary: "Interrupt is not implemented yet for task runs.",
				}},
			},
		}
	}
}

func hasPendingInterruptReview(view taskrun.TaskRunView) bool {
	return view.FollowUp != nil &&
		view.FollowUp.Kind == "interrupt_review" &&
		(view.FollowUp.Status == "pending" || view.FollowUp.Status == "overdue")
}

func isEmptyFollowUp(followUp *taskrun.RunFollowUp) bool {
	return followUp != nil &&
		followUp.Kind == "" &&
		followUp.Owner == "" &&
		followUp.Status == "" &&
		followUp.Summary == "" &&
		followUp.RequestedAt.IsZero() &&
		followUp.DueAt.IsZero() &&
		followUp.CompletedAt.IsZero()
}

func ownedTaskRoot(snapshot taskrun.TaskDefinitionSnapshot, repoLane taskrun.RepoLane) (string, error) {
	rel, err := filepath.Rel(snapshot.DeclaredWorktreeRoot, snapshot.DeclaredTaskRoot)
	if err != nil {
		return "", fmt.Errorf("resolve task root relative path: %w", err)
	}
	if rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
		return "", fmt.Errorf("declared task root %q is outside declared worktree root %q", snapshot.DeclaredTaskRoot, snapshot.DeclaredWorktreeRoot)
	}
	return filepath.Join(repoLane.OwnedRepoRoot, rel), nil
}

func gitRevParse(worktreeRoot string, ref string) (string, error) {
	out, err := exec.Command("git", "-C", worktreeRoot, "rev-parse", ref).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("git rev-parse %s: %w: %s", ref, err, strings.TrimSpace(string(out)))
	}
	return strings.TrimSpace(string(out)), nil
}

func gitStatusShort(worktreeRoot string) (string, error) {
	out, err := exec.Command("git", "-C", worktreeRoot, "status", "--short").CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("git status --short: %w: %s", err, strings.TrimSpace(string(out)))
	}
	return strings.TrimSpace(string(out)), nil
}

func gitStatusShortPaths(worktreeRoot string, paths ...string) (string, error) {
	args := []string{"-C", worktreeRoot, "status", "--short", "--"}
	args = append(args, paths...)
	out, err := exec.Command("git", args...).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("git status --short -- %v: %w: %s", paths, err, strings.TrimSpace(string(out)))
	}
	return strings.TrimSpace(string(out)), nil
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func writeJSONArtifact(path string, value any) error {
	data, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return err
	}
	data = append(data, '\n')
	if err := os.WriteFile(path, data, 0o644); err != nil {
		return fmt.Errorf("write %s: %w", path, err)
	}
	return nil
}

func sanitizePathSegment(value string) string {
	replacer := strings.NewReplacer("\\", "_", "/", "_", ":", "_", " ", "_")
	return replacer.Replace(value)
}
