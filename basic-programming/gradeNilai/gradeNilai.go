package main

import (
	"fmt"
)

func main() {
	var score int
	score = 90
	hasilKategori := MenentukanKategoriNilai(score)
	fmt.Println(hasilKategori)
}

func MenentukanKategoriNilai(score int) (hasil string) {
	fmt.Println("ini sebelum di kategorisasikan : ", hasil)
	if score <= 34 {
		hasil = "E"
	} else if score <= 49 {
		hasil = "D"
	} else if score <= 64 {
		hasil = "C"
	} else if score <= 79 {
		hasil = "B"
	} else if score <= 100 {
		hasil = "A"
	} 
	fmt.Println("ini setelah di kategorisasikan : ", hasil)
	return hasil
}
