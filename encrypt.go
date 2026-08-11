package main

import (
	"context"
	"fmt"
	"os"

	kms "github.com/sacloud/sacloud-sdk-go/api/kms"
	v1 "github.com/sacloud/sacloud-sdk-go/api/kms/apis/v1"
)

func Encrypt(client *v1.Client, keyID string, plaintext []byte) error {
	keyOp := kms.NewKeyOp(client)

	ciphertext, err := keyOp.Encrypt(context.Background(), keyID, plaintext, v1.KeyEncryptAlgoEnumAes256Gcm)
	if err != nil {
		return fmt.Errorf("failed to encrypt: %w", err)
	}

	if _, err := os.Stdout.WriteString(ciphertext); err != nil {
		return fmt.Errorf("failed to write output: %w", err)
	}

	fmt.Fprintf(os.Stderr, "encrypted %d bytes with key %s\n", len(plaintext), keyID)
	return nil
}
