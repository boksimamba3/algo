package main

import "fmt"

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxProfit(prices []int, start int, end int) int {
	if start >= end {
		return 0
	}

	profit := 0
	for i := start; i < end; i++ {
		for j := i + 1; j <= end; j++ {
			if prices[j] > prices[i] {
				currProfit := prices[j] - prices[i] + maxProfit(prices, start, i-1) + maxProfit(prices, j+1, end)
				profit = max(profit, currProfit)
			}
		}
	}

	return profit
}

func main() {
	arr := []int{8, 2, 13, 7, 1}
	fmt.Println(maxProfit(arr, 0, len(arr)-1))
}
