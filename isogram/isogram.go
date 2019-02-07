package isogram

import "strings"

// IsIsogram takes a word and returns whether or not it is an isogram
func IsIsogram(word string) bool {

	previous := make(map[rune]bool)

	for _, char := range strings.ToLower(word) {
		if char != '-' && char != ' ' {
			if previous[char] {
				return false
			}
			previous[char] = true
		}
	}

	return true

}
