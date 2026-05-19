package main

import (
	"fmt"
	"os"
)

func pwd() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(dir)

}
