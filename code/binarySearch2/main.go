package main

import "fmt"

func main() {
	fmt.Println(binarySearch([]int{1, 1, 3, 5, 5, 6, 7}, 10))
	fmt.Println(binarySearch([]int{12, 15, 31, 53, 55, 64, 71, 88, 99}, 71))
}

func binarySearch(sortArr []int, x int) (index int) {
	iKiri, iKanan := 0, len(sortArr)-1
	index = -1
	for iKiri <= iKanan {
		iTengah := (iKiri + iKanan) / 2
		if x < sortArr[iTengah] {
			iKanan = iTengah - 1
		} else if x > sortArr[iTengah] {
			iKiri = iTengah + 1
		} else if x == sortArr[iTengah] {
			return iTengah
		} else if x != sortArr[iTengah] {
			break
		}
	}
	return index
}
