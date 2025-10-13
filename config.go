package main

import (
	"fmt"
	"os"
)

type Config struct {
	Token  string
	Secret string
	Zone   string
	KeyID  string
}

func LoadConfig(zone string) (*Config, error) {
	cfg := &Config{
		Token:  os.Getenv("SAKURACLOUD_ACCESS_TOKEN"),
		Secret: os.Getenv("SAKURACLOUD_ACCESS_TOKEN_SECRET"),
		KeyID:  os.Getenv("SAKURACLOUD_KMS_KEY_ID"),
		Zone:   zone,
	}

	var missing []string
	if cfg.Token == "" {
		missing = append(missing, "SAKURACLOUD_ACCESS_TOKEN")
	}
	if cfg.Secret == "" {
		missing = append(missing, "SAKURACLOUD_ACCESS_TOKEN_SECRET")
	}
	if cfg.KeyID == "" {
		missing = append(missing, "SAKURACLOUD_KMS_KEY_ID")
	}

	if len(missing) > 0 {
		return nil, fmt.Errorf("required environment variables not set: %v", missing)
	}

	return cfg, nil
}
