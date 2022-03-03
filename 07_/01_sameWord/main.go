package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Compare("KANGGORO", "KANG"))
	fmt.Println(Compare("KI", "KIJANG"))
	fmt.Println(Compare("KUPU-KUPU", "KUPU"))
}

func Compare(a, b string) string {
	shortest, longest := a, b
	if len(a) > len(b) {
		longest = a
		shortest = b
	}
	if strings.Contains(longest, shortest) {
		return shortest
	}
	return ""
}
