package main

import (
	"fmt"
)

func isPalindrome(n int) bool {
	r := 0
	o := n
	for n > 0 {
		r = r*10 + n%10
		n /= 10
	}
	return o == r
}

func main() {
	m := 0
	for a := 0; a < 1000; a++ {
		for b := 0; b < 1000; b++ {
			p := a * b
			if isPalindrome(p) && p > m {
				m = p
			}
		}
	}
	fmt.Println(m)
}
