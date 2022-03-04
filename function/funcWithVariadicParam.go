package main

import "fmt"

func NamaFunc(a int) int {
	return a + 1
}

// PrintMessage 
func PrintMessage(a int, sArr ...string) {
	fmt.Println(a)
	for idx, row := range sArr {
		fmt.Println("index", idx, "value", row)
	}
}

func main() {
	a := "Hello"
	b := "World"
	c := "This"
	d := "Is"
	e := "Amazing"
	PrintMessage(10, a, b, c, d, e, a, b, c, e, d, c, a)
}
