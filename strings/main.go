package main

import (
	"fmt"
	"strings"
)

func permute(str string) []string {
	if len(str) == 1 {
		return []string{str}
	}

	var permutations []string

	for i := 0; i < len(str); i++ {
		firstChar := str[i]
		remainingChars := str[:i] + str[i+1:]

		for _, permutation := range permute(remainingChars) {
			permutations = append(permutations, string(firstChar)+permutation)
		}
	}

	return permutations
}

func combinations(n, k int, input []string) [][]string {
	if k > n {
		return [][]string{}
	}
	if k == n {
		return [][]string{input}
	}
	if k == 0 {
		return [][]string{{}}
	}

	var result [][]string
	var first = input[0]
	var rest = input[1:]

	for _, subset := range combinations(n-1, k-1, rest) {
		subset = append(subset, first)
		result = append(result, subset)
	}

	for _, subset := range combinations(n-1, k, rest) {
		result = append(result, subset)
	}

	return result
}

// Given a string, remove characters until the string is made up of any two alternating characters.
// When you choose a character to remove, all instances of that character must be removed. Determine the longest string possible that contains just two alternating letters.
func alternate(s string) int {
	fmt.Printf("input: %s\n", s)
	maxL := 0
	// load string s to map
	// create a map to store the characters
	var let = make(map[byte]bool)
	n := len(s)
	for i := 0; i < n; i++ {
		let[s[i]] = true
	}

	var sb strings.Builder
	for b := range let {
		fmt.Printf("%s\n", string(b))
		sb.WriteString(string(b))
	}

	k := len(let)
	fmt.Printf("let: %v\n", let)
	a := make([]string, 0)
	for b := range let {
		a = append(a, string(b))
	}
	comb := combinations(k, k-2, a)
	fmt.Printf("combinations: %+v\n", comb)
	for _, p := range comb {
		fmt.Printf("combination: %+v\n", p)
		ss := s
		for _, q := range p {
			fmt.Printf("remove %s\n", q)
			ss = strings.Replace(ss, q, "", -1)
			if len(ss) < 3 {
				break
			}
			if invalid(ss) {
				fmt.Printf("invalid %s\n", ss)
				continue
			}
			fmt.Printf("%s\n", ss)
			if alternating(ss) {
				fmt.Printf("FOUND alternating %s\n", ss)
				l := len(ss)
				if l > maxL {
					maxL = l
				}
				break
			}
			// not alternating - try removing more
		}

	}

	//perm := permute(sb.String())
	//fmt.Printf("permutations %+v\n", perm)
	//for _, p := range perm {
	//	l := try(p)
	//}
	return maxL
}

// returns true (invalid) if 2 consecutive chars are the same
func invalid(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return true
		}
	}
	return false
}

//	babab. This is a valid  as there are only two distinct characters (a and b),
//
// and they are alternating within the string.
func alternating(s string) bool {
	if len(s) < 2 {
		return false
	}
	a := s[0]
	b := s[1]
	l := len(s) - 2
	for i := 2; l > 0; {
		if s[i] != a {
			return false
		}
		l--
		i++
		if l == 0 {
			break
		}
		if s[i] != b {
			return false
		}
		l--
		i++
	}
	return true
}

func digits(s string) string {
	r := []rune(s)
	var res []rune
	for _, c := range r {
		if c >= '0' && c <= '9' {
			res = append(res, c)
		}
	}
	return (string)(res)
}

func isValidUSPhone(input string) bool {
	const USPhoneSize = 10
	number := digits(input)
	l := len(number)
	if l < USPhoneSize {
		return false
	}
	switch l {
	case USPhoneSize:
		// not guaranteed to be valid - just syntax check
		return true
	case 1 + USPhoneSize:
		return strings.HasPrefix(number, "1")
	default:
		// allow > 11 digits if ext present?
		// who uses extensions anymore?
		return strings.Contains(strings.ToLower(input), "ext")
	}
}

func main() {
	//s := "abaacdabd"
	//s := "beabeefeab"
	check := []string{"+1 (800) 123 4567", "6197575696", "619 757 5696", "619757569"}
	for _, phone := range check {
		fmt.Printf("digits(%q): %q, isValidUSPhone: %v\n",
			phone, digits(phone), isValidUSPhone(phone))
	}
	//s := "asdcbsdcagfsdbgdfanfghbsfdab"
	//l := alternate(s)
	//fmt.Printf("len %d\n", l)
}

/*
cheat


int validSize(string s, char first, char second){
    string ans = "";
    for(int i = 0; i < s.size(); i++){
        if(s[i] == first || s[i] == second){
            if(ans.size() > 0 && ans[ans.size() - 1] == s[i]) return 0;
            else ans+=s[i];
        }
    }
    if(ans.size() < 2) return 0;
    return ans.size();
}

int alternate(string s) {
    int ans = 0;
    for(char i = 'a'; i < 'z'; i++){
        for(char j = i + 1; j <= 'z'; j++){
           int r = validSize(s, i, j);
           if(r > ans) ans = r;
        }
    }
    return ans;
}

*/
