package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go Print("hallo", &wg)
	go Print("hallo", &wg)
	wg.Wait()
	fmt.Println("next is --> exit")
}

func Print(s string, wg *sync.WaitGroup) {
	fmt.Println(s)
	wg.Done()
}
