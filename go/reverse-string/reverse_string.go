package reverse

// String takes a string and returns a reversed versio of it

// This is not the optimal algorithm for this task ...
func String(s string) string {
	reversed := ""
	for _, char := range s {
		reversed = string(char) + reversed
	}
	return reversed
}