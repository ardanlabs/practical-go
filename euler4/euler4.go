package main

import (
	"fmt"
	"time"
)

/*
A palindromic number reads the same both ways. The largest palindrome made from the product of two
2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.

Answer: 906609

Check for palindrome: Convert to string using fmt.Sprintf and then go over the digits
*/

func main() {
	start := time.Now()
	maxP := 0
	for i := 100; i <= 999; i++ {
		for j := 100; j <= 999; j++ {
			n := i * j
			if n > maxP && isPalindrome(n) {
				maxP = n
			}
		}
	}
	fmt.Println("duration:", time.Since(start))
	fmt.Println(maxP)
}

func isPalindrome(n int) bool {
	s := fmt.Sprintf("%d", n)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}

	return true
}
