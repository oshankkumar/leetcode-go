package minimum_array_length_after_pair_removals

func minLengthAfterRemovals(nums []int) int {
	left, right := 0, len(nums)-1
	k := 0
	for left < right {
		if nums[left] == nums[right] {
			break
		}
		left++
		right--
		k++
	}
	return len(nums) - 2*k
}
