package main

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func Test_Merge(t *testing.T) {
	a := [][]int{
		{1, 2},
		{2, 3},
		{5, 10},
		{6, 7},
	}
	result := Merge(a)
	log.Printf("result %v", result)
	assert.Equal(t, 2, len(result))
	assert.Equal(t, 1, result[0][0])
	assert.Equal(t, 3, result[0][1])
}
