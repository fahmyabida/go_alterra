package main

import "fmt"

func hitungLuasSegitga(alas, tinggi float64) func() float64 {
	return func() float64 {
		return 0.5 * alas * tinggi
	}
}

func main() {
	luas := hitungLuasSegitga(4.0, 2.0)
	fmt.Println(luas())
}
