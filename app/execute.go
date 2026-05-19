package main

import (
	"os"
	"os/exec"
)

func runExecutable(path string, args []string) {
	cmd := exec.Command(path, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
