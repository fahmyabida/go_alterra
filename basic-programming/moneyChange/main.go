package main

import (
	"fmt"
)

func moneyCoins(money int) []int {
	fmt.Println("\nInput :", money)
	fmt.Print("Output: ")
	var pecahan = []int{10000, 5000, 2000, 1000, 500, 200, 100, 50, 20, 10, 1}
	var hasil []int
	var banyak int
	for _, v := range pecahan {
		banyak = money / v
		for i := 0; i < banyak; i++ {
			hasil = append(hasil, v)
		}
		money %= v
	}
	return hasil
}

func main() {
	fmt.Println(moneyCoins(123))
	fmt.Println(moneyCoins(432))
	fmt.Println(moneyCoins(543))
	fmt.Println(moneyCoins(7752))
	fmt.Println(moneyCoins(15321))
}
