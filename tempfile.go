package main

import (
	_ "embed"

	"fmt"
	"os"
)

//go:embed tutor
var tutorTextBytes []byte

//go:embed tutor-ja
var tutorJATextBytes []byte

func createENTutorFile() (string, error) {
	return createTutorFile(tutorTextBytes)
}

func createJATutorFile() (string, error) {
	return createTutorFile(tutorJATextBytes)
}

func createTutorFile(text []byte) (string, error) {
	dir := os.TempDir()
	f, err := os.CreateTemp(dir, "nanotutor_*")
	if err != nil {
		return "", fmt.Errorf("failed to create tutor temp file: %w", err)
	}

	_, err = f.Write(text)
	if err != nil {
		return "", fmt.Errorf("failed to initialize tutor temp file: %w", err)
	}

	// return absolute path
	return f.Name(), nil
}
