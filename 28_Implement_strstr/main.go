package main

func main() {
	strStr("hello", "ll")
}
func strStr(haystack string, needle string) int {
	if len(needle) == 0 || (haystack == needle) {
		return 0
	}
	j := 0
	k := 0
	for i := 0; i < len(haystack); i++ {
		if j >= len(needle) {
			return k
		} else if haystack[i] == needle[j] {
			// fmt.Printf("Matches i=%v, j=%v, k=%v \n",i,j,k)
			j++
		} else {
			j = 0
			//fmt.Printf("Not Matches i=%v, j=%v, k=%v \n",i,j,k)
			i = k
			k++
		}
	}
	if j == len(needle) {
		return k
	} else {
		return -1
	}
}
