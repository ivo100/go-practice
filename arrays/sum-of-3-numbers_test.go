package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_find3(t *testing.T) {
	aa := [][]int{
		{1, 2, 3, 4},
		{-1, 1, 2, 3, 4},
		{-2, -1, 3, 2, 1},
		{-2, -1, 2, 2, 2},
		{-2, -1, 2, 2, 2},
		{4, 0, 2},
		{6, 0, 0},
		{8, -1, -1},
	}
	for _, a := range aa {
		assert.True(t, Find3(a, 6))
	}
}
