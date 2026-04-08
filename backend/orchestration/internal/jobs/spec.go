package jobs

type Spec struct {
	APIVersion   string        `json:"api_version"`
	JobID        string        `json:"job_id"`
	Label        string        `json:"label"`
	Description  string        `json:"description"`
	DesiredState string        `json:"desired_state"`
	Triggers     []Trigger     `json:"triggers"`
	Executor     Executor      `json:"executor"`
	Runtime      RuntimeConfig `json:"runtime"`
	Source       *SourceRef    `json:"source,omitempty"`
}

type Trigger struct {
	Type     string `json:"type"`
	Cron     string `json:"cron,omitempty"`
	Timezone string `json:"timezone,omitempty"`
	Path     string `json:"path,omitempty"`
}

type Executor struct {
	Type        string   `json:"type"`
	Cwd         string   `json:"cwd"`
	Entrypoint  string   `json:"entrypoint,omitempty"`
	ScriptPath  string   `json:"script_path,omitempty"`
	Args        []string `json:"args,omitempty"`
	ManualArgs  []string `json:"manual_args,omitempty"`
	WebhookArgs []string `json:"webhook_args,omitempty"`
}

type RuntimeConfig struct {
	WorkflowType string `json:"workflow_type"`
	TaskQueue    string `json:"task_queue"`
}

type SourceRef struct {
	Kind         string `json:"kind"`
	LegacyJobID  string `json:"legacy_job_id,omitempty"`
	LegacySource string `json:"legacy_source,omitempty"`
}
