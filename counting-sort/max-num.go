package main

import "fmt"

/*

https://www.geeksforgeeks.org/counting-sort/

For simplicity, consider the data in the range 0 to 9.
Input data: 1, 4, 1, 2, 7, 5, 2
  1) Take a count array to store the count of each unique object.
  Index:     0  1  2  3  4  5  6  7  8  9
  Count:     0  2  2  0  1  1  0  1  0  0

  2) Modify the count array such that each element at each index
  stores the sum of previous counts.
  Index:     0  1  2  3  4  5  6  7  8  9
  Count:     0  2  4  4  5  6  6  7  7  7

The modified count array indicates the position of each object in
the output sequence.

  3) Output each object from the input sequence followed by
  decreasing its count by 1.
  Process the input data: 1, 4, 1, 2, 7, 5, 2. Position of 1 is 2.
  Put data 1 at index 2 in output. Decrease count by 1 to place
  next data 1 at an index 1 smaller than this index.
*/

func countSort(a []int, min, max int) (res []int) {
	fmt.Printf("input %+v\n", a)
	if len(a) == 0 {
		return a
	}
	if min >= max {
		return []int{}
	}
	l := max - min + 1
	// counts
	cnt := make([]int, l)
	for _, x := range a {
		cnt[x]++
	}
	fmt.Printf("cnt %+v\n", cnt)
	// index
	idx := make([]int, l)
	idx[0] = cnt[0]
	for i := 1; i < l; i++ {
		idx[i] = idx[i-1] + cnt[i]
	}
	fmt.Printf("idx %+v\n", idx)
	// output
	res = make([]int, len(a))
	j := 0
	for i := 0; i < l; i++ {
		n := a[j]
		pos := idx[n]
		for {
			pos--
			if pos <= 0 {
				break
			}
			res[pos] = n
		}
		j++
	}
	fmt.Printf("result %+v\n", res)
	return res
}

func main() {
	// min 0 max 3
	a := []int{1, 0, 3, 1, 3, 1}
	s := countSort(a, 0, 3)
	_ = s

}
