package src

import "strings"

// WordBreakII is a function to return all possible sentences that can be formed by
// breaking the input string into words from the dictionary.
// The function takes a string s and a list of strings wordDict.
// The function returns all possible sentences that can be formed by breaking the input string into words from the dictionary.
func WordBreakII(s string, wordDict []string) []string {
	words := make(map[string]bool)
	sizes := make(map[int]bool)
	for _, word := range wordDict {
		words[word] = true
		sizes[len(word)] = true
	}

	output := make([]string, 0)
	var f func(array []string, idx int)

	f = func(array []string, idx int) {
		if idx == len(s) {
			output = append(output, strings.Join(array, " "))

			return
		}

		for size := range sizes {
			if idx+size > len(s) {
				continue
			}

			w := s[idx : idx+size]
			if _, ok := words[w]; !ok {
				continue
			}

			f(append(array, w), idx+size)
		}
	}

	f(nil, 0)

	return output
}
