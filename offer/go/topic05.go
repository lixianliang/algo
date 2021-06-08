package main

import (
	"fmt"
)

func replaceSpace(s string) string {
	if len(s) == 0 {
		return ""
	}

	newStr := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			//newStr = append(newStr, []byte("%20")...)
			newStr = append(newStr, '%', '2', '0')
		} else {
			newStr = append(newStr, s[i])
		}
	}

	return string(newStr)
}

func main() {
	s := "We are happy."
	fmt.Println(replaceSpace(s))
}
