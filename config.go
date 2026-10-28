package main

import (
	"fmt"
	"os"
)

func loadKeyID() (string, error) {
	if v := os.Getenv("SAKURA_KMS_KEY_ID"); v != "" {
		return v, nil
	}
	return "", fmt.Errorf("required environment variable not set: SAKURA_KMS_KEY_ID")
}
