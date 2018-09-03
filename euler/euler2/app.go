package main

import (
	"fmt"
)

func main() {
	s := 0
	a := 0
	b := 1
	for a < 4000000 {
		if a%2 == 0 {
			s += a
		}
		a, b = b, a+b
	}
	fmt.Println(s)
}
