package primes_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/JulzDiverse/al-go-rithms/numbers/primes"
)

var _ = Describe("Sieve", func() {
	Context("Find Primes", func() {
		It("should return all primes up to a given number N", func() {
			primes := FindPrimes(7)
			Expect(primes).To(ConsistOf(2, 3, 5, 7))
		})
	})

	Context("Sieve of Eratos Thenes", func() {
		It("should return all primes up to a given number N", func() {
			primes := Sieve(15)
			Expect(primes).To(ConsistOf(2, 3, 5, 7, 11, 13))
		})
	})
})
