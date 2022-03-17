package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseComplmentOfDNAWithReplace("ACGT"))
}

func reverseComplmentOfDNAWithReplace(s string) (result string) {
	// 1. reverse
	// ACCGGGTTTT
	sReverse := reverseString(s)
	// A --> T(W) | C --> G(Y)
	// A->T | T->A 
	// C->G | G->C
	sReverse = strings.Replace(sReverse, "A", "#", -1) 
	sReverse = strings.Replace(sReverse, "T", "A", -1)
	sReverse = strings.Replace(sReverse, "C", "$", -1) 
	sReverse = strings.Replace(sReverse, "G", "C", -1)

	sReverse = strings.Replace(sReverse, "#", "T", -1)
	sReverse = strings.Replace(sReverse, "$", "G", -1)
	return sReverse
}

func reverseComplmentOfDNA(s string) (result string) {
	sReverse := reverseString(s)
	for i := 0; i < len(sReverse); i++ {
		x := string(sReverse[i])

		if x == "A" {
			x = "T"
		} else if x == "T" {
			x = "A"
		} else if x == "C" {
			x = "G"
		} else if x == "G" {
			x = "C"
		}
		result += x
	}
	return result
} 

func reverseString(s string) string {
	hasilReverse := ""
	for i := len(s) - 1; i >= 0; i-- {
		hasilReverse += string(s[i])
	}
	return hasilReverse
}
