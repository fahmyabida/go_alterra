package main

import "fmt"

func main() {
	fmt.Println(palindrome("abcdeffedcba"))
	
}

func palindrome(input string) (isPalindrome bool) {
	for i := 0; i < len(input); i++ {
		if input[i] != input[len(input)-(i+1)] {
			return false
		}
	}
	return true
}
