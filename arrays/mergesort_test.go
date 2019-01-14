package arrays_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/JulzDiverse/al-go-rithms/arrays"
	"github.com/JulzDiverse/al-go-rithms/helpers"
)

var _ = Describe("Mergesort", func() {
	Context("When an array is not sorted", func() {
		It("should be sorted", func() {
			sorted := MergeSort([]int{2, 1, 3, 6, 4, 7, 5})
			Expect(helpers.StringifySl(sorted)).To(Equal("1 2 3 4 5 6 7"))
		})
	})
})
