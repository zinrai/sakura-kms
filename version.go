package main

import "fmt"

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func runVersion() {
	fmt.Printf("sakura-kms version %s\n", version)
	fmt.Printf("commit: %s\n", commit)
	fmt.Printf("built: %s\n", date)
}
