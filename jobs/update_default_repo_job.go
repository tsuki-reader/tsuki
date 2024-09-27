package jobs

import (
	"tsuki/core"
	"tsuki/extensions"
)

func UpdateDefaultRepoJob() {
	err := extensions.InstallRepository("https://raw.githubusercontent.com/tsuki-reader/tsuki-repo/refs/heads/main/repo.json", true)
	if err != nil {
		core.CONFIG.Logger.Println("Could not install/update tsuki-repo")
	}
}
