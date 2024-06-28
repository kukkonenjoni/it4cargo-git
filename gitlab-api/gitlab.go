package gitlabapi

import (
	"fmt"
	"os"

	"github.com/xanzy/go-gitlab"
)

func CreateRelease(projectId string, tagName string, releaseName string, description string) {
	gitlabToken := "glpat-M2UKZpFA5n8cYtk_qVKe"
	if gitlabToken == "" {
		fmt.Println("Please set the GITLAB_TOKEN environment variable")
		os.Exit(1)
	}

	// Create a new GitLab client
	git, err := gitlab.NewClient(gitlabToken, gitlab.WithBaseURL("https://gitlab.com/api/v4"))
	if err != nil {
		fmt.Printf("Failed to create client: %v\n", err)
		os.Exit(1)
	}

	// Fetch the latest commit hash of the repository
	latestCommit, _, err := git.Commits.ListCommits(projectId, &gitlab.ListCommitsOptions{
		ListOptions: gitlab.ListOptions{PerPage: 1},
	})
	if err != nil {
		fmt.Printf("Failed to fetch latest commit: %v\n", err)
		os.Exit(1)
	}

	if len(latestCommit) == 0 {
		fmt.Println("No commits found in the repository")
		os.Exit(1)
	}

	//latestCommitHash := latestCommit[0].ID
	branch := "master"

	// Create the release using the latest commit hash
	release, _, err := git.Releases.CreateRelease(projectId, &gitlab.CreateReleaseOptions{
		Name:        &releaseName,
		TagName:     &tagName,
		Description: &description,
		Ref:         &branch,
	})
	if err != nil {
		fmt.Printf("Failed to create release: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Release created: %v\n", release)
}
