package pqueue_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPqueue(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pqueue Suite")
}
