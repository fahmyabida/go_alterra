package main

import (
	"fmt"
	"sync"
)

func main() {
	text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua"

	wg.Add(len(text))
	for i := 0; i < len(text); i++ {
		go hitungFrekuensiHuruf(string(text[i]))
	}
	wg.Wait()
	fmt.Println(totalHuruf)

}

var totalHuruf map[string]int = map[string]int{}
var wg = sync.WaitGroup{}
var mx = sync.Mutex{}

func hitungFrekuensiHuruf(s string) {
	/*
		"o" = 2
		process bg A --> "o" = 2 + 1 = 3
		process bg B --> "o" = 3 + 1 = 4
		"o" = 4
	*/
	mx.Lock()
	value, isFound := totalHuruf[s]
	if !isFound {
		totalHuruf[s] = 1
	}
	totalHuruf[s] = value + 1
	mx.Unlock()
	wg.Done()
}
