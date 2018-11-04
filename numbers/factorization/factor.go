package factorization

import (
	"math"

	"github.com/JulzDiverse/al-go-rithms/helpers"
)

/*
Problem:
Given a number N, find all factors of N.
*/

// Easy approach: O(n) -> linear time complexity
func FactorsByTrialDivision(number int) []int {
	factors := []int{}
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}

// Improved approach: O(n) -> still linear time complexity
func FactorsByTrialDivisionImproved(number int) []int {
	factors := []int{1, number}
	for i := 2; i <= number/2; i++ {
		if number%i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}

// Better Approach: T propotional to square root of N -> O(SQRT(N))
func FactorsUnsorted(num int) []int {
	factors := []int{}
	for i := 1; i <= helpers.SquareRootOf(num); i++ {
		if num%i == 0 {
			factors = append(factors, i)
			if float64(i) != math.Sqrt(float64(num)) {
				factors = append(factors, num/i)
			}
		}
	}
	return factors
}
