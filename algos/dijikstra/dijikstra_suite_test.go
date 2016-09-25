package dijikstra_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDijikstra(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dijikstra Suite")
}
