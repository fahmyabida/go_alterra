package main

import "fmt"

func main() {
	a := 10
	b := 20
	fmt.Println("before", a, b)
	swap(&a, &b)
	fmt.Println("after", a, b)
}

func swap(a, b *int) {
	tmp := a
	*a = *b
	*b = *tmp
}
