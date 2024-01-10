package besttimetobuyandsellstock

func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}

	minPriceSoFar := prices[0]
	var maxProfit int

	for _, p := range prices[1:] {
		minPriceSoFar = min(minPriceSoFar, p)
		maxProfit = max(maxProfit, p-minPriceSoFar)
	}

	return maxProfit
}
