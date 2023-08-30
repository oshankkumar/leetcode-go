package findtargetindicesaftersortingarray

import (
	"sort"
)

func targetIndices(nums []int, target int) []int {
	var (
		eqTarget int
		ltTarget int
	)

	for _, n := range nums {
		if n == target {
			eqTarget++
		}
		if n < target {
			ltTarget++
		}
	}

	idxx := make([]int, 0, eqTarget)
	for i := 0; i < eqTarget; i++ {
		idxx = append(idxx, ltTarget+i)
	}

	return idxx
}

func targetIndicesUsingSort(nums []int, target int) []int {
	sort.Ints(nums)
	idxx := SearchInt(nums, 0, len(nums)-1, target, nil)
	sort.Ints(idxx)
	return idxx
}

func SearchInt(nums []int, start int, end int, target int, idxx []int) []int {
	if start > end {
		return idxx
	}

	if start == end {
		if nums[start] == target {
			idxx = append(idxx, start)
		}
		return idxx
	}

	mid := (start + end) / 2

	if nums[mid] == target {
		idxx = append(idxx, mid)
		idxx = SearchInt(nums, start, mid-1, target, idxx)
		idxx = SearchInt(nums, mid+1, end, target, idxx)
		return idxx
	}

	if nums[mid] > target {
		end = mid - 1
	}

	if nums[mid] < target {
		start = mid + 1
	}

	return SearchInt(nums, start, end, target, idxx)
}
