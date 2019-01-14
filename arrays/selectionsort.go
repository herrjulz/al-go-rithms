package arrays

func SelectionSort(sl []int) []int {
	for i := 0; i < len(sl); i++ {
		for j := i + 1; j < len(sl); j++ {
			if sl[j] < sl[i] {
				hold := sl[i]
				sl[i] = sl[j]
				sl[j] = hold
			}
		}
	}
	return sl
}
