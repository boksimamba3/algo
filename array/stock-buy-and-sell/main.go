package main

import "fmt"

// Question: Stock buy and sell. We need to find when to
// buy and when to sell stocks to get maximum profit. We
// assume that we know prices of stock for upcoming days.
// We can buy or sell stock on any day.

// Solution: O(N)
/*
	We buy stock when they are at lowest. If the stock are
	rising we keep then until they reach the peek. When
	they have reached the peek we sell them and we will buy
	them again when they start rising. On every bottom point buy
	stock and on every peek sell the stock.

	In = {1, 5, 3, 1, 2, 8}, Out: 11
	profit = 0
	I1: i = 1, 5 > 1, profit = 0 + 5 - 1 = 4;
	I2: i = 2, 3 < 5, profit = 4;
	I3: i = 3, 1 < 3, profit = 4;
	I4: i = 4, 2 > 1, profit = 4 + 2 - 1 = 5;
	I5: i = 5, 8 > 2, profit = 5 + 8 - 2 = 11;
*/
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
	fmt.Println(maxProfit([]int{8, 2, 13, 7, 1}))
}
