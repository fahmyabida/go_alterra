package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	iArr := []int{5,3,6}
	fmt.Println(iArr)
	sort.Ints(iArr) // ASC
	fmt.Println(iArr)
	sort.Slice(iArr, func(i, j int) bool {
		return iArr[i] > iArr[j]
	}) // DESC
	fmt.Println(iArr)


	// fmt.Println(ChiperSubtitution("zyx"))
}

func ChiperSubtitution(s string) string {
	result := ""
	alphabetic := "abcdefghijklmnopqrstuvwxyz" // <-- 26 character
	kebalikan := "zyxwvutsrqponmlkjihgfedcba"  // <-- 26 character
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
				result += string(alphabetic[idxChiperCaesar])
			} else {
				idxChiperCaesar = maxAlphabetic - idxChiperCaesar
				result += string(alphabetic[idxChiperCaesar])
			}
		}
	}
	return result
}
