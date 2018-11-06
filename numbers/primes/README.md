# Primes Package

Algorithms related to Prime Numbers.

## Determine if a number is a Prime number

--> `prime.go`

### Algorithm 1: Trial Divison

- Obvious approach
- Time Complexity: `O(n)`

*Idea:* Divide the given number `N` with every number smaller then N. If N is dividalbe by any smaller number it is not a prime number.

### Algorithm 2: Trial Divison Improved

- Better approach (still simple)
- Time Complexity: `O(SquareRoot(n))`

*Idea:* Same approach as `Algorithm 1` with the difference that we only interate up to the square root of the input number `N`. A number can't have any factors larger than the square root except the number itself.

## Find all Primes up to a given number N

--> `sieve.go`

### Algorithm 1: Trial Division

- Obvious Approach
- Time Complexity: `O(n*SquareRoot(n))`

*Idea:* For every number 2 to N check if the number is a prime. For the check use the Improved Trial Division algorithm. 

### Algorithm 2: Sieve of Erators Thenes 

- Smart Approach
- Time Compleixty: `O(n*log log(n)) + O(n)`

*Idea:* Consider every number up to N a prime number. For every prime number remove its multipliers.
