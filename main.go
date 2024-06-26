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
	// Example: git log
	output, err := runGitCommand("log", "--pretty=format:'%s'")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Output:\n", output)

	output, err = runGitCommand("pull")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Output: ", output)
}
