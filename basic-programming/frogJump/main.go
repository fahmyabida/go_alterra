package main

import (
	"fmt"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Frog(jumps []int) int {
	n := len(jumps)
	dp := make([]int, n+1)

	dp[1] = Abs(jumps[1] - jumps[0])

	for i := 2; i < n; i++ {
		jump_1 := Abs(jumps[i]-jumps[i-1]) + dp[i-1]
		jump_2 := Abs(jumps[i]-jumps[i-2]) + dp[i-2]
		dp[i] = Min(jump_1, jump_2)
	}

	return dp[n-1]
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))         // 30
	fmt.Println(Frog([]int{10, 10}))                 // 0
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40
}
