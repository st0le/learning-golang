package main

import (
	"fmt"
)

func sieve(n int) []bool {
	isprime := make([]bool, n+1)
	for i := 2; i < n; i += 2 {
		isprime[i] = false
		isprime[i+1] = true
	}

	isprime[2] = true
	for i := 3; i*i <= n; i += 2 {
		for j := i * i; j <= n; j += 2 * i {
			isprime[j] = false
		}
	}

	return isprime
}

func main() {

	primes := sieve(150000)

	c := 0
	for i, v := range primes {
		if v {
			c++
			if c == 10001 {
				fmt.Println(c, i)
			}
		}
	}

}
