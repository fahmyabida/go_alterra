package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))
	fmt.Println(Compare("KANGOORO", "KANG"))
	fmt.Println(Compare("KI", "KIJANG"))
	fmt.Println(Compare("KUPU-KUPU", "KUPU"))
	fmt.Println(Compare("ILALANG", "ILA"))
	a := 10
	b := 20
	Swap(&a, &b)
	fmt.Println(a, b)
}

func Compare(a, b string) string {
	arr := []string{}
	if len(a) > len(b) {
		a, b = b, a
	}
	for _, rowA := range a {
		for _, rowB := range b {
			if rowA == rowB {
				arr = append(arr, string(rowA))
				break
			}
		}
	}
	return strings.Join(arr, "")
}

func Swap(a, b *int) {
	*a, *b = *b, *a
}
