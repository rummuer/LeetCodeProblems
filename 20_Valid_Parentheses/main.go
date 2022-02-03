package main

import "fmt"

func main() {
	result := isValid("()")
	fmt.Println(result)
}

func isValid(s string) bool {
	a := make([]string, len(s))
	top := -1
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(', '{', '[':
			top++
			a[top] = string(s[i])
		case ')':
			if top != -1 && a[top] == string('(') {
				top--
			} else {
				return false
			}
		case '}':
			if top != -1 && a[top] == string('{') {
				top--
			} else {
				return false
			}
		case ']':
			if top != -1 && a[top] == string('[') {
				top--
			} else {
				return false
			}
		}
	}
	if top != -1 {
		return false
	}
	return true
}
