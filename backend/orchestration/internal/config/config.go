package config

import (
	"errors"
	"os"
	"path/filepath"
)

type Config struct {
	BindAddress string
	JobsRoot    string
	Namespace   string
	TaskQueue   string
}

func Load() (Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return Config{}, errors.New("resolve home directory")
	}

	cfg := Config{
		BindAddress: envOrDefault("CODEX_ORCHESTRATION_BIND_ADDRESS", "127.0.0.1:4318"),
		JobsRoot:    envOrDefault("CODEX_ORCHESTRATION_JOBS_ROOT", filepath.Join(home, ".codex", "Orchestration", "Jobs")),
		Namespace:   envOrDefault("CODEX_ORCHESTRATION_NAMESPACE", "default"),
		TaskQueue:   envOrDefault("CODEX_ORCHESTRATION_TASK_QUEUE", "codex-orchestration"),
	}

	if cfg.JobsRoot == "" {
		return Config{}, errors.New("jobs root must not be empty")
	}

	return cfg, nil
}

func envOrDefault(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
