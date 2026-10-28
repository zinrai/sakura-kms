package main

import (
	"context"
	"fmt"

	kms "github.com/sacloud/sacloud-sdk-go/api/kms"
	v1 "github.com/sacloud/sacloud-sdk-go/api/kms/apis/v1"
)

func decrypt(client *v1.Client, keyID string, ciphertext []byte) ([]byte, error) {
	keyOp := kms.NewKeyOp(client)

	plaintext, err := keyOp.Decrypt(context.Background(), keyID, string(ciphertext))
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt: %w", err)
	}
	return plaintext, nil
}
