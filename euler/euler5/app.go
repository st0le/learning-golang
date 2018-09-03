package main

import (
	"fmt"
)

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	l := 1
	for i := 1; i <= 20; i++ {
		l = lcm(l, i)
	}
	fmt.Println(l)
}
