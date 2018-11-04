package factorization_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/JulzDiverse/al-go-rithms/numbers/factorization"
)

var _ = Describe("Factor", func() {
	Context("FactorByTrialDivison", func() {
		It("should return all factors of a number", func() {
			factors := FactorsByTrialDivision(6)
			Expect(factors).To(ConsistOf(1, 2, 3, 6))
		})
	})

	Context("FactorByTrialDivison Improved", func() {
		It("should return all factors of a number", func() {
			factors := FactorsByTrialDivision(12)
			Expect(factors).To(ConsistOf(1, 2, 3, 4, 6, 12))
		})
	})

	Context("Better Factorization", func() {
		It("should return all factors of a number", func() {
			factors := FactorsUnsorted(12)
			Expect(factors).To(ConsistOf(1, 2, 3, 4, 6, 12))
		})

		It("should return all factors of a number", func() {
			factors := FactorsUnsorted(36)
			Expect(factors).To(ConsistOf(1, 2, 3, 4, 6, 9, 12, 18, 36))
		})
	})
})
