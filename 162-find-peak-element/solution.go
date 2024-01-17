package find_peak_element

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return 0
		}
		return 1
	}

	return findPeakElementUtil(nums, 0, len(nums)-1)
}

func findPeakElementUtil(nums []int, left, right int) int {
	if left >= right && left > 0 && right < len(nums)-1 {
		return -1
	}

	mid := (left + right) / 2
	if mid == 0 {
		if nums[mid] > nums[mid+1] {
			return mid
		}
		return findPeakElementUtil(nums, mid+1, right)
	}

	if mid == len(nums)-1 {
		if nums[mid] > nums[mid-1] {
			return mid
		}
		return findPeakElementUtil(nums, left, mid)
	}

	if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
		return mid
	}

	if idx := findPeakElementUtil(nums, left, mid); idx != -1 {
		return idx
	}

	return findPeakElementUtil(nums, mid+1, right)
}
