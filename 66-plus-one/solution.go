package plusone

func plusOne(digits []int) []int {
	digits[len(digits)-1]++
	var carry int
	if digits[len(digits)-1] == 10 {
		digits[len(digits)-1] = 0
		carry = 1
	}

	for i := len(digits) - 2; i >= 0 && carry > 0; i-- {
		digits[i] += carry
		if digits[i] < 10 {
			carry = 0
			break
		}

		digits[i] = 0
		carry = 1
	}

	if carry == 1 {
		digits = append([]int{1}, digits...)
	}

	return digits
}

func plusOneSolution2(digits []int) []int {
	digits[len(digits)-1]++

	for i := len(digits) - 1; i > 0 && digits[i] == 10; i-- {
		digits[i] = 0
		digits[i-1]++
	}

	if digits[0] == 10 {
		digits[0] = 0
		digits = append([]int{1}, digits...)
	}

	return digits
}
