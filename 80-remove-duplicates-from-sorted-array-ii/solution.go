package remove_duplicates_from_sorted_array_ii

func removeDuplicates(nums []int) int {
	i := 0
	unqCnt := 0

	for j := 0; j < len(nums); j++ {
		if nums[i] == nums[j] {
			unqCnt++
			continue
		}

		if unqCnt < 2 {
			i++
		} else {
			nums[i+1] = nums[i]
			i += 2
		}

		nums[i] = nums[j]
		unqCnt = 1
	}

	if unqCnt < 2 {
		return i + 1
	}
	nums[i+1] = nums[i]
	return i + 2
}
