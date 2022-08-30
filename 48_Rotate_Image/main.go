package main

import "fmt"

func main() {
	matrix := make([][]int, 3)
	matrix[0] = []int{1, 2, 3}
	matrix[1] = []int{4, 5, 6}
	matrix[2] = []int{7, 8, 9}

	rotate(matrix)
	fmt.Println(matrix)
}

func rotate(matrix [][]int) {
	for i := 0; i < len(matrix[0]); i++ {
		for j := i + 1; j < len(matrix[0]); j++ {
			if i == j {
				continue
			} else {
				fmt.Println(i, j, len(matrix[0]))
				swap(i, j, matrix)
			}
		}
	}
	for i := 0; i < len(matrix[0]); i++ {
		for j := 0; j < len(matrix[0])/2; j++ {
			reflect(i, j, len(matrix[0]), matrix)
		}
	}

}
func swap(i int, j int, matrix [][]int) {

	temp := matrix[i][j]
	matrix[i][j] = matrix[j][i]
	matrix[j][i] = temp

}
func reflect(i int, j int, n int, matrix [][]int) {
	temp := matrix[i][j]
	matrix[i][j] = matrix[i][n-j-1]
	matrix[i][n-j-1] = temp

}
