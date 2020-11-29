package main

import "fmt"

/* func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxProfit(price []int, start int, end int) int {
	if start >= end {
		return 0
	}

	profit := 0
	for i := start; i < end; i++ {
		for j := i + 1; j <= end; j++ {
			if price[j] > price[i] {
				currProfit := price[j] - price[i] + maxProfit(price, start, i-1) + maxProfit(price, j+1, end)
				profit = max(profit, currProfit)
			}
		}
	}

	return profit
} */

func maxProfit(price []int) int {
	profit := 0
	for i := 1; i < len(price); i++ {
		if price[i] > price[i-1] {
			profit += price[i] - price[i-1]
		}
	}

	return profit
}

func main() {
	arr := []int{8, 2, 13, 7, 1}
	fmt.Println(maxProfit(arr))
}
