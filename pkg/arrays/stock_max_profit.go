package main

import "fmt"

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
/*Say you have an array for which the ith element is the price of a given stock on day i.

If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock),
design an algorithm to find the maximum profit.

Note that you cannot sell a stock before you buy one.*/

func maxProfit(stocks []int) int {

	minPrice := stocks[0]
	maxProfit := 0

	for i:= 1; i < len(stocks); i++ {

		if stocks[i] < minPrice {
			minPrice = stocks[i]
		} else if stocks[i] - minPrice > maxProfit {
			maxProfit = stocks[i] - minPrice
		}
	}
	return maxProfit
}

func main() {
	stocks := []int{2, 3, 5, 7, 11, 13}

	maxProfit := maxProfit(stocks)
	fmt.Printf("Max profit : %d\n",maxProfit)

}
