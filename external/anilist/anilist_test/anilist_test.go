package anilist_test

import (
	"tsuki/external/anilist"
	"tsuki/test/mocks"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func PurgeGlobals() {
	anilist.CLIENT = nil
	anilist.TOKEN = ""
}

var mockLogger = &mocks.MockLogger{}

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

		Context("when client is not nil and token is an empty string", func() {
			BeforeEach(func() {
				anilist.SetupClient("InitialToken")
			})

			It("does not reset the token", func() {
				Expect(anilist.TOKEN).To(Equal("InitialToken"))
				anilist.SetupClient("")
				Expect(anilist.TOKEN).To(Equal("InitialToken"))
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

	// TODO: Figure out how to test the request vars if possible
	Describe("BuildAndSendRequest", func() {
		BeforeEach(func() {
			mocks.BuildMockConfig(mockLogger, 3)
		})

		Context("when query is not found", func() {
			It("logs fatal", func() {
				Expect(func() {
					anilist.BuildAndSendRequest[anilist.ViewerData]("bogus")
				}).To(PanicWith([]interface{}{"Could not find Anilist query"}))
				Expect(mockLogger.FatalCalled).To(BeTrue())
			})
		})

		Context("when an error occurs sending the request", func() {
			It("returns nil and an error", func() {
				// File does not exist so will return an error. We don't really care *why* the error has been thrown,
				// we only care that the error is returned if and when one occurs.
				mockClient := &mocks.MockClient{ResponseFile: "../../../test/data/viewer_with_error.json"}
				anilist.CLIENT = mockClient
				resp, err := anilist.BuildAndSendRequest[anilist.ViewerData]("viewer")
				Expect(resp).To(BeNil())
				Expect(err).NotTo(BeNil())
			})
		})

		Context("when the request succeeds", func() {
			It("returns the specified type", func() {
				mockClient := &mocks.MockClient{ResponseFile: "../../../test/data/viewer.json"}
				anilist.CLIENT = mockClient
				resp, err := anilist.BuildAndSendRequest[anilist.ViewerData]("viewer")
				Expect(err).To(BeNil())
				Expect(resp.Viewer.Name).To(Equal("hooligan"))
				Expect(resp.Viewer.BannerImage).To(Equal("https://example.com/print.png"))
				Expect(resp.Viewer.Avatar.Large).To(Equal("https://example.com/large.png"))
				Expect(resp.Viewer.Avatar.Medium).To(Equal("https://example.com/medium.png"))
			})
		})
	})
})
