package main

import "fmt"

func main() {
	p := isPalindrome(121)
	fmt.Println(p)
}

func isPalindrome(x int) bool {
	y := x
	sum := 0
	for y > 0 {
		sum = sum*10 + y%x
		y = y / 10
	}
	if x == sum {
		return true
	} else {
		return false
	}
}
