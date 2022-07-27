package two_sum

func twoSum(nums []int, target int) []int {
	seenSet := make(map[int]int, len(nums))

	for i, num := range nums {
		needed := target - num
		foundAt, ok := seenSet[needed]
		if ok {
			return []int{foundAt, i}
		}
		seenSet[num] = i
	}

	return nil
}
