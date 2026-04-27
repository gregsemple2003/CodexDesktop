package taskrepo

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type TestingT interface {
	Helper()
	Fatalf(format string, args ...any)
	TempDir() string
}

type TaskFixture struct {
	TaskMD        string
	TaskState     string
	PlanMD        string
	HandoffMD     string
	ConstraintsMD string
	ExtraFiles    map[string]string
}

type Repo struct {
	Root          string
	TrackingRoot  string
	InitialCommit string
}

func CanonicalTasks() map[string]TaskFixture {
	return map[string]TaskFixture{
		"Task-0008": Task0008BackendRuntime(),
		"Task-0009": Task0009TasksSurface(),
	}
}

func Task0008BackendRuntime() TaskFixture {
	return TaskFixture{
		TaskMD: `# Task 0008

## Title

Build the backend task dispatch layer.

## Summary

Create the durable backend task-run contract so later clients do not guess state.
`,
		TaskState: `{
  "task_id": "Task-0008",
  "status": "in_progress",
  "phase": "implementation",
  "plan_approved": true,
  "current_pass": "PASS-0002",
  "current_gate": "implementation",
  "blockers": [],
  "updated_at": "2026-04-24T17:10:00-04:00"
}`,
		PlanMD: `# Task-0008 Plan

Build backend readback, dispatch, pause, owned-lane cleanup, and durable run state.
`,
		HandoffMD: `# Task-0008 Handoff

## Current Status

Backend runtime semantics are available for task dispatch and supervision tests.
`,
		ConstraintsMD: `# Task-0008 Constraints

- Preserve rollback-safe runtime truth.
- Do not use the human dashboard lane for disposable proof.
`,
	}
}

func Task0009TasksSurface() TaskFixture {
	return TaskFixture{
		TaskMD: `# Task 0009

## Title

Build the committed-work Tasks tab.

## Summary

Show authored and promoted committed tasks, backend run state, Pause controls, and launch targets without surfacing unpromoted candidates.
`,
		TaskState: `{
  "task_id": "Task-0009",
  "status": "completed",
  "phase": "closure",
  "plan_approved": true,
  "current_pass": "PASS-0004",
  "current_gate": "complete",
  "blockers": [],
  "updated_at": "2026-04-26T21:35:00-04:00"
}`,
		PlanMD: `# Task-0009 Plan

Keep the Tasks tab true to committed work: no unpromoted candidates, no Candidate provenance labels, visible Pause copy, and no false AI-run progress bar.
`,
		HandoffMD: `# Task-0009 Handoff

## Current Status

The Tasks tab consumes backend-shaped task readback and renders the committed-work surface.
`,
	}
}

func WriteCanonicalGitRepo(t TestingT) Repo {
	t.Helper()
	root := WriteGitTaskTrackingRoot(t, CanonicalTasks())
	return Repo{
		Root:          root,
		TrackingRoot:  filepath.Join(root, "Tracking"),
		InitialCommit: GitOutput(t, root, "rev-parse", "HEAD"),
	}
}

func WriteTaskTrackingRoot(t TestingT, tasks map[string]TaskFixture) string {
	t.Helper()
	worktreeRoot := t.TempDir()
	writeTasks(t, worktreeRoot, tasks)
	return worktreeRoot
}

func WriteGitTaskTrackingRoot(t TestingT, tasks map[string]TaskFixture) string {
	t.Helper()
	worktreeRoot := WriteTaskTrackingRoot(t, tasks)
	RunCommand(t, worktreeRoot, "git", "init")
	RunCommand(t, worktreeRoot, "git", "config", "user.email", "taskrepo-fixture@example.com")
	RunCommand(t, worktreeRoot, "git", "config", "user.name", "Task Repo Fixture")
	CommitAll(t, worktreeRoot, "initial task fixtures")
	return worktreeRoot
}

func CommitAll(t TestingT, worktreeRoot string, message string) string {
	t.Helper()
	RunCommand(t, worktreeRoot, "git", "add", ".")
	RunCommand(t, worktreeRoot, "git", "commit", "-m", message)
	return GitOutput(t, worktreeRoot, "rev-parse", "HEAD")
}

func WriteFile(t TestingT, path string, contents string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("mkdir %s: %v", filepath.Dir(path), err)
	}
	if err := os.WriteFile(path, []byte(contents), 0o644); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
}

func RunCommand(t TestingT, dir string, exe string, args ...string) {
	t.Helper()
	cmd := exec.Command(exe, args...)
	cmd.Dir = dir
	if output, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("%s %v failed: %v\n%s", exe, args, err, string(output))
	}
}

func GitOutput(t TestingT, worktreeRoot string, args ...string) string {
	t.Helper()
	cmdArgs := append([]string{"-C", worktreeRoot}, args...)
	cmd := exec.Command("git", cmdArgs...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("git %v failed: %v\n%s", args, err, string(output))
	}
	return strings.TrimSpace(string(output))
}

func writeTasks(t TestingT, worktreeRoot string, tasks map[string]TaskFixture) {
	t.Helper()
	trackingRoot := filepath.Join(worktreeRoot, "Tracking")
	if err := os.MkdirAll(trackingRoot, 0o755); err != nil {
		t.Fatalf("mkdir tracking root: %v", err)
	}

	for taskID, fixture := range tasks {
		taskRoot := filepath.Join(trackingRoot, taskID)
		WriteFile(t, filepath.Join(taskRoot, "TASK.md"), fixture.TaskMD)
		WriteFile(t, filepath.Join(taskRoot, "TASK-STATE.json"), fixture.TaskState)
		if fixture.PlanMD != "" {
			WriteFile(t, filepath.Join(taskRoot, "PLAN.md"), fixture.PlanMD)
		}
		if fixture.HandoffMD != "" {
			WriteFile(t, filepath.Join(taskRoot, "HANDOFF.md"), fixture.HandoffMD)
		}
		if fixture.ConstraintsMD != "" {
			WriteFile(t, filepath.Join(taskRoot, "CONSTRAINTS.md"), fixture.ConstraintsMD)
		}
		for relativePath, contents := range fixture.ExtraFiles {
			WriteFile(t, safeTaskPath(t, taskRoot, relativePath), contents)
		}
	}
}

func safeTaskPath(t TestingT, taskRoot string, relativePath string) string {
	t.Helper()
	clean := filepath.Clean(relativePath)
	if clean == "." || filepath.IsAbs(clean) || clean == ".." || strings.HasPrefix(clean, ".."+string(os.PathSeparator)) {
		t.Fatalf("unsafe fixture relative path %q", relativePath)
	}
	return filepath.Join(taskRoot, clean)
}
