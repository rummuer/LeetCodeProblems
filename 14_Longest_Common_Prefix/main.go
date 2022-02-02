package main

import "fmt"

func main() {
	prefix := longestCommonPrefix([]string{"cir", "car"})
	fmt.Println(prefix)
}
func longestCommonPrefix(strs []string) string {
	var prefix string
	for i := 0; i < len(strs[0]); i++ {
		count := 0
		for j := 1; j < len(strs); j++ {
			if (i < len(strs[j])) && (strs[0][i] == strs[j][i]) {
				count++
			}

		}
		if count == len(strs)-1 {
			prefix = prefix + string(strs[0][i])
		} else {
			break
		}
	}
	return prefix
}
