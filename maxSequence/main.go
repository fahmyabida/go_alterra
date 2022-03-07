package main

import (
	"fmt"
	"sort"
)

func main() {
	MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}) // 6
	// MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6})   // 7
	// MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3})   // 7
	// MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6})   // 8
	// MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6})    // 12
	// MaxSequence([]int{-2, 9, 6, -2, 9, 8, -2, 3, 1, 6, -6})   // 29
}

func MaxSequence(arr []int) int {
	sum := []int{}
	hmap := make(map[int][]int)
	for i := 0; i < len(arr); i++ {
		arrGroup := []int{arr[i]}
		tmpSum := arr[i]
		for j := i + 1; j < len(arr); j++ {
			arrGroup = append(arrGroup, arr[j])
			tmpSum += arr[j] // tmpSum = tmpSum + arr[j]
		}
		sum = append(sum, tmpSum)
		hmap[tmpSum] = arrGroup
		// ---
		tempArr := arr[:len(arr)-1]
		for len(tempArr) != i {
			arrGroup = []int{arr[i]}
			tmpSum = arr[i]
			for j := i + 1; j < len(tempArr); j++ {
				arrGroup = append(arrGroup, tempArr[j])
				tmpSum += tempArr[j]
			}
			sum = append(sum, tmpSum)
			hmap[tmpSum] = arrGroup
			tempArr = tempArr[:len(tempArr)-1]
		}

	}
	sort.Slice(sum, func(i, j int) bool {
		return sum[i] > sum[j]
	})
	fmt.Println(sum[0])
	fmt.Println(hmap[sum[0]])
	return sum[0]
}
