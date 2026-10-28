package main

import (
	"fmt"
	"io"
	"os"
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
	if err := rejectArgs(args); err != nil {
		return err
	}

	keyID, err := LoadKeyID()
	if err != nil {
		return err
	}

	client, err := NewKMSClient()
	if err != nil {
		return err
	}

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf("failed to read stdin: %w", err)
	}

	if len(data) == 0 {
		return fmt.Errorf("no data provided via stdin")
	}

	return Encrypt(client, keyID, data)
}

func runDecrypt(args []string) error {
	if err := rejectArgs(args); err != nil {
		return err
	}

	keyID, err := LoadKeyID()
	if err != nil {
		return err
	}

	client, err := NewKMSClient()
	if err != nil {
		return err
	}

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf("failed to read stdin: %w", err)
	}

	if len(data) == 0 {
		return fmt.Errorf("no data provided via stdin")
	}

	return Decrypt(client, keyID, data)
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
