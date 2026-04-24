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
					reasonCode := "workload_step_executed"
					stateSummary := "Run executed the first backend workload step inside the owned lane."
					nextExpectedEvent := "Execution worker prepares or executes the next workload step."
					attentionReason := "Run executed its first workload step and is ready for the next backend step."
					attentionSortKey := "42-workload_step_executed"
					if request.TaskID == "Task-0008" {
						reasonCode = "task_0008_owned_lane_brief_written"
						stateSummary = "Run completed Task-0008 validation and wrote an owned-lane implementation brief."
						nextExpectedEvent = "Execution worker applies the next Task-0008-specific owned-lane change."
						attentionReason = "Run wrote a Task-0008 implementation brief inside the owned lane and is ready for the next backend step."
						attentionSortKey = "40-task_0008_owned_lane_brief_written"
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
	gitStatusAfter := ""
	if step.ExecutionKind == "task_0008_backend_validation" {
		executionSummary, stdoutPath, stderrPath, exitCode, workloadOutputPath, gitStatusAfter, err = executeTask0008Validation(repoLane, step)
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
		ProgressSummary:    executionSummary,
	}, nil
}

func executeTask0008Validation(repoLane taskrun.RepoLane, step workloadStepArtifact) (string, string, string, int, string, string, error) {
	if repoLane.RunArtifactRoot == "" {
		return "", "", "", 0, "", "", fmt.Errorf("run artifact root is missing")
	}
	if len(step.ExecutionCommand) == 0 {
		return "", "", "", 0, "", "", fmt.Errorf("execution command is missing")
	}
	workingDir := step.ExecutionWorkingDir
	if workingDir == "" {
		workingDir = repoLane.OwnedRepoRoot
	}
	if err := os.MkdirAll(repoLane.RunArtifactRoot, 0o755); err != nil {
		return "", "", "", 0, "", "", fmt.Errorf("create run artifact root: %w", err)
	}
	stdoutPath := filepath.Join(repoLane.RunArtifactRoot, "task-specific-validation.stdout.txt")
	stderrPath := filepath.Join(repoLane.RunArtifactRoot, "task-specific-validation.stderr.txt")
	stdoutFile, err := os.Create(stdoutPath)
	if err != nil {
		return "", "", "", 0, "", "", fmt.Errorf("create validation stdout log: %w", err)
	}
	defer stdoutFile.Close()
	stderrFile, err := os.Create(stderrPath)
	if err != nil {
		return "", "", "", 0, "", "", fmt.Errorf("create validation stderr log: %w", err)
	}
	defer stderrFile.Close()

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
		return "", stdoutPath, stderrPath, exitCode, "", "", fmt.Errorf("task-specific validation failed with exit code %d: %w", exitCode, err)
	}

	workloadOutputPath, err := writeTask0008OwnedLaneBrief(step)
	if err != nil {
		return "", stdoutPath, stderrPath, exitCode, "", "", err
	}
	gitStatusAfter, err := gitStatusShort(repoLane.OwnedRepoRoot)
	if err != nil {
		return "", stdoutPath, stderrPath, exitCode, workloadOutputPath, "", err
	}

	summary := "Executed Task-0008 backend validation and wrote an owned-lane implementation brief."
	return summary, stdoutPath, stderrPath, exitCode, workloadOutputPath, gitStatusAfter, nil
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
