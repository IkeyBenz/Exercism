package diffsquares

// SquareOfSum takes an integer, n, and returns the sum of all numbers from 0-n squared
func SquareOfSum(n int) int {
	return ((n * (n + 1)) / 2) * ((n * (n + 1)) / 2)
}

// SumOfSquares takes an integer, n, and returns the sum of squares of numbers from 0-n
func SumOfSquares(n int) int {
	// An^3 + Bn^2 + Cn
	return 1
}

// Difference returns the difference between SquareOfSum(n) and SumOfSquares(n)
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
