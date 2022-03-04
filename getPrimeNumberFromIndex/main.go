package main

import "fmt"

func main() {
	fmt.Println(primeX(3))
}

func primeX(n int) int {
	result, i := 0, 2
	for {
		if val, ok := isPrime(i); ok {
			result++
			if result == n {
				return val
			}
		}
		i ++
    }
}

func isPrime(n int) (int, bool) {
	if n <= 1 {return 0, false}
	for i := 2; i < n; i++ {
		if n % i == 0 {
			return 0, false
		}
	}
	return n, true
}