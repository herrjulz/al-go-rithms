package basesystems_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/JulzDiverse/al-go-rithms/numbers/basesystems"
)

var _ = Describe("Binary", func() {
	Context("Decimal to Binary", func() {
		It("should return the binary number for a int", func() {
			bin := DecimalToBinary(357)
			Expect(bin).To(ConsistOf(1, 0, 1, 1, 0, 0, 1, 0, 1))
		})
	})
})
