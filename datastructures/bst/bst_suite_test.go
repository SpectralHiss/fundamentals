package bst_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestBst(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Bst Suite")
}
