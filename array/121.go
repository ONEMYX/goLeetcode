package array

func maxProfit(prices []int) int {
	minNum := prices[0]
	ans := 0
	for i := range prices {
		ans = max(ans, prices[i]-minNum)
		minNum = min(minNum, prices[i])
	}

	return ans
}
