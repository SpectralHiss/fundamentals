package kmp_test

import (
	"io/ioutil"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/parallelKiller/fundamentals/algos/strings/kmp"
)

var _ = Describe("KMP exact string search", func() {
	var (
		corpus          string
		searcher        kmp.Searcher
		present         bool
		positions       []int
		pattern         string      = "Shaw"
		default_options kmp.Options = kmp.Options{Debug: false}
	)

	BeforeEach(func() {
		f, err := os.Open("../fixtures/corpus.txt")
		Expect(err).NotTo(HaveOccurred())

		corpusBytes, err := ioutil.ReadAll(f)
		Expect(err).NotTo(HaveOccurred())

		corpus = string(corpusBytes)
		searcher = kmp.New()
		present, positions = searcher.Search(pattern, corpus, default_options)
	})

	It("finds the location of a substring, if exists", func() {
		Expect(corpus).NotTo(BeNil())
		Expect(present).To(BeTrue())
		Expect(positions).NotTo(BeNil())
	})
})
