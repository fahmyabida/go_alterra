package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(EncCaesarChiper(3, "a"))
}

func EncCaesarChiper(offset int, s string) string {
	result := ""
	alphabetic := "abcdefghijklmnopqrstuvwxyz"
	maxAlphabetic := len(alphabetic)
	for _, row := range s {
		idxOfAlphabet := strings.Index(alphabetic, string(row))
		idxChiperCaesar := (idxOfAlphabet + offset) % maxAlphabetic
		result += string(alphabetic[idxChiperCaesar])
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
