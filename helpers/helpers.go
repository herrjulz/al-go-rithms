package helpers

import (
	"fmt"
	"math"
)

func SquareRootOf(n int) int {
	return int(math.Sqrt(float64(n)))
}

func StringifySl(sl []int) (result string) {
	for i, v := range sl {
		if i == 0 {
			result = fmt.Sprintf("%d", v)
			continue
		}
		result = fmt.Sprintf("%s %d", result, v)
	}
	return
}
