package jobs

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const (
	APIVersion                   = "codex.jobs/v1"
	DesiredStateEnabled          = "enabled"
	DesiredStateDisabled         = "disabled"
	ExecutorTypeCodexExec        = "codex_exec"
	ExecutorTypePowerShellScript = "powershell_script"
	TriggerTypeSchedule          = "schedule"
	TriggerTypeManual            = "manual"
	TriggerTypeWebhook           = "webhook"
)

func LoadSpecs(jobsRoot string) ([]Spec, error) {
	specsDir := filepath.Join(jobsRoot, "specs")
	entries := []string{}
	err := filepath.WalkDir(specsDir, func(path string, d os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if d.IsDir() {
			return nil
		}
		if !strings.HasSuffix(d.Name(), ".json") || strings.HasSuffix(d.Name(), ".schema.json") {
			return nil
		}
		entries = append(entries, path)
		return nil
	})
	if err != nil {
		return nil, err
	}

	sort.Strings(entries)
	seen := map[string]struct{}{}
	specs := make([]Spec, 0, len(entries))
	for _, path := range entries {
		spec, err := loadSpecFile(path)
		if err != nil {
			return nil, err
		}
		if err := ValidateSpec(spec); err != nil {
			return nil, fmt.Errorf("%s: %w", path, err)
		}
		if _, exists := seen[spec.JobID]; exists {
			return nil, fmt.Errorf("duplicate job_id %q", spec.JobID)
		}
		seen[spec.JobID] = struct{}{}
		specs = append(specs, spec)
	}
	return specs, nil
}

func loadSpecFile(path string) (Spec, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return Spec{}, err
	}
	var spec Spec
	if err := json.Unmarshal(raw, &spec); err != nil {
		return Spec{}, err
	}
	return spec, nil
}

func ValidateSpec(spec Spec) error {
	if spec.APIVersion != APIVersion {
		return fmt.Errorf("unsupported api_version %q", spec.APIVersion)
	}
	if spec.JobID == "" || spec.Label == "" || spec.Description == "" {
		return fmt.Errorf("job metadata is incomplete")
	}
	if spec.DesiredState != DesiredStateEnabled && spec.DesiredState != DesiredStateDisabled {
		return fmt.Errorf("unsupported desired_state %q", spec.DesiredState)
	}
	if len(spec.Triggers) == 0 {
		return fmt.Errorf("at least one trigger is required")
	}
	for _, trigger := range spec.Triggers {
		switch trigger.Type {
		case TriggerTypeSchedule:
			if trigger.Cron == "" || trigger.Timezone == "" {
				return fmt.Errorf("schedule trigger requires cron and timezone")
			}
		case TriggerTypeManual:
		case TriggerTypeWebhook:
			if trigger.Path == "" {
				return fmt.Errorf("webhook trigger requires path")
			}
		default:
			return fmt.Errorf("unsupported trigger type %q", trigger.Type)
		}
	}
	if spec.Executor.Cwd == "" {
		return fmt.Errorf("executor requires cwd")
	}
	if err := validateStringSlice(spec.Executor.Args, "executor args"); err != nil {
		return err
	}
	if err := validateStringSlice(spec.Executor.ManualArgs, "executor manual_args"); err != nil {
		return err
	}
	if err := validateStringSlice(spec.Executor.WebhookArgs, "executor webhook_args"); err != nil {
		return err
	}
	switch spec.Executor.Type {
	case ExecutorTypeCodexExec:
		if spec.Executor.Entrypoint == "" {
			return fmt.Errorf("codex_exec executor requires entrypoint")
		}
	case ExecutorTypePowerShellScript:
		if spec.Executor.ScriptPath == "" {
			return fmt.Errorf("powershell_script executor requires script_path")
		}
	default:
		return fmt.Errorf("unsupported executor type %q", spec.Executor.Type)
	}
	if spec.Runtime.WorkflowType == "" || spec.Runtime.TaskQueue == "" {
		return fmt.Errorf("runtime requires workflow_type and task_queue")
	}
	return nil
}

func validateStringSlice(values []string, label string) error {
	for _, value := range values {
		if value == "" {
			return fmt.Errorf("%s must not contain empty strings", label)
		}
	}
	return nil
}
