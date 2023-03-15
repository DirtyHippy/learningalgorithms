package main

import "fmt"

func main() {
	fmt.Println(isSubsequence("leeeeetcode", "yyylyyyyyyeyyyyeyyyyyeyyyyyyyyeyyyyyeyyyyytyyycyyyoyyyydyyyyyyeyyy"))
	fmt.Println(isSubsequence("abc", "ahbgdc"))
	fmt.Println(isSubsequence("axc", "ahbgdc"))
	fmt.Println(isSubsequence("abcb", "ahbgdcb"))
}

func isSubsequence(s string, t string) bool {
	sub := ""
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(t); j++ {
			if s[i] == t[j] {
				t = t[j+1:]
				sub += string(s[i])
				break
			}
		}
	}

	return len(s) == len(sub)
}
