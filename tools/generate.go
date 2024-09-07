package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	protoPath := os.Getenv("PROTOS_PATH")
	generatedPath := os.Getenv("GENERATED_PATH")

	// Validate environment variables
	if protoPath == "" || generatedPath == "" {
		fmt.Println("PROTOS_PATH or GENERATED_PATH is not set")
		os.Exit(1)
	}

	errorCreatingFolder := os.MkdirAll(generatedPath, 0755)
	if errorCreatingFolder != nil {
		fmt.Printf("Error creating directory: %v\n", errorCreatingFolder)
		return
	}

	// Run the protoc command
	cmd := exec.Command("protoc", "--go_out="+generatedPath, "--go-grpc_out="+generatedPath, "--proto_path="+protoPath, "password_policy.proto")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running protoc:", err)
		os.Exit(1)
	}
}
