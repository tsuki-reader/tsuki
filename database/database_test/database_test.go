package database_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"tsuki/database"
	"tsuki/test/mocks"
)

var mockLogger = &mocks.MockLogger{}

var _ = Describe("Database", func() {
	Describe("Connect", func() {
		BeforeEach(func() {
			mocks.BuildMockConfig(mockLogger, 2)
		})

		It("successfully connects to the database", func() {
			Expect(database.DATABASE).To(BeNil())
			Expect(database.Connect).NotTo(Panic())
			Expect(database.DATABASE).NotTo(BeNil())
		})

		It("logs the success message", func() {
			database.Connect()
			Expect(mockLogger.Args).To(Equal([]interface{}{"Database connected successfully"}))
		})

		Context("when there is an error connecting to the database", func() {
			It("logs a fatal error", func() {
				Skip("TODO: Mock an error")
			})
		})
	})
})
