// To execute Go code, please declare a func main() in a package "main"

/*
Given a binary array A[] of size N. The task is to arrange array in increasing order.

Note: The binary array contains only 0  and 1.

Input:
The first line of input contains an integer T, denoting the testcases. Every test case contains two lines, first line is N(size of array) and second line is space separated elements of array.

Output:
Space separated elements of sorted arrays. There should be a new line between output of every test case.

Constraints:
1 < = T <= 100
1 <= N <= 106
0 <= A[i] <= 1

Example:
Input:
2
5
1 0 1 1 0
10
1 0 1 1 1 1 1 0 0 0

Output:
0 0 1 1 1
0 0 0 0 1 1 1 1 1 1
*/
package main

import "fmt"

func main() {

	a := []uint8{1, 0, 1, 1, 0}
	sort(a)
	fmt.Printf("a %+v\n", a)

}

func sort(a []uint8) {
	fmt.Printf("a %+v\n", a)

	i := 0
	j := len(a) - 1
	for i < j {
		if a[i] == 0 {
			i++
			continue
		}
		if a[j] == 1 {
			j--
			continue
		}
		// a[i] 1, a[j] 0 - swap them
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
}
