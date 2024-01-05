package jumpgame

func jump(nums []int) int {
	var right, maxSoFar, jump int

	for left := 0; left < len(nums)-1; left++ {
		maxSoFar = max(maxSoFar, left+nums[left])
		if left == right {
			right = maxSoFar
			jump++
		}
	}

	return jump
}
