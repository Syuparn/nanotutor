package main

import (
	_ "embed"

	"fmt"
	"os"
)

//go:embed tutor
var tutorTextBytes []byte

func createTutorFile() (string, error) {
	dir := os.TempDir()
	f, err := os.CreateTemp(dir, "nanotutor_*")
	if err != nil {
		return "", fmt.Errorf("failed to create tutor temp file: %w", err)
	}

	_, err = f.Write(tutorTextBytes)
	if err != nil {
		return "", fmt.Errorf("failed to initialize tutor temp file: %w", err)
	}

	// return absolute path
	return f.Name(), nil
}
