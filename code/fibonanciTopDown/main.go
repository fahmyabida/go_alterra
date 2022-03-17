package main

import "fmt"

func main() {
	fmt.Println(fibonanciTopDown(4))
}

// key : idx fibo || value : nilai fibonanci
var memoTopDown = map[int]int{}

func fibonanciTopDown(n int) (fiboNumber int) {
	// kalau sudah pernah dihitung, coba di cek
	valueFibo, isAlreadyCount := memoTopDown[n]
	if isAlreadyCount {
		return valueFibo
	}
	// kalau blm pernah maka lanjut / dihitung
	if n <= 1 {
		memoTopDown[n] = n
	} else {
		memoTopDown[n] = fibonanciTopDown(n - 1) + fibonanciTopDown(n - 2)
	}
	return memoTopDown[n]
}
