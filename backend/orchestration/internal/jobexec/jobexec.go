package jobexec

import (
	"context"
	"errors"
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

	"github.com/gregsemple2003/CodexDesktop/backend/orchestration/internal/config"
	"github.com/gregsemple2003/CodexDesktop/backend/orchestration/internal/controlplane"
)

const RunCodexExecActivityName = "codex.exec.activity"

type CommandPlan struct {
	Executable       string
	Args             []string
	RunRoot          string
	EventLogPath     string
	FinalMessagePath string
	StderrPath       string
}

type codexExecActivity struct {
	cfg config.Config
}

func Register(w worker.Worker, cfg config.Config) {
	activityImpl := &codexExecActivity{cfg: cfg}
	w.RegisterWorkflowWithOptions(CodexExecWorkflow, workflow.RegisterOptions{Name: "codex.exec.job"})
	w.RegisterActivityWithOptions(activityImpl.Execute, activity.RegisterOptions{Name: RunCodexExecActivityName})
}

func CodexExecWorkflow(ctx workflow.Context, request controlplane.JobRunRequest) (controlplane.JobRunResult, error) {
	info := workflow.GetInfo(ctx)
	request.WorkflowID = info.WorkflowExecution.ID
	request.RunID = info.WorkflowExecution.RunID
	if request.RequestedAt.IsZero() {
		request.RequestedAt = workflow.Now(ctx).UTC()
	}

	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 6 * time.Hour,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts: 1,
		},
	})

	var result controlplane.JobRunResult
	if err := workflow.ExecuteActivity(ctx, RunCodexExecActivityName, request).Get(ctx, &result); err != nil {
		return result, err
	}
	return result, nil
}

func BuildCommandPlan(cfg config.Config, request controlplane.JobRunRequest) (CommandPlan, error) {
	if request.Spec.Executor.Cwd == "" {
		return CommandPlan{}, fmt.Errorf("job %s executor cwd is empty", request.JobID)
	}

	runRoot := filepath.Join(cfg.RunsRoot, sanitizePathSegment(request.JobID), sanitizePathSegment(request.WorkflowID))
	eventLogPath := filepath.Join(runRoot, "events.jsonl")
	finalMessagePath := filepath.Join(runRoot, "final-message.txt")
	stderrPath := filepath.Join(runRoot, "stderr.txt")

	args := []string{
		"exec",
		"--full-auto",
		"--skip-git-repo-check",
		"--json",
		"-C", request.Spec.Executor.Cwd,
		"-o", finalMessagePath,
		buildPrompt(request),
	}

	return CommandPlan{
		Executable:       cfg.CodexExecutable,
		Args:             args,
		RunRoot:          runRoot,
		EventLogPath:     eventLogPath,
		FinalMessagePath: finalMessagePath,
		StderrPath:       stderrPath,
	}, nil
}

func (a *codexExecActivity) Execute(ctx context.Context, request controlplane.JobRunRequest) (controlplane.JobRunResult, error) {
	plan, err := BuildCommandPlan(a.cfg, request)
	if err != nil {
		return controlplane.JobRunResult{}, err
	}

	if err := os.MkdirAll(plan.RunRoot, 0o755); err != nil {
		return controlplane.JobRunResult{}, fmt.Errorf("create run root: %w", err)
	}

	eventLogFile, err := os.Create(plan.EventLogPath)
	if err != nil {
		return controlplane.JobRunResult{}, fmt.Errorf("create event log: %w", err)
	}
	defer eventLogFile.Close()

	stderrFile, err := os.Create(plan.StderrPath)
	if err != nil {
		return controlplane.JobRunResult{}, fmt.Errorf("create stderr log: %w", err)
	}
	defer stderrFile.Close()

	cmd := exec.CommandContext(ctx, plan.Executable, plan.Args...)
	cmd.Dir = request.Spec.Executor.Cwd
	cmd.Stdout = eventLogFile
	cmd.Stderr = stderrFile

	startedAt := time.Now().UTC()
	runErr := cmd.Run()
	completedAt := time.Now().UTC()

	result := controlplane.JobRunResult{
		JobID:            request.JobID,
		TriggerType:      request.TriggerType,
		DesiredSpecHash:  request.DesiredSpecHash,
		RequestedAt:      request.RequestedAt,
		WorkflowID:       request.WorkflowID,
		RunID:            request.RunID,
		StartedAt:        startedAt,
		CompletedAt:      completedAt,
		EventLogPath:     plan.EventLogPath,
		FinalMessagePath: plan.FinalMessagePath,
		StderrPath:       plan.StderrPath,
		Command:          append([]string{plan.Executable}, plan.Args...),
	}

	if rawFinalMessage, err := os.ReadFile(plan.FinalMessagePath); err == nil {
		result.FinalMessage = strings.TrimSpace(string(rawFinalMessage))
	}

	if runErr == nil {
		return result, nil
	}

	result.ExitCode = exitCode(runErr)
	return result, fmt.Errorf("codex exec failed with exit code %d: %w", result.ExitCode, runErr)
}

func buildPrompt(request controlplane.JobRunRequest) string {
	argsText := "(none)"
	if len(request.Spec.Executor.Args) > 0 {
		argsText = strings.Join(request.Spec.Executor.Args, " ")
	}

	lines := []string{
		fmt.Sprintf("Use the `%s` skill or task entrypoint available in this Codex environment.", request.Spec.Executor.Entrypoint),
		fmt.Sprintf("Run it with arguments: %s.", argsText),
		fmt.Sprintf("Job id: %s.", request.JobID),
		fmt.Sprintf("Trigger type: %s.", request.TriggerType),
		fmt.Sprintf("Desired spec hash: %s.", request.DesiredSpecHash),
		fmt.Sprintf("Temporal workflow id: %s.", request.WorkflowID),
		fmt.Sprintf("Temporal run id: %s.", request.RunID),
	}
	if request.TriggerPath != "" {
		lines = append(lines, fmt.Sprintf("Webhook path: %s.", request.TriggerPath))
	}
	lines = append(lines, "Follow the entrypoint's normal behavior and finish with a concise summary of what happened and any operator follow-up required.")
	return strings.Join(lines, "\n")
}

func sanitizePathSegment(value string) string {
	replacer := strings.NewReplacer("\\", "_", "/", "_", ":", "_")
	return replacer.Replace(value)
}

func exitCode(runErr error) int {
	if runErr == nil {
		return 0
	}
	var exitError *exec.ExitError
	if errors.As(runErr, &exitError) {
		return exitError.ExitCode()
	}
	return 1
}
