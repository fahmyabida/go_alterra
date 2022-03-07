package main

import (
	"fmt"
)

type pair struct {
	name  string
	count int
}

func main() {
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))
}

func MostAppearItem(items []string) []pair {
	hMap := make(map[string]int)
	for i := 0; i < len(items); i++ {
		if val, ok := hMap[items[i]]; ok {
			hMap[items[i]] = val + 1
		} else {
			hMap[items[i]] = 1
		}
	}
	list := []pair{}
	for key, value := range hMap {
		pairObj := pair{
			name: key, 
			count: value,
		}
		list = append(list, pairObj)
	}
	return list
}
