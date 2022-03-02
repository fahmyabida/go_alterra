package main

import "fmt"

func main() {
	hmap := make(map[string]int)
	hmap["kayuza"] = 1
	hmap["iqram"] = 1234
	val, isFound := hmap["feng"]
	fmt.Println(val, "|", isFound)

	// arrSatu := []string{"abc", "def", "ghi", "qrs"}
	// arrDua := []string{"jkl", "mno"}
	// fmt.Println(ArrayMerge(arrSatu, arrDua))
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
