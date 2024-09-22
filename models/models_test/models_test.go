package models_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"tsuki/models"
)

var _ = Describe("Models", func() {
	Describe("Manga", func() {
		Describe("ParsedGenres", func() {
			Context("when there are no genres", func() {
				It("returns an empty array", func() {
					manga := models.Manga{}
					Expect(manga.ParsedGenres()).To(Equal([]string{}))
				})
			})

			Context("when there is one genre", func() {
				It("returns an array with one genre", func() {
					manga := models.Manga{Genres: "Action"}
					Expect(manga.ParsedGenres()).To(Equal([]string{"Action"}))
				})
			})

			Context("when there are multiple genres", func() {
				It("returns an array of genres", func() {
					manga := models.Manga{Genres: "Action,Adventure"}
					Expect(manga.ParsedGenres()).To(Equal([]string{"Action", "Adventure"}))
				})
			})
		})
	})
})
