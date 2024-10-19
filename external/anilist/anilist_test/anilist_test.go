package anilist_test

import (
	"tsuki/external/anilist"
	"tsuki/external/anilist/al_types"
	"tsuki/test/mocks"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var mockLogger = &mocks.MockLogger{}

var _ = Describe("Anilist", func() {
	// TODO: Figure out how to test the request vars if possible
	Describe("BuildAndSendRequest", func() {
		BeforeEach(func() {
			mocks.BuildMockConfig(mockLogger, 3)
		})

		Context("when query is not found", func() {
			It("logs fatal", func() {
				Expect(func() {
					anilist.BuildAndSendRequest[al_types.ALViewerData]("bogus", "altoken", nil)
				}).To(PanicWith([]interface{}{"Could not find Anilist query"}))
				Expect(mockLogger.FatalCalled).To(BeTrue())
			})
		})

		Context("when an error occurs sending the request", func() {
			It("returns nil and an error", func() {
				// File does not exist so will return an error. We don't really care *why* the error has been thrown,
				// we only care that the error is returned if and when one occurs.
				mockClient := &mocks.MockClient{ResponseFile: "../../../test/data/viewer_with_error.json"}
				resp, err := anilist.BuildAndSendRequest[al_types.ALViewerData]("viewer", "altoken", mockClient)
				Expect(resp).To(BeNil())
				Expect(err).NotTo(BeNil())
			})
		})

		Context("when the request succeeds", func() {
			It("returns the specified type", func() {
				mockClient := &mocks.MockClient{ResponseFile: "../../../test/data/viewer.json"}
				resp, err := anilist.BuildAndSendRequest[al_types.ALViewerData]("viewer", "altoken", mockClient)
				Expect(err).To(BeNil())
				Expect(resp.Viewer.Name).To(Equal("hooligan"))
				Expect(resp.Viewer.BannerImage).To(Equal("https://example.com/print.png"))
				Expect(resp.Viewer.Avatar.Large).To(Equal("https://example.com/large.png"))
				Expect(resp.Viewer.Avatar.Medium).To(Equal("https://example.com/medium.png"))
			})
		})
	})
})
