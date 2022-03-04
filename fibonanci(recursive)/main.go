package main

import "fmt"

func main() {
	fmt.Println(fibonanci(3))
}

func fibonanci(index int) int {
	before := 0
	current := 1
	return before + current
}
