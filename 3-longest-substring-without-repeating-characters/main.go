package longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	i, j := 0, 0
	maxLen := 0
	seen := make(map[byte]struct{})

	for j < len(s) {
		if _, ok := seen[s[j]]; ok {
			delete(seen, s[i])
			i++
			continue
		}
		seen[s[j]] = struct{}{}
		if len(seen) > maxLen {
			maxLen = len(seen)
		}
		j++
	}

	return maxLen
}
