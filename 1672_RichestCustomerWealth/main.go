package main

import "fmt"

func main() {
	accounts := [][]int{{1, 2, 3}, {3, 2, 1}}
	maxWealth := maximumWealth(accounts)
	fmt.Println(maxWealth)
}

func maximumWealth(accounts [][]int) int {
	m := len(accounts)
	n := len(accounts[0])
	if m >= 1 && n <= 50 {
		maxWealth := 0
		for i := range accounts {
			hisWealth := 0
			for _, val := range accounts[i] {
				if 1 <= val && val <= 100 {
					hisWealth = hisWealth + val
				} else {
					panic("Wealth should be between 1 and 100")
				}
			}
			if hisWealth >= maxWealth {
				maxWealth = hisWealth
			}

		}
		return maxWealth
	} else {
		panic(" At least one user should be there and maximum number of accounts per user are 50")
	}
}
