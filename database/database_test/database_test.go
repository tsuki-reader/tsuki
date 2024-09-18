package database_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"tsuki/database"
)

var _ = Describe("Database", func() {
	Describe("Connect", func() {
		It("successfully connects to the database", func() {
			Expect(database.DATABASE).To(BeNil())
			Expect(database.Connect).NotTo(Panic())
			Expect(database.DATABASE).NotTo(BeNil())
		})
	})
})
