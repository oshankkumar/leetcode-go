package container_with_most_water

func maxAreaBruteForce(height []int) int {
	var maxArea int

	for i := 0; i < len(height); i++ {
		for j := i; j < len(height); j++ {
			area := min(height[j], height[i]) * (j - i)
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

func maxArea(height []int) int {
	var maxArea int
	i, j := 0, len(height)-1

	for i < j {
		area := (j - i) * min(height[i], height[j])
		if area > maxArea {
			maxArea = area
		}
		if height[j] > height[i] {
			i++
		} else {
			j--
		}
	}
	return maxArea
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
