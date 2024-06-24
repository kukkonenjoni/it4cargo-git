package main

import (
	"fmt"
	"os/exec"
)

func runGitCommand(args ...string) (string, error) {
	// Create the command with "git" and the passed arguments
	cmd := exec.Command("git", args...)

	// Get the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func main() {
	// Example: git status
	output, err := runGitCommand("status")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Output:\n", output)

	// Example: git log
	output, err = runGitCommand("log")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Output:\n", output)
}
