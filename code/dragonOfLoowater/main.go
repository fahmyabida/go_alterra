package main

import (
	"fmt"
	"sort"
)

func main() {
	// var dragonHead, knightHeight []int
	// dragonHead = []int{5, 4, 4}
	// knightHeight = []int{7, 8, 3, 4}
	// DragonOfLoowater(dragonHead, knightHeight)

	for i := 0; i < 10; i++ {
		if i == 2 {
			break
		}
		fmt.Println(i)
	}
}

func DragonOfLoowater(dragonHead, knightHeight []int) {
	if len(dragonHead) > len(knightHeight) {
		fmt.Println("knight fall")
	}
	sort.Ints(dragonHead)
	sort.Ints(knightHeight)
	totalMinimumHeightOfKnight := 0
	pointTemp := 0
	for i := 0; i < len(dragonHead); i++ {
		for j := pointTemp; j < len(knightHeight); j++ {
			if knightHeight[j] >= dragonHead[i] {
				totalMinimumHeightOfKnight += knightHeight[j] // knight can crop th dragon head!
				pointTemp = j + 1
				break
			}
		}
	}
	if totalMinimumHeightOfKnight == 0 {
		fmt.Println("knight fall!")
	} else {
		fmt.Println(totalMinimumHeightOfKnight)

	}
}
