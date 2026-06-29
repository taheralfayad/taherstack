package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	homeDir, err := os.UserHomeDir()
	scanner := bufio.NewScanner(os.Stdin)
	if err != nil {
		fmt.Printf("Failed to get home directory: %v\n", err)
		return
	}

	fmt.Print("Enter the name of your project: ")
	var directoryName string

	if scanner.Scan() {
		directoryName = scanner.Text()
	}

	targetPath := filepath.Join(homeDir, directoryName)

	err = os.MkdirAll(targetPath, 0755)

	if err != nil {
		fmt.Printf("Failed to create directory: %v\n", err)
		return
	}

	cmd := exec.Command("git", "init", targetPath)

	_, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("git init failed: Output: %s\n", err)
	}

	var subdirs = [2]string{"backend", "frontend"}

	for _, dir := range subdirs {
		subPath := filepath.Join(targetPath, dir)
		err = os.MkdirAll(subPath, 0755)

		if err != nil {
			fmt.Printf("Failed to create subdir %v\n", err)
		}
	}

	fmt.Printf("Success! Project initialized at: %s\n", targetPath)
}
