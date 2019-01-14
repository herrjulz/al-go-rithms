package basesystems_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/JulzDiverse/al-go-rithms/numbers/basesystems"
)

var _ = Describe("Decimal", func() {
	Context("Find Digits in a Decimal Number", func() {
		It("should return all single digits of a number", func() {
			digits := FindDigitsInDecimal(2345)
			Expect(digits).To(ConsistOf(2, 3, 4, 5))
		})

		It("should word for a single number", func() {
			digits := FindDigitsInDecimal(1)
			Expect(digits).To(ConsistOf(1))
		})

		It("should work for negative numbers", func() {
			digits := FindDigitsInDecimal(-125)
			Expect(digits).To(ConsistOf(1, 2, 5))
		})
	})
})
