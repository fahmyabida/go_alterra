// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println(binarySearch(70, items))
}

func binarySearch(n int, arr []int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := (low + high) / 2
		if arr[mid] < n {
			low = mid + 1
			continue
		}

		high = mid - 1
	}

	if arr[low] == n {
		return low
	}

	return -1
}
