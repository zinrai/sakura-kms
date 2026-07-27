package main

import (
	"flag"
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
	fs := flag.NewFlagSet("encrypt", flag.ExitOnError)
	zone := fs.String("zone", "is1a", "SAKURA Cloud zone")
	fs.Parse(args)

	keyID, err := LoadKeyID()
	if err != nil {
		return err
	}

	client, err := NewKMSClient(*zone)
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
	fs := flag.NewFlagSet("decrypt", flag.ExitOnError)
	zone := fs.String("zone", "is1a", "SAKURA Cloud zone")
	fs.Parse(args)

	keyID, err := LoadKeyID()
	if err != nil {
		return err
	}

	client, err := NewKMSClient(*zone)
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

func printUsage() {
	fmt.Fprintf(os.Stderr, `Usage: sakura-kms <command> [options]

Commands:
  encrypt    Encrypt data from stdin, write ciphertext to stdout
  decrypt    Decrypt data from stdin, write plaintext to stdout
  version    Print version

Encrypt options:
  -zone string          SAKURA Cloud zone (default "is1a")

Decrypt options:
  -zone string          SAKURA Cloud zone (default "is1a")

Environment variables:
  SAKURA_KMS_KEY_ID                 KMS key resource ID (required)

  API credentials are resolved by saclient-go. Set either static API keys:
    SAKURA_ACCESS_TOKEN, SAKURA_ACCESS_TOKEN_SECRET
  or service principal credentials:
    SAKURA_SERVICE_PRINCIPAL_ID, SAKURA_SERVICE_PRINCIPAL_KEY_ID,
    SAKURA_PRIVATE_KEY_PATH (or SAKURA_PRIVATE_KEY)

Example:
  sakura-kms encrypt < secret.txt > secret.enc
  sakura-kms encrypt -zone tk1a < secret.txt > secret.enc
  sakura-kms decrypt < secret.enc > secret.txt
  sakura-kms decrypt < secret.enc | my-application --password-stdin
`)
}
