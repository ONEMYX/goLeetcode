package main

import "fmt"

func maxProfit(prices []int) int {
	mx := 0
	for i := 1; i < len(prices); i++ {
		mx += max(0, prices[i]-prices[i-1])
	}
	return mx
}

func main() {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := maxProfit(list)
	fmt.Println(s)
}
