package count_primes

func countPrimes(n int) int {
	notPrimes := make([]bool, n)

	var count int

	for i := 2; i < n; i++ {
		if notPrimes[i] {
			continue
		}

		count++

		for j := i * i; j < n; j += i {
			notPrimes[j] = true
		}

	}

	return count
}
