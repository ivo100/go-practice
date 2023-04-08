package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// https://practice.geeksforgeeks.org/problems/find-the-missing-no-in-string/1
/*
Given a string consisting of some numbers, not separated by any separator.
The numbers are positive integers and the sequence increases by one at each number
except the missing number.
The task is to complete the function missingNumber which return's the missing number.
The numbers will have no more than six digits.
Print -1 if input sequence is not valid.

Note: Its is guaranteed that if the string is valid, then it is sure that at least
one number would be missing from the string.

Input:
The first line of input contains an integer T denoting the no of test cases. Then T test cases follow. Each test case contains an string s representing a number.

Output:
For each test case in a new line output will be the missing number. Output will be -1 if the input is invalid.

Constraints:
1<=T<=100
1<=Length of string<=100

Example(To be used only for expected output):
Input:
2
9899100102
1112141519

Output:
101
-1
*/

func check(s string, n int) bool {
	s1 := strconv.Itoa(n)
	return strings.HasPrefix(s, s1)
}

// try all lengths from 1 to 6
// let a is 1st number and b is 2nd
// check if a+1 or a+2 match b
// if yes - assume correct len found
func findFirst2(s string) (int, int) {
	// len can increase 98 99 100
	for l := 1; l <= 6; l++ {
		//a := s[0 : l+1]
		n, err := strconv.Atoi(s[0 : l+1])
		if err != nil {
			continue
		}
		if check(s[l+1:], n+1) {
			return n, n + 1
		}
		if check(s[l+1:], n+2) {
			return n, n + 1
		}
	}
	return -1, -1
}

func findMissing(s string) int {
	log.Printf("findMissing %v", s)
	n1, n2 := findFirst2(s)
	//log.Printf("findFirst2 %v %v", n1, n2)

	numMissing := 0
	missing := -1
	if n1 == n2-2 {
		numMissing = 1
		missing = n1 + 1
	}

	s2 := fmt.Sprintf("%d%d", n1, n2)
	s = s[len(s2):]
	nxt := n2 + 1
	for len(s) > 0 {
		// +1?
		if check(s, nxt) {
			snxt := strconv.Itoa(nxt)
			s = s[len(snxt):]
			nxt++
			continue
		}
		// +2?
		nxt++
		if check(s, nxt) {
			missing = nxt - 1
			//log.Printf("FOUND Missing %v", missing)
			numMissing++
			snxt := strconv.Itoa(nxt)
			s = s[len(snxt):]
			//return nxt
			// keep going - must have only 1 missing
		}
	}
	if numMissing > 1 {
		missing = -1
	}
	log.Printf("FOUND Missing %v", missing)
	return missing
}

func main() {
	s := "9899100102"
	n := findMissing(s)
	if n != 101 {
		panic("expected 101")
	}
	s = "1112141519"
	n = findMissing(s)
	if n != -1 {
		panic("expected -1")
	}
}
