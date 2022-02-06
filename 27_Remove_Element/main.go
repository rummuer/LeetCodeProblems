package main

import "fmt"

func main() {
	r := removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2)
	fmt.Println(r)
}
func removeElement(nums []int, val int) int {
	times := 0
	for j := 0; j < len(nums)-1; j++ {
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] == val {
				temp := nums[i]
				nums[i] = nums[i+1]
				nums[i+1] = temp
				times++
			}
		}
	}
	k := 0
	for k = 0; k < len(nums); k++ {
		if nums[k] == val {
			break
		}
	}
	fmt.Println(nums, " ", k)
	return k
}
