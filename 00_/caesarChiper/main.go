package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(ChiperSubtition("zyx"))
	// s := "abc"
	// fmt.Println(string(s[0]))
}

func ChiperSubtition(s string) string {
	result := ""
	alphabetic := "abcdefghijklmnopqrstuvwxyz" // <-- 26 character
	kebalikan :=  "zyxwvutsrqponmlkjihgfedcba" // <-- 26 character
	// maxAlphabetic := len(alphabetic) // 26
	for _, row := range s {
		idxOfAlphabet := strings.Index(alphabetic, string(row))
		result += string(kebalikan[idxOfAlphabet])
	}
	return result
}



















func DecCaesarChiper(offset int, s string) string {
	result := ""
	alphabetic := "abcdefghijklmnopqrstuvwxyz"
	maxAlphabetic := len(alphabetic)
	for _, row := range s {
		idxOfAlphabet := strings.Index(alphabetic, string(row))
		idxChiperCaesar := (idxOfAlphabet - offset)
		if idxChiperCaesar < 0 {
			idxChiperCaesar = idxChiperCaesar * -1
			if idxChiperCaesar > maxAlphabetic {
				idxChiperCaesar = maxAlphabetic - idxChiperCaesar
				idxChiperCaesar = idxChiperCaesar * -1
				idxChiperCaesar = idxChiperCaesar % maxAlphabetic
			} else {
				idxChiperCaesar = maxAlphabetic - idxChiperCaesar
			}
		}
		// fmt.Println(idxChiperCaesar)
		// fmt.Println(string(alphabetic[idxChiperCaesar]))
		result += string(alphabetic[idxChiperCaesar])
	}
	return result
}
