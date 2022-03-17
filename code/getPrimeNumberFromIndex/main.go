package main

import "fmt"

func main() {
	fmt.Println(primeX(3))
}

func primeX(idx int) int {
	result, number := 1, 5
	for {
		if val, ok := isPrime(number); ok {
			result = result + 1
			if result == idx {
				return val
			}
		}
		number ++
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