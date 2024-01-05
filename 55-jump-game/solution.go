package jumpgame

func canJump(nums []int) bool {
	var maxSoFar int
	for i := 0; i <= maxSoFar && maxSoFar < len(nums); i++ {
		maxSoFar = max(maxSoFar, i+nums[i])
	}
	return maxSoFar >= len(nums)-1
}
