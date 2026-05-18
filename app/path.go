package main

import (
	"os"
	"path/filepath"
	"strings"
)

func findInPath(name string) (string, bool) {
	pathDirs := strings.Split(os.Getenv("PATH"), string(os.PathListSeparator))
	for _, dir := range pathDirs {
		files, err := os.ReadDir(dir)
		if err != nil {
			continue
		}
		for _, file := range files {
			if file.Name() == name {
				info, err := file.Info()
				if err != nil {
					continue
				}
				if info.Mode().Perm()&0111 != 0 {
					return filepath.Join(dir, name), true
				}
			}
		}
	}
	return "", false
}
