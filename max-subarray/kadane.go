package main

import (
	"fmt"
	"log"
)

// kadane algorithm
func maxSubarray(arr []int) int {
	log.Printf("maxSubarray %v", arr)
	maxEndingHere, maxSoFar := arr[0], arr[0]
	for i := 1; i < len(arr); i++ {
		n := arr[i]
		log.Printf("i %v, n %v", i, n)
		log.Printf("maxEndingHere %v, maxSoFar %v", maxEndingHere, maxSoFar)
		maxEndingHere = max(n, maxEndingHere+n)
		//maxSoFar = max(maxSoFar, maxEndingHere)
		if maxSoFar < maxEndingHere {
			maxSoFar = maxEndingHere
			log.Printf(">>> i %v, n %v", i, n)
			log.Printf(">>> maxEndingHere %v, maxSoFar %v", maxEndingHere, maxSoFar)
		}
		//log.Printf("maxEndingHere %v, maxSoFar %v", maxEndingHere, maxSoFar)
	}

	return maxSoFar
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	//p := []int{23, 13, 25, 29, 33, 19, 34, 45, 65, 67}
	//arr := []int{0, -10, 12, 4, 16, -14, 15, 11, 20}

	maxSum := maxSubarray(arr)
	fmt.Println("Maximum subarray sum:", maxSum)
}
