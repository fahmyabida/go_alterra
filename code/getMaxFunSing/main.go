package main

import (
	"fmt"
	"sort"
)

func main() {
	getMaxFun([]int{1, 1, 2}, []int{5, 4, 3})
	getMaxFun([]int{1, 2, 3, 4}, []int{3, 2, 1, 4})

}

func getMaxFun(singer, length []int) {
	songs := make(map[int][]int)
	for i := 0; i < len(singer); i++ {
		j := songs[singer[i]]
		j = append(j, length[i])
		songs[singer[i]] = j
	}
	tot := 0
	q := []int{}
	for i, _ := range songs {
		sort.Ints(songs[i])
		q = append(q, songs[i][0])
		for j := 1; j < len(songs[i]); j++ {
			tot += songs[i][j]
		}
	}
	ans, t := 0, 0
	sort.Ints(q)
	for i, _ := range q {
		t += 1
		ans += t * q[i]
	}
	ans += t * tot
	fmt.Println(ans)

}
