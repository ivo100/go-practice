package sherlock

import (
	"fmt"
	"sort"
	"strings"
)

// https://www.hackerrank.com/challenges/sherlock-and-anagrams
/*
cheating

from collections import Counter
def sherlockAndAnagrams(s):
    count = 0
    print("INPUT", s)
    for l in range(1,len(s)+1):
        print("len", l, "j in 0:", len(s)-l+1)
        a = ["".join(sorted(s[j:j+l])) for j in range(len(s)-l+1)]
        print(a)
        hash = Counter(a)
        for k in hash:
            c = hash[k] * (hash[k]-1)/2
            count += c
            print("k", k, "c", c, "count", count)

    return int(count)


INPUT abba

len 1 j in 0: 4
['a', 'b', 'b', 'a']
k a c 1.0 count 1.0
k b c 1.0 count 2.0

len 2 j in 0: 3
['ab', 'bb', 'ab']
k ab c 1.0 count 3.0
k bb c 0.0 count 3.0

len 3 j in 0: 2
['abb', 'abb']
k abb c 1.0 count 4.0

len 4 j in 0: 1
['aabb']
k aabb c 0.0 count 4.0

*/
func substrWithLen(s string, l int) []string {
	res := make([]string, 0)
	L := len(s)
	for i := 0; i <= L-l; i++ {
		res = append(res, s[i:i+l])
	}
	return res
}

//func isAnagram(a, b string) bool {
//	aa := strings.Split(a, "")
//	sort.Strings(aa)
//	a = strings.Join(aa, "")
//	bb := strings.Split(b, "")
//	sort.Strings(bb)
//	b = strings.Join(bb, "")
//	return a == b
//}

func findAnagrams(list []string) int {
	m := make(map[string]int)
	for _, s := range list {
		a := strings.Split(s, "")
		sort.Strings(a)
		k := strings.Join(a, "")
		m[k]++
	}
	c := 0
	for _, v := range m {
		c += v * (v - 1) / 2
	}
	return c
}

// Complete the sherlockAndAnagrams function below.
func SherlockAndAnagrams(s string) int32 {
	L := len(s)
	if L < 2 {
		return 0
	}
	tot := 0
	for l := 1; l < L; l++ {
		lst := substrWithLen(s, l)
		fmt.Printf("s %v, l %d, list %+v\n", s, l, lst)
		n := findAnagrams(lst)
		fmt.Printf("n %v\n", n)
		tot += n
	}
	return int32(tot)
}
