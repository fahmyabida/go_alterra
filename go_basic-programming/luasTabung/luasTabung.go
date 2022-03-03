package main

import "fmt"

const phi = 3.41

func main() {
	// L = (2 * phi * (r * r)) + (2 * phi * r * t)
	var T float64
	var r float64
	fmt.Println("## Menghitung luas tabung ##")
	fmt.Println("input tinggi tabung")
	fmt.Scanf("%v", &T)

	fmt.Println("input ruas tabung")
	fmt.Scanf("%v", &r)

	L := 2 * phi * r * (r + T)
	fmt.Println("hasilnya adalah ", L)
	// // L = 2 * phi * r * (r + t)
}






