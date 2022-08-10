package threesum

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int

	for i, num := range nums {
		// remove duplicates
		if i > 0 && num == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		target := -num
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				result = append(result, []int{num, nums[left], nums[right]})
				left++
				for left < right && nums[left] == nums[left-1] {
					left++
				}

			}
			if sum > target {
				right--
			}
			if sum < target {
				left++
			}
		}
	}

	return result
}
