package main

import (
	"fmt"
	"os"
)

// LoadKeyID resolves the KMS key resource ID from environment variables.
// SAKURA_KMS_KEY_ID takes precedence over the legacy SAKURACLOUD_KMS_KEY_ID.
func LoadKeyID() (string, error) {
	for _, name := range []string{"SAKURA_KMS_KEY_ID", "SAKURACLOUD_KMS_KEY_ID"} {
		if v := os.Getenv(name); v != "" {
			return v, nil
		}
	}
	return "", fmt.Errorf("required environment variable not set: SAKURA_KMS_KEY_ID or SAKURACLOUD_KMS_KEY_ID")
}
