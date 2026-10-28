package main

import (
	"context"
	"fmt"

	kms "github.com/sacloud/sacloud-sdk-go/api/kms"
	v1 "github.com/sacloud/sacloud-sdk-go/api/kms/apis/v1"
)

func encrypt(client *v1.Client, keyID string, plaintext []byte) (string, error) {
	keyOp := kms.NewKeyOp(client)

	ciphertext, err := keyOp.Encrypt(context.Background(), keyID, plaintext, v1.EncryptionRequestAlgoAes256Gcm)
	if err != nil {
		return "", fmt.Errorf("failed to encrypt: %w", err)
	}
	return ciphertext, nil
}
