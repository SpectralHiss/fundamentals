package kmp

type Searcher interface {
	Search(pattern string, text string) (present bool, positionsInText []int)
}

type kmpSearch struct{}

func New() Searcher {
	return kmpSearch{}
}

func (kmp kmpSearch) Search(pattern string, text string) (bool, []int) {

	π := buildPi(pattern) // pi unicode character is U+03C0

	matched := 0
	for cursor := 0; cursor < len(text); cursor++ {
		if matched == len(pattern) {
			return true, []int{cursor}
		}

		if text[cursor] == pattern[matched] {
			matched++
		} else {
			matched = 0
			cursor += len(pattern) - π[matched]
		}

	}

	return false, []int{}
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
