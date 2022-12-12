// Write a function 'canConstruct(target, wordBank)' that accepts a target string and array of strings
// The function should return a boolena indicating whether or not the 'target'
//can be construcyed by concatenating elements of the 'wordbank' array
// You may reuse elements of 'wordbank' as many times as needed.

//canConstruct(abcdef, [ab,abc,cd,def,abcd]) --> true
//canConstruct(skateboard,[bo,rd,ate,t,ska,sk,boar]) --> false
//canConstruct('',[dog,mouse,cat]) --> true

//Do not take strings in the middle. Always take prefixes only
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := [5]string{"ab", "abc", "cd", "def", "abcd"}
	fmt.Println(canConstruct("abcdef", s))
}

func canConstruct(target string, wordBank [5]string) bool {
	if target == "" {
		return true
	}
	for _, word := range wordBank {
		if strings.Index(target, word) == 0 {
			suffix := strings.TrimPrefix(target, word)

			if canConstruct(suffix, wordBank) == true {
				return true
			}
		}
	}
	return false
}
