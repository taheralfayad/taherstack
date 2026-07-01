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
		workDir, err := os.Getwd()

		if err != nil {
			fmt.Printf("failed to get workdir %v\n", err)
			continue
		}

		currPath := filepath.Join(workDir, dir)
		targetPath := filepath.Join(targetPath, dir)

		cmd := exec.Command("cp", "-r", currPath, targetPath)

		_, err = cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Failed to create subdir %v\n", err)
			continue
		}
	}

	var files = [5]string{".gitignore", "Caddyfile", ".env.template", "SAMPLE_README.md", "docker-compose.yml"}

	for _, file := range files {
		var targetFile string
		workDir, err := os.Getwd()

		if err != nil {
			fmt.Printf("failed to get workdir %v\n", err)
			continue
		}

		currPath := filepath.Join(workDir, file)

		if file == "SAMPLE_README.md" {
			targetFile = "README.md"
		} else {
			targetFile = file
		}

		targetPath := filepath.Join(targetPath, targetFile)

		cmd := exec.Command("cp", currPath, targetPath)

		_, err = cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Failed to create subdir %v\n", err)
			continue
		}
	}
	fmt.Printf("Success! Project initialized at: %s\n", targetPath)
}
