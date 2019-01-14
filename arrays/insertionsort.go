package arrays

/*
	Time Complexity is O(n^2)
*/
func InsertionSort(sl []int) []int {
	for i := 1; i < len(sl); i++ {
		key := sl[i]
		left := i - 1

		for left >= 0 && sl[left] > key {
			sl[left+1] = sl[left]
			left--
		}

		sl[left+1] = key
	}
	return sl
}
