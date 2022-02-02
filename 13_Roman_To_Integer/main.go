package main

import "fmt"

func main() {
	num := romanToInt("MCMXCIV")
	fmt.Println(num)
}
func romanToInt(s string) int {
	num := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'I':
			if i+1 >= len(s) || (s[i+1] != 'V' && s[i+1] != 'X') {
				num = num + 1
			} else if s[i+1] == 'V' {
				num = num + 4
				i++
			} else if s[i+1] == 'X' {
				num = num + 9
				i++
			}
		case 'V':
			num = num + 5
		case 'X':
			if i+1 >= len(s) || (s[i+1] != 'L' && s[i+1] != 'C') {
				num = num + 10
			} else if s[i+1] == 'L' {
				num = num + 40
				i++
			} else if s[i+1] == 'C' {
				num = num + 90
				i++
			}
		case 'L':
			num = num + 50
		case 'C':
			if i+1 >= len(s) || (s[i+1] != 'D' && s[i+1] != 'M') {
				num = num + 100
			} else if s[i+1] == 'D' {
				num = num + 400
				i++
			} else if s[i+1] == 'M' {
				num = num + 900
				i++
			}
		case 'D':
			num = num + 500
		case 'M':
			num = num + 1000
		}
	}

	return num
}
