package primes

import "github.com/JulzDiverse/al-go-rithms/helpers"

/*
Problem:
Given a number N, find all prime numbers N (N included)
*/

//Simple and obvious approach, with time complexity -> O(n*SQRT(n))
func FindPrimes(number int) []int {
	primes := []int{}
	for i := 2; i <= number; i++ {
		if PrimeByTrialDivisionImproved(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

/*
  Better approach with time complexity -> O(n*log log n) + O(n)
  Approach: Eliminate all multipliers of a prime
*/
func Sieve(number int) []int {
	primes := make([]bool, number+1)
	primes[0], primes[1] = true, true
	primeCount := 0

	for i := 2; i <= helpers.SquareRootOf(number); i++ {
		if primes[i] == false {
			for j := 2; (i * j) <= number; j++ {
				primes[i*j] = true
				primeCount++
			}
		}
	}

	result := make([]int, 0, primeCount)
	for i, v := range primes {
		if v == false {
			result = append(result, i)
		}
	}

	return result
}
