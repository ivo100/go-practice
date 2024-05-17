package main

import "log"

/*
	Kadane algorithm
	https://www.geeksforgeeks.org/largest-sum-contiguous-subarray/
*/

func main() {
	//p := []int{7, 1, 5, 3, 6, 4}
	//p := []int{100, 180, 260, 310, 40, 535, 695}
	p := []int{23, 13, 25, 29, 33, 19, 34, 45, 65, 67}

	profit := buy_sell_a_stock(p)
	_ = profit
}

func buy_sell_a_stock(p []int) int {
	log.Printf("stock prices: %v", p)
	buy := p[0]
	profit := 0
	sold := 0
	daySold := 0
	for day := 1; day < len(p); day++ {
		diff := p[day] - buy
		if diff < 0 {
			buy = p[day]
			profit = 0
			log.Printf("Day %v, bought at $%v", day, buy)
			continue
		}
		if diff > profit {
			profit = diff
			sold = p[day]
			daySold = day
		}
	}
	if profit > 0 {
		log.Printf("Day %v, sold at $%v", daySold, sold)
		log.Printf("max profit $%v", profit)
	}
	return profit
}
