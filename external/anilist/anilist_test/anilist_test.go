package anilist_test

import (
	"tsuki/external/anilist"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func PurgeGlobals() {
	anilist.CLIENT = nil
	anilist.TOKEN = ""
}

var _ = Describe("Anilist", func() {
	// TODO: Add some mocking for the http client
	AfterEach(func() {
		PurgeGlobals()
	})

	Describe("SetupClient", func() {
		Context("when client is nil", func() {
			It("sets the client", func() {
				Expect(anilist.CLIENT).To(BeNil())
				anilist.SetupClient("this_is_a_token")
				Expect(anilist.CLIENT).NotTo(BeNil())
			})

			It("sets the token", func() {
				Expect(anilist.TOKEN).To(Equal(""))
				anilist.SetupClient("this_is_a_new_token")
				Expect(anilist.TOKEN).To(Equal("this_is_a_new_token"))
			})
		})

		Context("when client is not nil and given token is not the same as the current one", func() {
			BeforeEach(func() {
				anilist.SetupClient("InitialToken")
			})

			It("sets a new client", func() {
				previousClient := anilist.CLIENT
				Expect(previousClient).NotTo(BeNil())
				anilist.SetupClient("NewToken")
				Expect(anilist.CLIENT).NotTo(Equal(previousClient))
			})

			It("sets the new token", func() {
				Expect(anilist.TOKEN).To(Equal("InitialToken"))
				anilist.SetupClient("NewToken")
				Expect(anilist.TOKEN).To(Equal("NewToken"))
			})
		})

		Context("when client is not nil and given token is the same", func() {
			BeforeEach(func() {
				anilist.SetupClient("InitialToken")
			})

			It("does not reset the client", func() {
				previousClient := anilist.CLIENT
				Expect(previousClient).NotTo(BeNil())
				anilist.SetupClient("InitialToken")
				Expect(anilist.CLIENT).To(Equal(previousClient))
			})

			It("does not change the token", func() {
				Expect(anilist.TOKEN).To(Equal("InitialToken"))
				anilist.SetupClient("InitialToken")
				Expect(anilist.TOKEN).To(Equal("InitialToken"))
			})
		})
	})
})
