package main

import (
	"fmt"
	"os"
)

func runEcho(arg string) {
	fmt.Println(arg)
}

func runType(arg string) {
	if arg == "echo" || arg == "pwd" || arg == "type" || arg == "exit" {
		fmt.Println(arg + " is a shell builtin")
	} else {
		path, found := findInPath(arg)
		if found {
			fmt.Println(arg + " is " + path)
		} else {
			fmt.Println(arg + ": not found")
		}
	}
}

func runExit() {
	os.Exit(0)
}
