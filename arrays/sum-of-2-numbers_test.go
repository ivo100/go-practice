package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sum2(t *testing.T) {
	aa := [][]int{
		{1, 2, 3, 4},
		{-1, 1, 2, 3, 4},
		{-2, -1, 3, 2, 1},
		{0, -1, 0, 0, 4},
	}
	for _, a := range aa {
		assert.True(t, Find2Sum(a, 3))
	}
}
