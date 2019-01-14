package numbers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/JulzDiverse/al-go-rithms/numbers/gcd"
)

var _ = Describe("Gcd", func() {

	Context("For two given numbers", func() {
		Context("When the numbers have common divisors", func() {
			It("should find the greatest divisor", func() {
				a, b := 30, 48
				Expect(Gcd(a, b)).To(Equal(6))
			})
		})

		Context("When the numbers dones not have other common divisors than 1", func() {
			It("should find the greatest divisor", func() {
				a, b := 2, 3
				Expect(Gcd(a, b)).To(Equal(1))
			})
		})

		Context("When one of the two numbers is a zero", func() {
			It("should find the greatest divisor", func() {
				a, b := 30, 0
				Expect(Gcd(a, b)).To(Equal(30))
			})
		})
	})
})
