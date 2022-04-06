package main

import (
	"fmt"
)

func main() {
	c := make(chan string)
	go Print(c)
	s := <-c
	fmt.Println(s)

	// go Print2()
	// time.Sleep(2 * time.Second)

}

func Print2() {
	fmt.Println("hallo")
}

func Print(c chan string) {
	s := "print something"
	c <- s
}
