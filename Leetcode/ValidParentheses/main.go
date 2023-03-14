package main

import "fmt"

type stack []string

func (s *stack) Push(b string) {
	*s = append(*s, b)
}

func (s *stack) Pop() string {
	if s.Len() != 0 {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element
	}

	return ""
}

func (s stack) Len() int {
	return len(s)
}

func isValid(s string) bool {
	brackets := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}
	st := new(stack)
	for _, ch := range s {
		b := string(ch)
		if isClosedBracket(b) {
			if brackets[b] != st.Pop() {
				return false
			}
		} else {
			st.Push(b)
		}
	}

	return st.Len() == 0
}

func isClosedBracket(b string) bool {
	return b == ")" || b == "]" || b == "}"
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([{}])"))
	fmt.Println(isValid("([{}]}"))
	fmt.Println(isValid("]"))
}
