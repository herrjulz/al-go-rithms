package primes

import "github.com/JulzDiverse/al-go-rithms/helpers"

/*
A prime number is a number that is devisible exactly by two disticts
positive numbers 1 and the number itself.
*/

// Simple Approach: Linear Time Complexity -> O(n)
func PrimeByTrialDivision(number int) bool {
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return number > 1
}

// Better Approach: Time Complexity -> O(n)
func PrimeByTrialDivisionImproved(number int) bool {
	for i := 2; i <= helpers.SquareRootOf(number); i++ {
		if number%i == 0 {
			return false
		}
	}
	return number > 1
}
