package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		hmap[string]int
		hmap itu jika ketemu maka int/value nya ditambahkan +1
					  tidak ketemu maka inisialisasi / int atau value = 1
		loop hmap dicari value == 1 nanti ditaruh di array
	*/
	s := "1234524"
	// 1 = 1
	// 2 = 2
	// 3 = 1
	// 4 = 2
	// 5 = 1
	fmt.Println(AngkaMunculSekali(s))
}

//map dengan key "...", value nya adalah ...
func AngkaMunculSekali(s string) []int {
	hmap := make(map[string]int)
	for i := 0; i < len(s); i++ {
		key := string(s[i])
		val, isFound := hmap[key]
		if isFound {
			hmap[key] = val + 1 // karena sudah ada sebelumnya, valuenya diupdate
		} else {
			hmap[key] = 1 //inisialisasi
		}
	}
	var sliceAngkaYgMunculSekali []int
	for key, value := range hmap {
		if value == 1 {
			intVarDariKey, _ := strconv.Atoi(key) // string to int
			sliceAngkaYgMunculSekali = append(sliceAngkaYgMunculSekali, intVarDariKey)
		}
	}
	return sliceAngkaYgMunculSekali
}
