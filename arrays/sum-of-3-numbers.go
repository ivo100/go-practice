package main

/*
https://www.codinginterview.com/apple-interview-questions

Determine if the sum of three integers is equal to the given value
Problem statement
Given an array of integers and a value, determine if there are any three integers in the array whose sum equals the given value.
*/

func Find3(a []int, x int) bool {
	m := make(map[int]int)
	l := len(a)
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			m[x-a[i]-a[j]] = 1
			if _, ok := m[a[j]]; ok {
				return true
			}
		}
	}
	return false
}
