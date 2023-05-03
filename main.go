package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	tempFilePath, err := createTutorFile()
	if err != nil {
		fmt.Fprintf(os.Stderr, "FATAL: %v\n", err)
		os.Exit(1)
	}
	defer os.Remove(tempFilePath)

	// run nano
	cmd := exec.Command("nano", tempFilePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "WARNING: nano closed unexpectedly: %v\n", err)
		os.Exit(1)
	}
}
