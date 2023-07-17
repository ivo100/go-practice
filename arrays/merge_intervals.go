package main

import "log"

/*
Merge overlapping intervals

Given a list of intervals, merge all the overlapping intervals to produce a list
that has only mutually exclusive intervals.
*/

type Range struct {
	Start int
	End   int
}

type IntervalCalculator interface {
	Before(Range, Range) bool
	After(Range, Range) bool
	Contains(Range, Range) bool
	Overlaps(Range, Range) bool
	Result() []Range
	Merge(Range, Range)
	Add(Range)
}

type intervalCalculator struct {
	input  []Range
	merged []Range
}

func (c *intervalCalculator) Result() []Range {
	return c.merged
}

func (c *intervalCalculator) Before(r Range, r2 Range) bool {
	return r.End < r2.Start
}

func (c *intervalCalculator) After(r Range, r2 Range) bool {
	return r2.Start > r.End
}

func (c *intervalCalculator) Contains(r Range, r2 Range) bool {
	return (r2.Start >= r.Start && r2.End <= r.End) || (r.Start >= r2.Start && r.End <= r2.End)
}

func (c *intervalCalculator) Overlaps(r Range, r2 Range) bool {
	if r2.Start >= r.Start && r2.Start <= r.End {
		return true
	}
	if r2.End >= r.Start && r2.End <= r.End {
		return true
	}
	return false
}

func (c *intervalCalculator) Add(r Range) {
	c.merged = append(c.merged, r)
}

func (c *intervalCalculator) Merge(r Range, r2 Range) {
	s := min(r.Start, r2.Start)
	e := max(r.End, r2.End)
	//log.Printf("merged start %v, end %v", s, e)
	// we have to update r
	for i, a := range c.merged {
		if a.Start == r2.Start && a.End == r2.End {
			c.merged[i].Start = s
			c.merged[i].End = e
			log.Printf("merged %v", c.merged)
			return
		}
	}
}

func NewIntervalCalculator(intervals []Range) IntervalCalculator {
	return &intervalCalculator{input: intervals, merged: make([]Range, 0)}
}

func Merge(a [][]int) [][]int {
	ret := make([][]int, 0)
	rr := make([]Range, 0)
	for _, b := range a {
		r := Range{
			Start: b[0],
			End:   b[1],
		}
		rr = append(rr, r)
	}

	c := NewIntervalCalculator(rr)

	for _, in := range rr {
		toAdd := true
		for _, res := range c.Result() {
			log.Printf("compare %v with %v", in, res)
			if c.Contains(in, res) {
				log.Printf("contains true")
				toAdd = false
				break
			}
			if c.Overlaps(in, res) {
				log.Printf("%v overlaps %v", in, res)
				c.Merge(in, res)
				toAdd = false
				break
			}
		}
		if toAdd {
			c.Add(in)
		}
	}
	log.Printf("result %v", c.Result())

	for _, r := range c.Result() {
		log.Printf("b %v", r)
		ret = append(ret, []int{r.Start, r.End})
	}
	return ret
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
