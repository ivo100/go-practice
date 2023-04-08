package main

import "log"

// https://practice.geeksforgeeks.org/problems/max-rectangle/1

/*
Given a binary matrix, Your task is to complete the function maxArea
which the maximum size rectangle area in a binary-sub-matrix with all 1’s.
The function takes 3 arguments the first argument is the Matrix M[ ] [ ]
and the next two are two  integers n and m which denotes the size of the matrix M.
Your function should return an integer denoting the area of the maximum rectangle .

Input:
The first line of input is an integer T denoting the no of test cases . Then T test cases follow. The first line of each test case are 2 integers n and m denoting the size of the matrix M . Then in the next line are n*m space separated values of the matrix M.

Output:
For each test case output will be the area of the maximum rectangle .

Constraints:
1<=T<=50
1<=n,m<=50
0<=M[][]<=1

Example:
Input
1
4 4
0 1 1 0 1 1 1 1 1 1 1 1 1 1 0 0

Output
8

Explanation
For the above test case the matrix will look like
0 1 1 0
1 1 1 1
1 1 1 1
1 1 0 0
the max size rectangle is
1 1 1 1
1 1 1 1
and area is 4*2 = 8
*/
func main() {
	a := [][]int{
		{0, 1, 1, 0},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 0, 0},
	}
	s := NewSolver(a)
	ma := s.findMaxArea()
	log.Printf("max area: %v", ma)
}

type Solver struct {
	a       [][]int
	nr      int
	nc      int
	r       int
	c       int
	maxArea int
}

func NewSolver(a [][]int) *Solver {
	s := &Solver{}
	s.a = a
	s.nr = len(a)
	s.nc = len(a[0])
	log.Printf("NewSolver nr %v, nc %v, a %+v", s.nr, s.nc, a)
	return s
}

func (s *Solver) findMaxArea() int {
	for {
		s.r, s.c = s.findCell()
		if s.r < 0 || s.c < 0 {
			return s.maxArea
		}
		w, h := s.findRect()
		ar := w * h
		if ar > s.maxArea {
			s.maxArea = ar
		}
		// next - to the right of top left corner of the found rect
		s.c += w
		if s.c >= s.nc {
			s.c = 0
			s.r += h
			if s.r >= s.nr {
				break
			}
		}

	}
	return s.maxArea
}

func (s *Solver) findCell() (r, c int) {
	for r = s.r; r < s.nr; r++ {
		for c = s.c; c < s.nc; c++ {
			if s.a[r][c] == 1 {
				//log.Printf("findCell returns r %v, c %v", r, c)
				return r, c
			}
		}
	}
	return -1, -1
}

// given cell r,c with val 1 - find max rect by expanding to the right and down
func (s *Solver) findRect() (w, h int) {
	// sanity check
	if s.a[s.r][s.c] != 1 {
		return 0, 0
	}

	w = 0
	h = 1
	// current row
	r := s.r
	c := s.c
	for ; c < s.nc; c++ {
		// find range of 1111
		if s.a[r][c] != 1 {
			break
		}
		w++
	}
	// s.r has 1s from s.c to c
	// now try cells of next row - under 1 cells
	if r == s.nr-1 {
		// last row
		return w, h
	}
	r++
	//i := s.c
	//j := c - 1
	for c1 := s.c; c1 < c; c++ {
		// find range of 1111
		if s.a[r][c] != 1 {
			break
		}
		w++
	}

	// s.a[r][c] == 1
	// mark visited
	//s.a[r][c] = 2
	log.Printf("findRect returns w %v, h %v", w, h)
	return w, h
}
