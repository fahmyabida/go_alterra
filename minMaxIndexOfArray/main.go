package main

import "fmt"

func main() {
	min, minIndex, max, maxIndex := getMinMaxAndTheIndex([]int{5, 7, 4, -2, -1, 0})
	fmt.Printf("min: %v index %v | max: %v index %v \n", min, minIndex, max, maxIndex)
}

func getMinMaxAndTheIndex(numbers []int) (min, minIndex, max, maxIndex int) {
	min, max = numbers[0], numbers[0]
	for i, v := range numbers {
		if max < v {
			max = v
			maxIndex = i
		}
		if min > v {
			min = v
			minIndex = i
		}
	}

	return min, minIndex, max, maxIndex
}
