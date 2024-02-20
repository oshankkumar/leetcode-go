package majorityelement

func majorityElement(nums []int) int {
	freq := make(map[int]int)

	currMajor := nums[0]
	for _, n := range nums {
		freq[n]++
		if freq[n] > freq[currMajor] {
			currMajor = n
		}
	}

	return currMajor
}

func majorityElementMooreAlgo(nums []int) int {
	count, major := 0, nums[0]

	for _, curr := range nums {
		switch {
		case curr == major:
			count++
		case count == 0:
			major = curr
		default:
			count--
		}
	}

	return major
}
