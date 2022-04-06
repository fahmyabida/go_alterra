package main

import (
	"fmt"
	"math"
)

func PrimeNumber(N int) (isPrimeNumber bool) {
	if N <= 1 {
		return false // bukan bilangan prima
	}
	for i := 2; i < N; i++ {
		sisaBagi := N % i
		fmt.Printf("sisaBagi dari %v dibagi %v adalah %v\n", N, i, sisaBagi)
		if sisaBagi == 0 {
			return false // bukan bilangan prima
		}
	}
	return true // bilangin prima
}

func PrimeNumber2(number int) bool {
	if number <= 2 {
		return false
	}
	akar := math.Sqrt(float64(number))
	root := int(akar)
	for i := 2; i <= root; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("hasil : ", PrimeNumber2(127))
}
