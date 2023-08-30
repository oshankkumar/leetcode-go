package uglynumberii

func nthUglyNumber(n int) int {
	if 1 <= n && n <= 6 {
		return n
	}

	num := 6
	count := 6
	for i := 7; count < n; i++ {
		if !isUgly(i) {
			continue
		}

		num = i
		count++
	}

	return num
}

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}

	for n%2 == 0 {
		n /= 2
	}

	for n%3 == 0 {
		n /= 3
	}

	for n%5 == 0 {
		n /= 5
	}

	return n == 1
}
