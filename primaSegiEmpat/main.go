package main

import "fmt"

func main() {
	fmt.Println(PrimaSegiEmpat(2, 3, 13))
	fmt.Println("-------")
	fmt.Println(PrimaSegiEmpat(5, 2, 1))
}

func PrimaSegiEmpat(column, row, start int) int {
	total := 0
	for i := 0; i < row; i++ {
		j := 0
		for j < column {
			start = start + 1
			if val, ok := isPrime(start); ok {
				fmt.Printf("%v\t", val)
				total = total + val
				j++
			}
		}
		fmt.Printf("\n")
	}
	return total
}

func isPrime(start int) (int, bool) {
	if start <= 1 {
		return 0, false
	}
	for i := 2; i < start; i++ {
		if start%i == 0 {
			return 0, false
		}
	}
	return start, true
}
