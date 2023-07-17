package main

import "log"

/*
Determine if the sum of 2 integers is equal to the given value
Problem statement
*/

func Find2Sum(a []int, x int) bool {
	m := make(map[int]int)
	l := len(a)
	log.Printf("a %v, find %d", a, x)
	for i := 0; i < l; i++ {
		n := a[i]
		s := x - n
		//log.Printf("i %d, s %d", i, s)
		m[s] = 1
		if _, ok := m[n]; ok {
			log.Printf("found %d, m %v", n, m)
			return true
		}
	}
	return false
}
