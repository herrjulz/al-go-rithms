package basesystems

import (
	"math"
	"sort"
)

func FindDigitsInDecimal(n int) []int {
	n = int(math.Abs(float64(n)))
	digits := []int{}
	for n > 0 {
		remainder := n % 10
		digits = append(digits, remainder)
		n = n / 10
	}
	sort.Sort(sort.Reverse(sort.IntSlice(digits)))
	return digits
}
