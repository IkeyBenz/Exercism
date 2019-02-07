package diffsquares

// SquareOfSum takes an integer, n, and returns the sum of all numbers from 0-n squared
func SquareOfSum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum*sum
}

// SumOfSquares takes an integer, n, and returns the sum of squares of numbers from 0-n
func SumOfSquares(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i*i
	}
	return sum
}

// Difference returns the difference between SquareOfSum(n) and SumOfSquares(n)
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}