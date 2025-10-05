package main

import (
	"fmt"
	"os"
)

type Config struct {
	Token  string
	Secret string
	Zone   string
}

func LoadConfig(zone string) (*Config, error) {
	cfg := &Config{
		Token:  os.Getenv("SAKURACLOUD_ACCESS_TOKEN"),
		Secret: os.Getenv("SAKURACLOUD_ACCESS_TOKEN_SECRET"),
		Zone:   zone,
	}

	var missing []string
	if cfg.Token == "" {
		missing = append(missing, "SAKURACLOUD_ACCESS_TOKEN")
	}
	if cfg.Secret == "" {
		missing = append(missing, "SAKURACLOUD_ACCESS_TOKEN_SECRET")
	}

	if len(missing) > 0 {
		return nil, fmt.Errorf("required environment variables not set: %v", missing)
	}

	return cfg, nil
}
