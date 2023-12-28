package sortcolors

func sortColors(nums []int) {
	pivot := 1
	sm, eq, lg := 0, 0, len(nums)-1

	for eq <= lg {
		switch {
		case nums[eq] < pivot:
			nums[sm], nums[eq] = nums[eq], nums[sm]
			sm++
			eq++
		case nums[eq] == pivot:
			eq++
		case nums[eq] > pivot:
			nums[lg], nums[eq] = nums[eq], nums[lg]
			lg--
		}
	}
}
