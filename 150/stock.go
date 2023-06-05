package main

// maxProfit 股票1
func maxProfit(prices []int) int {
	minPrice, ans := prices[0], 0
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		}
		if ans < price-minPrice {
			ans = price - minPrice
		}
	}
	return ans
}
