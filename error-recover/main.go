package main

import (
	"fmt"
	"strconv"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("code nya ttep lanjut walaupun ada error ini -->", err)
		}
	}()
	s := "A"
	number, err := strconv.Atoi(s) // <-- func untuk mengubah string menjadi integer
	if err != nil {
		panic(err)
	}
	fmt.Println(number)

}
