package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(10 - 20)
	fmt.Println(int(math.Abs(10 - 20)))
	// fmt.Println(fiboAdvance(4))
}

func fiboAdvance(n int) int {

	if n <= 1 {
		return n
	}
	// n = 2 --> 1
	// (n-2) = 0
	// (n-1) = 1
	fiboIMinSatu, fiboIMinDua, fiboI := 1, 0, -1
	for i := 2; i <= n; i++ {
		fiboI = fiboIMinSatu + fiboIMinDua
		fiboIMinDua = fiboIMinSatu
		fiboIMinSatu = fiboI
	}
	return fiboI
}

// 0,1,1,2,3,5,?
