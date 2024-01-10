package besttimetobuyandsellstock

func maxProfit(prices []int) int {
	var maxProfit int

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxProfit += prices[i] - prices[i-1]
		}
	}

	return maxProfit
}
