package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println(isIsomorphic("egg", "add"))
	fmt.Println(isIsomorphic("foo", "bar"))
	fmt.Println(isIsomorphic("paper", "title"))
	fmt.Println(isIsomorphic("badc", "baba"))
}

func isIsomorphic(s string, t string) bool {
	if utf8.RuneCountInString(s) != utf8.RuneCountInString(t) {
		return false
	}

	mapping := make(map[string]string)
	for i, cht := range t {
		val, ok := mapping[string(cht)]
		if !ok {
			mapping[string(cht)] = string(s[i])
		} else if val != string(s[i]) {
			return false
		}
	}

	mapping = make(map[string]string)
	for i, chs := range s {
		val, ok := mapping[string(chs)]
		if !ok {
			mapping[string(chs)] = string(t[i])
		} else if val != string(t[i]) {
			return false
		}
	}

	return true
}
