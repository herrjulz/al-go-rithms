package basesystems

func DecimalToBinary(n int) []int {
	digits := []int{}
	for n > 0 {
		digits = append(digits, n%2)
		n = n / 2
	}
	return digits
}
