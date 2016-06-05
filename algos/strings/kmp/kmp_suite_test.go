package kmp_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestKmp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Kmp Suite")
}
