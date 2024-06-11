package src

import "strings"

// ReplaceWords finds the root of the words in the dictionary and replaces them in the sentence.
// It returns the sentence with the roots of the words in the dictionary.
func ReplaceWords(dictionary []string, sentence string) string {
	type Trie struct {
		children [26]*Trie
		isWord   bool
	}

	insert := func(root *Trie, word string) {
		node := root
		for _, ch := range word {
			if node.children[ch-'a'] == nil {
				node.children[ch-'a'] = &Trie{}
			}
			node = node.children[ch-'a']
		}
		node.isWord = true
	}

	getRoot := func(root *Trie, word string) string {
		node := root
		for size, c := range word {
			if node.children[c-'a'] == nil {
				return word
			}

			node = node.children[c-'a']
			if node.isWord {
				return word[:size+1]
			}
		}

		return word
	}

	root := &Trie{}
	for _, word := range dictionary {
		insert(root, word)
	}

	words := strings.Split(sentence, " ")
	output := make([]string, 0, len(words))
	for _, word := range words {
		output = append(output, getRoot(root, word))
	}

	return strings.Join(output, " ")
}
