package main

import "fmt"

type bidangDatar interface {
	area() float64
}

type kotak struct {
	panjang, tinggi float64
}

func (k kotak) area() float64 {
	return k.panjang * k.tinggi
}

func hitungBidang(bd bidangDatar) {
	fmt.Println("hasil bd area", bd.area())
}

func main() {
	ktk := kotak{10, 5}
	hitungBidang(ktk)
}
