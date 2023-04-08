package main

import "log"

// https://practice.geeksforgeeks.org/problems/stock-buy-and-sell/0
/*
The cost of stock on each day is given in an array A[] of size N. Find all the days on which you buy and sell the stock so that in between those days your profit is maximum.

Input:
First line contains number of test cases T. First line of each test case contains an integer value N denoting the number of days, followed by an array of stock prices of N days.

Output:
For each testcase, output all the days with profit in a single line. And if there is no profit then print "No Profit".

Constraints:
1 <= T <= 100
2 <= N <= 103
0 <= Ai <= 104

Example
Input:
2
7
100 180 260 310 40 535 695
10
23 13 25 29 33 19 34 45 65 67

Output:
(0 3) (4 6)
(1 4) (5 9)

Explanation:
Testcase 1: We can buy stock on day 0, and sell it on 3rd day, which will give us maximum profit.

Note: Output format is as follows - (buy_day sell_day) (buy_day sell_day)
For each input, output should be in a single line.
*/
func main() {
	var prof int
	p := []int{100, 180, 260, 310, 40, 535, 695}
	prof = stocks(p)
	log.Print("-----")
	p = []int{23, 13, 25, 29, 33, 19, 34, 45, 65, 67}
	prof = stocks(p)
	log.Print("-----")
	_ = prof
}

func stocks(p []int) int {
	l := len(p)
	bal := 0
	i := 0
	j := 1
	profit := 0
	bought := 0
	for j < l && i < j {
		d := p[j] - p[j-1]
		//log.Printf("i %v, j %v, d %v", i, j, d)
		// up?
		if d > 0 {
			if bal == 0 {
				// buy p[i]
				bal = p[i]
				bought = i
				log.Printf("day %v BUY at %v, tot.profit %v", bought, bal, profit)
			}
			// next day
			j++
			if j < l-1 {
				continue
			}
			// last day - sell
			if bal > 0 {
				pr := p[j] - p[bought]
				profit += pr
				log.Printf("day %v SELL at %v, profit %v, tot.prof %v", j, p[j], pr, profit)
				return profit
			}
		}
		// down?
		if d < 0 {
			if bal > 0 {
				pr := p[j-1] - p[bought]
				if pr > 0 {
					// time to sell
					profit += pr
					log.Printf("day %v SELL at %v, profit %v, tot.prof %v", j-1, p[j-1], pr, profit)
					bal = 0
				}
				i = j
				j++
				continue
			}
			// next day
			i++
			j++
		}
	}
	return profit
}
