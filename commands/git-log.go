package it4cargoGit

import (
	"fmt"

	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/kukkonenjoni/it4cargo-git/util"
)

func GetGitLogs(taggedHash object.Commit) (string, error) {
	output, err := util.RunGitCommand("log", "--pretty=format:'%s'", taggedHash.Hash.String()+"..")
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	return output, nil
}
