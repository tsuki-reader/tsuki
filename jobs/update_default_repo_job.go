package jobs

import (
	"tsuki/core"
	"tsuki/extensions"
)

func UpdateDefaultRepoJob() {
	var repository extensions.Repository
	err := extensions.InstallRepository("https://raw.githubusercontent.com/tsuki-reader/tsuki-repo/refs/heads/main/repo.json", true, &repository)
	if err != nil {
		core.CONFIG.Logger.Println("Could not install/update tsuki-repo")
	}
}
