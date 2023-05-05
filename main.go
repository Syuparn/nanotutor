package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

var (
	ja = flag.Bool("ja", false, "use Japanese Text")
)

func main() {
	flag.Parse()
	createTutorFile := createENTutorFile
	if *ja {
		createTutorFile = createJATutorFile
	}

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
