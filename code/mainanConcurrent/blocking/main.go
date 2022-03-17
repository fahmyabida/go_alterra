package main

import (
	"fmt"
	"time"
)

// masak mie
func main() {
	// fmt.Println("---start---")
	// fmt.Println("buka bungkus")
	// fmt.Println("=======")

	rebusAirSudahSelesai := make(chan string)
	go MasukanMieKeRebusanAir(rebusAirSudahSelesai) // kirim background :: 1 2 - - 3 4
	go RebusAir(rebusAirSudahSelesai)               // kirim background ::   A B C D
	time.Sleep(1 * time.Nanosecond)

	// fmt.Println("=======")
	// fmt.Println("buka & tuang bumbu ke mangkok")
	// fmt.Println("campurkan mie dgn bumbu")
	// fmt.Println("---end---")
}

func RebusAir(rebusAirSudahSelesai chan string) {
	fmt.Println("ambil air secukupnya")      // A
	fmt.Println("masukan ke panci")          // B
	fmt.Println("rebus air")                 // C
	rebusAirSudahSelesai <- "sudah mendidih" // D
}

func MasukanMieKeRebusanAir(rebusAirSudahSelesai chan string) {
	fmt.Println("// belum boleh lanjut")           // 1
	s := <-rebusAirSudahSelesai                    // 2
	fmt.Println("// sudah boleh lanjut")           // 3
	fmt.Println(s)                                 // -
	fmt.Println("masukan mie kedalam air rebusan") // 4
}
