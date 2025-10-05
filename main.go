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
	default:
		printUsage()
		os.Exit(1)
	}
}

func runEncrypt(args []string) error {
	fs := flag.NewFlagSet("encrypt", flag.ExitOnError)
	output := fs.String("output", "", "Output file path (required)")
	resourceID := fs.String("resource-id", "", "KMS key resource ID (required)")
	zone := fs.String("zone", "is1a", "SAKURA Cloud zone")
	fs.Parse(args)

	if *output == "" || *resourceID == "" {
		fs.Usage()
		return fmt.Errorf("both -output and -resource-id are required")
	}

	cfg, err := LoadConfig(*zone)
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

	return Encrypt(cfg, *resourceID, data, *output)
}

func runDecrypt(args []string) error {
	fs := flag.NewFlagSet("decrypt", flag.ExitOnError)
	output := fs.String("output", "", "Output file path (required)")
	resourceID := fs.String("resource-id", "", "KMS key resource ID (required)")
	zone := fs.String("zone", "is1a", "SAKURA Cloud zone")
	fs.Parse(args)

	if *output == "" || *resourceID == "" {
		fs.Usage()
		return fmt.Errorf("both -output and -resource-id are required")
	}

	cfg, err := LoadConfig(*zone)
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

	return Decrypt(cfg, *resourceID, data, *output)
}

func printUsage() {
	fmt.Fprintf(os.Stderr, `Usage: sakura-kms <command> [options]

Commands:
  encrypt    Encrypt data from stdin
  decrypt    Decrypt data from stdin

Encrypt options:
  -output string        Output file path (required)
  -resource-id string   KMS key resource ID (required)
  -zone string          SAKURA Cloud zone (default "is1a")

Decrypt options:
  -output string        Output file path (required)
  -resource-id string   KMS key resource ID (required)
  -zone string          SAKURA Cloud zone (default "is1a")

Environment variables (required):
  SAKURACLOUD_ACCESS_TOKEN          API token
  SAKURACLOUD_ACCESS_TOKEN_SECRET   API secret

Example:
  cat secret.txt | sakura-kms encrypt -output secret.enc -resource-id 110000000000
  cat secret.txt | sakura-kms encrypt -output secret.enc -resource-id 110000000000 -zone tk1a
  cat secret.enc | sakura-kms decrypt -output secret.txt -resource-id 110000000000
`)
}
