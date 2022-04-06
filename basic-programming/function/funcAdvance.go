package main

import "fmt"

func Print() {
	defer func() {
		fmt.Println("1")
	}()

	world := func() {
		fmt.Println("2")
	}
	world()

	fmt.Println("3")
	fmt.Println("4")
}
