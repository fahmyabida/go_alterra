package main

import "fmt"

func main() {
	fmt.Println(lastStoneWeight([]int{1, 2, 3, 6, 7, 7}))
	fmt.Println(lastStoneWeight([]int{2, 4, 5}))
}

func lastStoneWeight(arrStone []int) (lastStone int) {
	sisaPecahanBatu := 0
	for i := len(arrStone)-1; i > 0; i-- {
		if sisaPecahanBatu == 0 {
			if arrStone[i] >= arrStone[i-1] {
				sisaPecahanBatu = arrStone[i] - arrStone[i-1]
			} else {
				sisaPecahanBatu = arrStone[i-1] - arrStone[i]
			}
		} else if sisaPecahanBatu > 0 {
			if sisaPecahanBatu >= arrStone[i-1] {
				sisaPecahanBatu = sisaPecahanBatu - arrStone[i-1]
			} else {
				sisaPecahanBatu = arrStone[i-1] - sisaPecahanBatu
			}
		}
		if sisaPecahanBatu == 0 {
			i--
			continue
		}
	}
	return sisaPecahanBatu
}
