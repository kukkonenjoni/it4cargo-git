package main

import (
	"fmt"
	"os"

	it4cargoGit "github.com/kukkonenjoni/it4cargo-git/commands"
	"github.com/kukkonenjoni/it4cargo-git/util"
)

func main() {
	argsWithoutProg := os.Args[1:]
	fmt.Println(argsWithoutProg[0])
	//Luo changelogin HEAD -> viimeisin tagi väliltä ja näistä commit messaget tulevat releasen kuvaukseen
	mdContent, err := it4cargoGit.CreateChangeLog()
	if err != nil {
		fmt.Println("Failed to create changelog")
		os.Exit(1)
	}
	fmt.Println(mdContent)

	//Load JSON config
	config, err := util.LoadConfig("config.json")
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Config: %+v\n", config)

	//Luo releasen
	//gitlabapi.CreateRelease(config.ProjectId, config.TagName, config.ReleaseName, mdContent)
	util.ChangeDirectory(config.Base)
	//Jos directoryä ei ole olemassa niin pullataan se
	for i := range config.Repositories {
		fmt.Println(config.Base + config.Repositories[i].FileName)
		if _, err := os.Stat(config.Base + config.Repositories[i].FileName); err != nil {
			if os.IsNotExist(err) {
				output, err := util.RunGitCommand("clone", "--depth", "1", config.Repositories[i].GitUrl)
				if err != nil {
					fmt.Println("Failed to clone repository")
					os.Exit(1)
				}
				fmt.Println(output)
			}
		}

		util.ChangeDirectory(config.Base + config.Repositories[i].FileName)
		output, err := util.RunGitCommand("submodule", "update", "--init", "--recursive")
		if err != nil {
			//Submodulea ei ole olemassa -> Luo submodule
			fmt.Println("Failed to update submodule, initializing new one")
			output2, err2 := util.RunGitCommand("submodule", "add", config.BaseRepo)
			if err2 != nil {
				fmt.Println("Unable to initialize new submodule")
				os.Exit(1)
			}
			fmt.Println(output2)
		}
		fmt.Println(output)

		err = util.ChangeDirectory(config.Base + config.Repositories[i].FileName + "/apha-inc")
		if err != nil {
			fmt.Println("Failed to change directory to submodule of " + config.Repositories[i].FileName)
			os.Exit(1)
		}
		output, err = util.RunGitCommand("checkout", "main")
		if err != nil {
			fmt.Println("Failed to checkout 'main' branch")
			os.Exit(1)
		}
		fmt.Println(output)

		output, err = util.RunGitCommand("pull", "origin", "main")
		if err != nil {
			fmt.Println("Failed to pull data to submodule from base repo and 'main' branch")
			os.Exit(1)
		}
		fmt.Println(output)

		err = util.ChangeDirectory(config.Base + config.Repositories[i].FileName)
		if err != nil {
			fmt.Println("Failed to change directory to submodule path of " + config.Repositories[i].FileName)
			os.Exit(1)
		}

		output, err = util.RunGitCommand("commit", "-am", "v1.0.0")
		if err != nil {
			fmt.Println(output)
			fmt.Println("Failed to create new commit")
			os.Exit(1)
		}

		fmt.Println(output)
		output, err = util.RunGitCommand("push")
		if err != nil {
			fmt.Println("Failed to push new submodule")
			os.Exit(1)
		}
		fmt.Println(output)
	}
}
