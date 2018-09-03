package main

import (
	"fmt"
)

func main() {

	n := 600851475143
	f := 3
	lastPrimeFactor := 0
	for n > 1 {
		if n%f == 0 {
			lastPrimeFactor = f
			n /= f
		} else {
			f += 2
		}
	}

	fmt.Println(lastPrimeFactor)

}
