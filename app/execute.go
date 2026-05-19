package main

import (
	"os"
	"os/exec"
)

func runExecutable(name string, path string, args []string) {
	cmd := exec.Command(path, args...)
	cmd.Args := append([]string{name}, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
