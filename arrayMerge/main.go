package main

import "fmt"

func main() {
	arrSatu := []string{"abc", "def", "ghi", "qrs"}
	arrDua := []string{"jkl", "mno"}
	fmt.Println(ArrayMerge(arrSatu, arrDua))
}

func ArrayMerge(arrSatu, arrDua []string) (hasilGabung []string) {
	hmap := make(map[string]bool)
	for _, valueSatu := range arrSatu {
		hmap[valueSatu] = true // kazuya = true dst
		hasilGabung = append(hasilGabung, valueSatu)
	}
	for _, valueDua := range arrDua {
		_, isFound := hmap[valueDua] // false
		if isFound == false {
			hmap[valueDua] = true
			hasilGabung = append(hasilGabung, valueDua)
		}
	}
	return hasilGabung
}
