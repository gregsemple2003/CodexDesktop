package httpapi

import (
	"encoding/json"
	"net/http"

	"github.com/gregsemple2003/CodexDesktop/backend/orchestration/internal/config"
)

func NewMux(cfg config.Config) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		_ = json.NewEncoder(w).Encode(map[string]string{
			"status":    "ok",
			"jobs_root": cfg.JobsRoot,
			"namespace": cfg.Namespace,
			"task_queue": cfg.TaskQueue,
		})
	})
	return mux
}
