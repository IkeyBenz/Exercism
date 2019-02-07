package raindrops

import "strconv"

// Convert takes a number and returns a weird string based on the numbers factors
func Convert(number int) string {

	factors := [3]int{3, 5, 7}
	strings := [3]string{"Pling", "Plang", "Plong"}

	s := ""
	for i, factor := range factors {
		if number%factor == 0 {
			s += strings[i]
		}
	}

	if s == "" {
		return strconv.Itoa(number)
	}

	return s
}
