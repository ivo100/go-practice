package main

import "log"

type solver struct {
	// goal
	k int
	// input array
	a     []int
	start int
	// min len of subseq
	min int
}

func NewSolver(k int, a []int) *solver {
	s := &solver{a: a, k: k, min: len(a) + 1}
	return s
}

// seq. of land plots with area a[i]
// given area k - find min num conseq. plots with area sum = k
func (s *solver) find(beg int) (idx, minLen int) {
	ps := make([]int, 1)
	ps[0] = s.a[beg]
	i := 0
	for j := beg + 1; j < len(s.a); j++ {
		sum := ps[j-beg-1] + s.a[j]
		ps = append(ps, sum)
		//log.Printf("i %v, j %v, sum %v, ps %+v",i,j,sum, ps)
		if sum < s.k {
			continue
		}
		if sum == s.k {
			// got it
			len := j - i - beg + 1
			if len < s.min {
				s.min = len
				s.start = beg + i
			}
			return beg + i, len
		}
		// s > k
		// incr i - left index until s <= k
		for i < j {
			x := s.a[i]
			sum -= x
			i++
			//log.Printf("== i %v, j %v, x %v, sum %v",i,j,x,sum)
			if sum < s.k {
				j++
				sum += s.a[j]
			}
			if sum == s.k {
				// got it
				len := j - i - beg + 1
				if len < s.min {
					s.min = len
					s.start = beg + i
				}
				return beg + i, len
			}
		}
	}
	//}
	return -1, -1
}

func main() {
	a := []int{1, 3, 2, 1, 4, 1, 3, 2, 1, 1, 2}
	k := 8
	s := NewSolver(k, a)

	idx := -1
	n := 0
	for {
		idx, n = s.find(idx + 1)
		log.Printf("idx %v, n %v, s.min %v, s.start %v", idx, n, s.min, s.start)
		if idx < 0 {
			break
		}
	}

	//idx, n = s.find(idx+1)
	//log.Printf("222 idx %v, n %v, s.min %v", idx, n, s.min)
	//idx, n = s.find(idx+1)
	//log.Printf("333 idx %v, n %v, s.min %v", idx, n, s.min)
	//idx, n = s.find(idx+1)
	//log.Printf("444 idx %v, n %v, s.min %v", idx, n, s.min)
	//idx, n = s.find(idx+1)
	//log.Printf("555 idx %v, n %v, s.min %v", idx, n, s.min)

}
