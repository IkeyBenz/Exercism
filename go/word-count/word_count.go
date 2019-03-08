package wordcount

import (
	"strings"
	"unicode"
)

//Frequency is just a dictionary of string:int key/val pairs
type Frequency map[string]int

//WordCount returns a map of the frequency of words in the sentence.
func WordCount(sentence string) Frequency {
	counts := make(Frequency)
	words := strings.FieldsFunc(sentence, split)
	for _, word := range words {
		counts[strings.ToLower(strings.Trim(word, "'"))]++
	}
	return counts
}

func split(r rune) bool {
	return unicode.IsSpace(r) || r == '\n' || (unicode.IsPunct(r) && r != '\'') || unicode.IsSymbol(r)
}
