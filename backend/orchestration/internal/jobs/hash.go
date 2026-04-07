package jobs

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

func HashSpec(spec Spec) (string, error) {
	payload, err := json.Marshal(spec)
	if err != nil {
		return "", fmt.Errorf("marshal spec %s: %w", spec.JobID, err)
	}

	digest := sha256.Sum256(payload)
	return hex.EncodeToString(digest[:]), nil
}
