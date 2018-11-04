package primes_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/JulzDiverse/al-go-rithms/numbers/primes"
)

var _ = Describe("Prime", func() {
	Context("Get Prime By Trial Divison", func() {
		It("should return true if number is a prime", func() {
			isPrime := PrimeByTrialDivision(7)
			Expect(isPrime).To(BeTrue())
		})

		It("should return false if number is not a prime", func() {
			isPrime := PrimeByTrialDivision(4)
			Expect(isPrime).ToNot(BeTrue())
		})
	})

	Context("Get Prime By Trial Division Improved", func() {
		It("should return true if number is a prime", func() {
			isPrime := PrimeByTrialDivisionImproved(7)
			Expect(isPrime).To(BeTrue())
		})

		It("should return false if number is not a prime", func() {
			isPrime := PrimeByTrialDivisionImproved(4)
			Expect(isPrime).ToNot(BeTrue())
		})

		It("should return true if a very large number is a prime", func() {
			isPrime := PrimeByTrialDivisionImproved(1500450271)
			Expect(isPrime).To(BeTrue())
		})
	})
})
