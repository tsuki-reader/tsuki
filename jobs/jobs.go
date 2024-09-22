package jobs

import "time"

func Run() {
	refreshManga := time.NewTicker(10 * time.Minute)

	go func() {
		for range refreshManga.C {
			RefreshMangaJob()
		}
	}()
}
