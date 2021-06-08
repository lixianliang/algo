package main

import (
	"fmt"
)

/*func isMatch(s string, p string) bool {
	if len(s) == 0 || len(p) == 0 {
		return false
	}

	return matchRec(s, p)
}

func matchRec(s string, p string) bool {
	if len(s) == 0 && len(p) == 0 {
		return true
	}
	if len(s) != 0 && len(p) == 0 {
		return false
	}

	if len(p) > 2 && p[1] == byte('*') { // p的长度大于2
		if s[0] == p[0] || (p[0] == byte('.') && len(s) != 0) {
			return matchRec(s, p[2:]) || matchRec(s[1:], p) || matchRec(s[1:], p[2:])
		} else {
			return matchRec(s[1:], p[2:])
		}
	} else if s[0] == p[0] || (p[0] == byte('.') && len(s) != 0) {
		return matchRec(s[1:], p[1:])
	}

	return false
}*/

func main() {
	fmt.Println(isMatch("aab", "c*a*b"))
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "a*"))
}
