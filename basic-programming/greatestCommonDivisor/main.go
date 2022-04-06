package main

import "fmt"

func main() {
	fmt.Println(gcd(10, 15))
}

// gcd : Greatest Common Divisor
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}


