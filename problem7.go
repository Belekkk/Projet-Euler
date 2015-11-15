package main

import (
	"fmt"
	"math"
)

func isPrime(num int) bool {
	if num == 1 {
		return false
	}
	if num%2 == 0 {
		return false
	}
	for i := 3; i < int(math.Sqrt(float64(num)))+1; i += 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	nbPrime := 1
	k := 1

	for nbPrime != 10001 {
		k += 2
		if isPrime(k) {
			nbPrime += 1
		}
	}
	fmt.Println(k)
}
