package util

import (
	"os"
	"os/exec"
)

func RunGitCommand(args ...string) (string, error) {
	// Create the command with "git" and the passed arguments
	cmd := exec.Command("git", args...)

	// Get the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func ChangeDirectory(path string) error {
	err := os.Chdir(path)
	if err != nil {
		return err
	}
	return nil
}
