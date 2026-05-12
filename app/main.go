package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print
var pathDirs = []string(strings.Split(os.Getenv("PATH"), string(os.PathListSeparator)))

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
				cmdName := command[5:]
				found := false
				for _, dir := range pathDirs {
					files, err := os.ReadDir(dir)
					if err != nil {
						continue
					}
					for _, file := range files {
						if file.Name() == cmdName {
							info, err := file.Info()
							if err != nil {
								continue
							}
							if info.Mode().Perm()&0111 != 0 {
								fullPath := filepath.Join(dir, cmdName)
								fmt.Println(cmdName + " is " + fullPath)
								found = true
								break
							}
						}
					}
					if found {
						break
					}
				}
				if !found {
					fmt.Println(command[5:] + ": not found")
				}
			}
		} else {
			fmt.Println(command + ": command not found")
		}
	}
}
