package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print
pathFiles := []string(os.ReadDir(PATH))

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
			break
		} else if strings.HasPrefix(command, "echo ") {
			fmt.Println(command[5:])
		} else if strings.HasPrefix(command, "type ") {
			if command[5:] == "echo" || command[5:] == "type" || command[5:] == "exit" {
				fmt.Println(command[5:] + " is a shell builtin")
			} else {
				for _, file := range pathFiles {
					if file.Name() == command[5:] && file.type().IsExec() {
						println(command[5:] + " is" + file.Path())
						break
					}
				}
						
				fmt.Println(command[5:] + ": not found")
			}
		} else {
			fmt.Println(command + ": command not found")
		}
	}
}
