package main

import (
	"fmt"
)

// 2D identity matrix
func I(order int) [][]float64 {
	matrix := make([][]float64, order)
	for i := 0; i < order; i++ {
		matrix[i] = make([]float64, order)
		matrix[i][i] = 1.0
	}
	return matrix
}

// main method
func main() {
	fmt.Println(I(3))
}
