package main

import (
	"context"
	"fmt"
	"os"

	kms "github.com/sacloud/kms-api-go"
	v1 "github.com/sacloud/kms-api-go/apis/v1"
)

func Decrypt(client *v1.Client, keyID string, ciphertext []byte) error {
	keyOp := kms.NewKeyOp(client)

	plaintext, err := keyOp.Decrypt(context.Background(), keyID, string(ciphertext))
	if err != nil {
		return fmt.Errorf("failed to decrypt: %w", err)
	}

	if _, err := os.Stdout.Write(plaintext); err != nil {
		return fmt.Errorf("failed to write output: %w", err)
	}

	fmt.Fprintf(os.Stderr, "decrypted %d bytes with key %s\n", len(plaintext), keyID)
	return nil
}
