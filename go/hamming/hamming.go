package hamming

import "errors"

// Distance takes two strings and returns the 'hamming' distance between them
// Returns the number of different characters in the string and/or an error
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return -1, errors.New("both strands must have the same length")
	}

	differences := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			differences++
		}
	}

	return differences, nil

}
