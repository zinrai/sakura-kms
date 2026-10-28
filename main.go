package main

import (
	"fmt"
	"io"
	"os"

	v1 "github.com/sacloud/sacloud-sdk-go/api/kms/apis/v1"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "encrypt":
		if err := runEncrypt(os.Args[2:]); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	case "decrypt":
		if err := runDecrypt(os.Args[2:]); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	case "version":
		runVersion()
	default:
		printUsage()
		os.Exit(1)
	}
}

func runEncrypt(args []string) error {
	in, err := prepare(args)
	if err != nil {
		return err
	}

	ciphertext, err := encrypt(in.client, in.keyID, in.data)
	if err != nil {
		return err
	}

	if _, err := os.Stdout.WriteString(ciphertext); err != nil {
		return fmt.Errorf("failed to write output: %w", err)
	}

	fmt.Fprintf(os.Stderr, "encrypted %d bytes with key %s\n", len(in.data), in.keyID)
	return nil
}

func runDecrypt(args []string) error {
	in, err := prepare(args)
	if err != nil {
		return err
	}

	plaintext, err := decrypt(in.client, in.keyID, in.data)
	if err != nil {
		return err
	}

	if _, err := os.Stdout.Write(plaintext); err != nil {
		return fmt.Errorf("failed to write output: %w", err)
	}

	fmt.Fprintf(os.Stderr, "decrypted %d bytes with key %s\n", len(plaintext), in.keyID)
	return nil
}

type input struct {
	client *v1.Client
	keyID  string
	data   []byte
}

func prepare(args []string) (*input, error) {
	if err := rejectArgs(args); err != nil {
		return nil, err
	}

	keyID, err := loadKeyID()
	if err != nil {
		return nil, err
	}

	client, err := newKMSClient()
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return nil, fmt.Errorf("failed to read stdin: %w", err)
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("no data provided via stdin")
	}

	return &input{client: client, keyID: keyID, data: data}, nil
}

func rejectArgs(args []string) error {
	// Not ignored: stdin would be read instead, waiting on the terminal.
	if len(args) > 0 {
		return fmt.Errorf("unexpected argument: %s", args[0])
	}
	return nil
}

func printUsage() {
	fmt.Fprintf(os.Stderr, `Usage: sakura-kms <command>

Commands:
  encrypt    Encrypt data from stdin, write ciphertext to stdout
  decrypt    Decrypt data from stdin, write plaintext to stdout
  version    Print version

Environment variables:
  SAKURA_KMS_KEY_ID     KMS key resource ID (required)

  API credentials and the API endpoint are resolved by saclient-go.
`)
}
