package unique_number_of_occurrences

func uniqueOccurrences(arr []int) bool {
	numFreq := make(map[int]int)
	for _, n := range arr {
		numFreq[n]++
	}

	occurSet := make(map[int]struct{})

	for _, f := range numFreq {
		if _, ok := occurSet[f]; ok {
			return false
		}
		occurSet[f] = struct{}{}
	}

	return true
}
