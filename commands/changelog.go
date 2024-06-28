package it4cargoGit

import (
	"fmt"
	"strings"

	"log"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func CreateChangeLog() (string, error) {
	lastTaggedCommit := getCommitMessages()
	fmt.Println("Viimeisen commitin hash, jossa tagi", lastTaggedCommit.Hash)

	commitMessages, err := GetGitLogs(*lastTaggedCommit)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// Split commit messages by newline
	messages := strings.Split(strings.TrimSpace(commitMessages), "\n")

	// Create Markdown content
	mdContent := "# Changes since last version\n\n"
	for _, msg := range messages {
		mdContent += fmt.Sprintf("- %s\n", strings.Trim(msg, "'"))
	}
	return mdContent, nil
}

// Fetch all commit messages to the point where last tag is located.
func getCommitMessages() *object.Commit {
	// Open the existing repository
	repo, err := git.PlainOpen(".")
	if err != nil {
		log.Fatalf("Failed to open repository: %v", err)
	}

	// Get the latest tagged commit
	lastTaggedCommit, err := getLastTaggedCommit(repo)
	if err != nil {
		log.Fatalf("Failed to get last tagged commit: %v", err)
	}
	return lastTaggedCommit
}

func getLastTaggedCommit(repo *git.Repository) (*object.Commit, error) {
	tags, err := repo.Tags()
	if err != nil {
		return nil, err
	}

	var lastTaggedCommit *object.Commit

	err = tags.ForEach(func(ref *plumbing.Reference) error {
		tag, err := repo.TagObject(ref.Hash())
		if err != nil {
			// Not a tag object, it might be a lightweight tag
			commit, err := repo.CommitObject(ref.Hash())
			if err == nil {
				if lastTaggedCommit == nil || commit.Committer.When.After(lastTaggedCommit.Committer.When) {
					lastTaggedCommit = commit
				}
			}
			return nil
		}

		commit, err := repo.CommitObject(tag.Target)
		if err != nil {
			return err
		}

		if lastTaggedCommit == nil || commit.Committer.When.After(lastTaggedCommit.Committer.When) {
			lastTaggedCommit = commit
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	if lastTaggedCommit == nil {
		return nil, fmt.Errorf("no tagged commits found")
	}

	return lastTaggedCommit, nil
}
