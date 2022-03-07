package main

import "fmt"

var MONEY_LIST = []int{10000, 5000, 2000, 1000, 500, 200, 100, 50, 20, 10, 1}

func main() {
	fmt.Println("result", moneyChange(26123))
}

func moneyChange(amount int) (result []int) {
	fmt.Printf("uang awal %v\n", amount)
	for _, row := range MONEY_LIST {
		div := amount / row
		fmt.Printf("pecahan %v dapat sebanyak %v\n", row, div)
		if div == 0 {
			continue
		}
		for i := 0; i < div; i++ {
			result = append(result, row)
		}
		amount = amount - (row * div)
		fmt.Printf("sisa uang sekarang %v\n", amount)
		if amount == 0 {
			break
		}
	}
	return result
}
