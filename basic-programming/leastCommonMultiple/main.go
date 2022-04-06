package main

import (
	"fmt"
)

func main() {
	fmt.Println(lcm(2, 6))
}

// lcm : least common multiple :: penjumlahan paling terkecil dari a & b
func lcm(a, b int) int {
	return a * (b / gcd(a, b))
}

// gcd : Greatest Common Divisor :: pembagi terbesar dari a & b
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
