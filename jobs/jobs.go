package jobs

import "time"

func Run() {
	go func() {
		// Run jobs on start up
		RefreshMangaJob()
	}()

	refreshManga := time.NewTicker(10 * time.Minute)

	go func() {
		for range refreshManga.C {
			RefreshMangaJob()
		}
	}()
}
