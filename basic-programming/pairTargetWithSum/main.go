package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 6}
	fmt.Println(PairSum(arr, 6))
}

/*
	target = a + b
	target - b = a
*/
func PairSum(arr []int, target int) (resultIndexs []int) {
	hmap := make(map[int]int)
	for index, b := range arr {
		var a = target - b // 2 = 6 - 4
		if _, ok := hmap[a]; ok {
			resultIndexs = append(resultIndexs, hmap[a], index) // 1 & 3
			break;
		}
		hmap[b] = index
	}
	return resultIndexs

}
