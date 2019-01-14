package arrays_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/JulzDiverse/al-go-rithms/arrays"
	"github.com/JulzDiverse/al-go-rithms/helpers"
)

var _ = Describe("Selectionsort", func() {

	Context("When an slice is unsorted", func() {

		var (
			sl     []int
			sorted []int
		)

		It("should sort the slice", func() {
			sl = []int{3, 1, 4, 2, 5, 6, 9, 8}
			sorted = SelectionSort(sl)
			Expect(helpers.StringifySl(sorted)).To(Equal("1 2 3 4 5 6 8 9"))
		})

		It("should sort slices that contain the same value twice or more", func() {
			sl = []int{4, 1, 4, 4, 5, 6, 9, 8}
			sorted = SelectionSort(sl)
			Expect(helpers.StringifySl(sorted)).To(Equal("1 4 4 4 5 6 8 9"))
		})
	})
})
