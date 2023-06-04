package main

import (
	"fmt"
	"strings"
)

// 3 roses
const rrr = "000"

// 1 rose, 1 cosmos
const rc = "01"

// p is price for rrr
// q is price for rc
func flowerBouquets(p int, q int, s string) int {
	// Write your code here
	rev := 0
	// greedy
	fmt.Printf("input: %s, p %d, q % d\n", s, p, q)
	ss := s
	for len(ss) > 0 {
		fmt.Println(ss)
		if p > q {
			i := strings.Index(ss, rrr)
			if i >= 0 {
				rev += p
				ss = ss[0:i] + "***" + ss[i+3:]
			} else {
				i := strings.Index(ss, rc)
				if i >= 0 {
					rev += q
					ss = ss[0:i] + "** " + ss[i+2:]
				} else {
					break
				}
			}
		} else {
			i := strings.Index(ss, rc)
			if i >= 0 {
				rev += p
				ss = ss[0:i] + "**" + ss[i+2:]
			} else {
				i := strings.Index(ss, rrr)
				if i >= 0 {
					rev += q
					ss = ss[0:i] + "***" + ss[i+3:]
				} else {
					break
				}
			}
		}
	}
	return rev
}

func main() {
	n := flowerBouquets(3, 2, "0001000")
	// expect 5
	fmt.Printf("n %d\n", n)
}
