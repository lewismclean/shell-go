package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")

		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
		command = strings.TrimSpace(command)
		if command == "exit" {
			runExit()
		} else if strings.HasPrefix(command, "echo ") {
			runEcho(command[5:])
		} else if strings.HasPrefix(command, "type ") {
			runType(command[5:])
		} else {
			parts := strings.SplitN(command, " ", 0)
			path, found := findInPath(parts[0])
			if found {
				runExecutable(path, parts[1:])
			} else {
				fmt.Println(command + ": command not found")
			}
		}
	}
}
