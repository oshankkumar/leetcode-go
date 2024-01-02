package multiplystrings

import "strings"

func multiply(s1 string, s2 string) string {
	num1, num2 := convertToIntSlice(s1, s2)
	result := make([]int32, len(num1)+len(num2))

	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			result[i+j+1] += num1[i] * num2[j]
			result[i+j] += result[i+j+1] / 10
			result[i+j+1] %= 10
		}
	}

	var i int
	for ; i < len(result) && result[i] == 0; i++ {
	}

	if i == len(result) {
		return "0"
	}

	var builder strings.Builder
	for _, n := range result[i:] {
		builder.WriteRune(n + '0')
	}
	return builder.String()
}

func convertToIntSlice(num1 string, num2 string) ([]int32, []int32) {
	n1, n2 := make([]int32, len(num1)), make([]int32, len(num2))

	for i, r := range num1 {
		n1[i] = r - '0'
	}

	for i, r := range num2 {
		n2[i] = r - '0'
	}

	return n1, n2
}
