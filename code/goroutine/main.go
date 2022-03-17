package main

import (
	"fmt"
	"time"
)

func main() {
	go Print("Hello")
	Print("World")
	time.Sleep(1*time.Second)
}

func Print(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1*time.Millisecond)
		fmt.Println(s)
	}
}
