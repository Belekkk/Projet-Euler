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
	res := 2
	for i := 3; i < 2000000; i += 2 {
		if isPrime(i) {
			res += i
		}
	}
	fmt.Println(res)
}
