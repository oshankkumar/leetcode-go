package besttimetobuyandsellstock

func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}

	if len(prices) == 2 {
		return max(0, prices[1]-prices[0])
	}

	minSoFar := prices[0]
	var maxmProfit int
	maxProfitTillDay := make([]int, len(prices))

	for i, p := range prices[1:] {
		minSoFar = min(minSoFar, p)
		maxmProfit = max(maxmProfit, p-minSoFar)
		maxProfitTillDay[i] = maxmProfit
	}

	maxSoFar := prices[len(prices)-1]
	for i := len(prices) - 2; i >= 1; i-- {
		maxSoFar = max(maxSoFar, prices[i])
		maxmProfit = max(maxmProfit, maxProfitTillDay[i-1]+maxSoFar-prices[i])
	}

	return maxmProfit
}
