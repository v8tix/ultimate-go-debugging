package main

import "fmt"

func callMe(i int) int {
	if i == 0 {
		return 100
	}
	return callMe(i - 1)
}

func main() {
	fmt.Println(callMe(10))
}
