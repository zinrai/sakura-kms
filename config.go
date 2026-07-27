package main

import (
	"fmt"
	"os"
)

// LoadKeyID resolves the KMS key resource ID from the environment.
func LoadKeyID() (string, error) {
	if v := os.Getenv("SAKURA_KMS_KEY_ID"); v != "" {
		return v, nil
	}
	return "", fmt.Errorf("required environment variable not set: SAKURA_KMS_KEY_ID")
}
