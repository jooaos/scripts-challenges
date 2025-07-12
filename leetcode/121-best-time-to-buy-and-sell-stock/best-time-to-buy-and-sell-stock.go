package main

func MaxProfit(prices []int) int {
	lowerValue := 0
	result := 0

	for i := 0; i < len(prices); i++ {
		if i == 0 {
			lowerValue = prices[0]
			continue
		}

		if prices[i] <= lowerValue {
			lowerValue = prices[i]
		} else if prices[i]-lowerValue > result {
			result = prices[i] - lowerValue
		}
	}

	return result
}
