package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
)

type EncryptRequest struct {
	Key struct {
		Plain string `json:"Plain"`
		Algo  string `json:"Algo"`
	} `json:"Key"`
}

type EncryptResponse struct {
	Key struct {
		Cipher string `json:"Cipher"`
	} `json:"Key"`
}

func Encrypt(cfg *Config, resourceID string, plaintext []byte, outputPath string) error {
	client := NewClient(cfg)

	req := EncryptRequest{}
	req.Key.Plain = base64.StdEncoding.EncodeToString(plaintext)
	req.Key.Algo = "aes-256-gcm"

	path := fmt.Sprintf("/kms/keys/%s/encrypt", resourceID)
	respData, err := client.doRequest("POST", path, req)
	if err != nil {
		return err
	}

	var resp EncryptResponse
	if err := json.Unmarshal(respData, &resp); err != nil {
		return fmt.Errorf("failed to parse response: %w", err)
	}

	if err := os.WriteFile(outputPath, []byte(resp.Key.Cipher), 0600); err != nil {
		return fmt.Errorf("failed to write output: %w", err)
	}

	fmt.Fprintf(os.Stderr, "Successfully encrypted to %s\n", outputPath)
	return nil
}
