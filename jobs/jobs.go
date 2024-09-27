package jobs

import "time"

func Run() {
	go func() {
		// Run jobs on start up
		UpdateDefaultRepoJob()
		RefreshMangaJob()
	}()

	refreshManga := time.NewTicker(10 * time.Minute)
	updateTsukiRepo := time.NewTicker(2 * time.Hour)

	go func() {
		for range refreshManga.C {
			RefreshMangaJob()
		}
	}()

	go func() {
		for range updateTsukiRepo.C {
			UpdateDefaultRepoJob()
		}
	}()
}
