package main

import "fmt"

func main() {
	var a string
	a = "Fahmy"
	fmt.Println("address a", &a)
	fmt.Println("value a", a)

	fmt.Println("=====")
	
	var b *string
	b = &a
	fmt.Println("address b", &b)
	fmt.Println("value raw b", b)
	fmt.Println("value pointer b", *b)
	fmt.Println("=====")
	var c *string
	c = b
	fmt.Println("address c", &c)
	fmt.Println("value raw c", c)
	// fmt.Println("value pointer c", *c)
}
