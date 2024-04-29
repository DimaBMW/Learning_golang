package main

import "fmt"

func main(){
	fmt.Println(isPalindrome(121))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	reversed := 0
	original := x
	for original > 0 {
		reversed = reversed*10 + original%10
		original /= 10
	}
	return x == reversed
} 

