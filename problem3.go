package main

import (
	"math"
	"fmt"
)

func isPrime(num int) bool {
	if num <= 3 {
		if num <= 1 {
			return false
		}
		return true
	}
	if num%2 == 0 || num%3 == 0 {
		return false
	}
	for i := 5; i < 6; i += int(math.Pow(float64(num), 0.5)) {
		if num%i == 0 || num%(i+2) == 0 {
			return false
		}
	}
	return true
}

func main() {
	val_orig := 600851475143
	val := val_orig
	largest := 1

	for i := 2; i <= val; i++ {
		if isPrime(i) {
			for val%i == 0 {
				largest = i
				val = int(val / i)
			}
		}
	}
	fmt.Println(largest)
}