package numbers

func GcdV1(a, b int) int {
	if a == 0 || b == 0 {
		if b != 0 {
			return b
		}
		if a != 0 {
			return a
		}
	}

	if a == b {
		return a
	}

	gcd := 1
	smaller := a
	bigger := b

	if a >= b {
		smaller = b
		bigger = a
	}

	if bigger%smaller == 0 {
		return smaller
	}

	for i := 2; i <= smaller/2; i++ {
		if smaller%i == 0 {
			if bigger%i == 0 {
				gcd = i
			}
		}
	}

	return gcd
}

func Gcd(a, b int) int {
	if a >= b {
		a, b = swap(a, b)
	}

	if a == 0 {
		return b
	}

	return Gcd(a, b%a)
}

func swap(a, b int) (int, int) {
	hold := a
	a = b
	b = hold
	return a, b
}
