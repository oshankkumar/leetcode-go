package two_sum_ii_input_array_is_sorted

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		}
		if sum > target {
			right--
		} else {
			left++
		}
	}
	return []int{0, 0}
}
