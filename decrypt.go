package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
)

type DecryptRequest struct {
	Key struct {
		Cipher string `json:"Cipher"`
	} `json:"Key"`
}

type DecryptResponse struct {
	Key struct {
		Plain string `json:"Plain"`
	} `json:"Key"`
}

func Decrypt(cfg *Config, resourceID string, ciphertext []byte, outputPath string) error {
	client := NewClient(cfg)

	req := DecryptRequest{}
	req.Key.Cipher = string(ciphertext)

	path := fmt.Sprintf("/kms/keys/%s/decrypt", resourceID)
	respData, err := client.doRequest("POST", path, req)
	if err != nil {
		return err
	}

	var resp DecryptResponse
	if err := json.Unmarshal(respData, &resp); err != nil {
		return fmt.Errorf("failed to parse response: %w", err)
	}

	plaintext, err := base64.StdEncoding.DecodeString(resp.Key.Plain)
	if err != nil {
		return fmt.Errorf("failed to decode base64: %w", err)
	}

	if err := os.WriteFile(outputPath, plaintext, 0600); err != nil {
		return fmt.Errorf("failed to write output: %w", err)
	}

	fmt.Fprintf(os.Stderr, "Successfully decrypted to %s\n", outputPath)
	return nil
}
