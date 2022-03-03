package main

import "fmt"

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

func main() {
	fmt.Println("hasil : ", PrimeNumber(5))
}
