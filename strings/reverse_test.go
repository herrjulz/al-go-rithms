package strings_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/JulzDiverse/al-go-rithms/strings"
)

var _ = Describe("Reverse", func() {
	Context("Reverse slice", func() {
		It("should reverse a string slice", func() {
			Expect(Reverse("julz")).To(Equal("zluj"))
		})
	})
})
