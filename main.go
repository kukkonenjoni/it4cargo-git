package main

import (
	"fmt"

	changelog "github.com/kukkonenjoni/it4cargo-git/commands"
	"github.com/kukkonenjoni/it4cargo-git/util"
)

func main() {
	// Example: git log
	output, err := util.RunGitCommand("log", "--pretty=format:'%s'")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Output:\n", output)

	output, err = util.RunGitCommand("pull")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Output: ", output)

	changelog.CreateChangeLog()
}
