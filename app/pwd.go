package main

import (
	"os"
)

func pwd() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	} else {
		return dir, nil
	}

}
