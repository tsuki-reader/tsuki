package anilist_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAnilist(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Anilist Suite")
}
