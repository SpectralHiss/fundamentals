package kmp

import (
	"fmt"
	"math"
)

type Searcher interface {
	Search(pattern string, text string, options Options) (present bool, positionsInText []int)
}

type kmpSearch struct{}

type Options struct {
	Debug         bool
	caseSensitive bool
}

func New() Searcher {
	return kmpSearch{}
}

func (kmp kmpSearch) Search(pattern string, text string, options Options) (bool, []int) {

	π := buildPi(pattern) // pi unicode character is U+03C0

	allMatches := []int{}
	hasMatches := false

	matched := 0
	for cursor := 0; cursor < len(text); {
		if matched == len(pattern) {
			allMatches = append(allMatches, cursor)
			matched = 0
			hasMatches = true
		}

		if text[cursor] == pattern[matched] {
			matched++
			cursor++
		} else {
			matched = π[matched]
			cursor += 1 + π[matched]
		}
	}

	if options.Debug {
		debug_output(allMatches, text, pattern)
	}

	return hasMatches, allMatches
}

func buildPi(pattern string) []int {
	// precompute table of longest prefix which is also a suffix

	π := make([]int, len(pattern))

	for index, _ := range pattern {
		for pre := 0; pre < index/2; pre++ {
			if pattern[pre] == pattern[index-pre] {
				π[index]++
			} else {
				break
			}
		}
	}

	return π
}

func debug_output(positions []int, corpus string, pattern string) {
	for _, position := range positions {
		min := int(math.Max(0, float64(position-100)))
		max := int(math.Min(float64(len(corpus)-1), float64(position+100)))

		fmt.Printf("\n !!! %pattern is found here:%#v", pattern, (corpus[min:max]))

	}
}
