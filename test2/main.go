package main

import "math"

func main() {

}

func maxSellStock(stockPrices []int) int {
	max := math.MinInt32
	for i := 0; i < len(stockPrices); i++ {
		for j := i + 1; j < len(stockPrices); j++ {
			if stockPrices[j]-stockPrices[i] > max {
				max = stockPrices[j] - stockPrices[i]
			}
		}
	}
	return max
}
