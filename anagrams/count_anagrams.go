// package anagrams
package main

import (
	"log"
	"strings"
)

/*
Given a word S and a text C.
Return the count of the occurrences of anagrams of the word in the text.

a=for
s=forxxorfxdofr
count=3 for, orf, ofr
*/
func count(aa string, ss string) int {
	a := strings.Split(aa, "")
	s := strings.Split(ss, "")
	//log.Printf("a %v, s %v", a, s)
	la := len(a)
	ls := len(s)
	if la >= ls {
		return 0
	}
	m := make(map[string]int)
	for _, c := range a {
		m[c]++
	}
	//log.Printf("map %+v", m)
	cnt := 0

loop1:
	for i := 0; i < ls-la+1; i++ {
		mm := make(map[string]int)
		for _, r := range s[i : i+la] {
			_, ok := m[r]
			if !ok {
				continue loop1
			}
			mm[r]++
		}
		for k, v := range m {
			if mm[k] != v {
				continue loop1
			}
		}
		cnt++
		//log.Printf("i %v, cnt %v, subs %+v", i, cnt, s[i:i+la])
	}

	return cnt
}

func main() {
	c := 0
	a := "for"
	s := "forxxorfxdofr"
	c = count(a, s)
	log.Printf("count %v, OK: %v", c, c == 3)

	s = "aabaabaa"
	a = "aaba"
	c = count(a, s)
	log.Printf("count %v, OK: %v", c, c == 4)
}
