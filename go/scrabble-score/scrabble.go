package scrabble

import "unicode"

// Score takes a word and returns its scrabble score
func Score(word string) int {

	score := 0

	for _, char := range word {

		switch unicode.ToUpper(char) {

		case 'D', 'G':
			score += 2
			break
		case 'B', 'C', 'M', 'P':
			score += 3
			break
		case 'F', 'H', 'V', 'W', 'Y':
			score += 4
			break
		case 'K':
			score += 5
			break
		case 'J', 'X':
			score += 8
			break
		case 'Q', 'Z':
			score += 10
			break
		default:
			score++
		}

	}

	return score
}
