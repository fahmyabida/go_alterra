package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "fahmy"
	b := "FAHMY"
	fmt.Println(input == b)                  // case sensitive
	fmt.Println(strings.EqualFold(input, b)) // case in-sensitive

	s := "Hello World This is Amazing"
	arrS := strings.Split(s, " ")
	for idx, row := range arrS {
		fmt.Println("index", idx, "value", row)
	}
	s = strings.Join(arrS, ",")
	fmt.Println(s) // Hello,World,This,is,Amazing

	sNew := "Hello,World,This,is,Amazing,,,,,,,,,,,"
	sNew = strings.ReplaceAll(sNew, ",", " | ")
	fmt.Println(sNew)


	
}
