package main

import "fmt"

func isPythTriplet(a int, b int, c int) bool {
	if a*a+b*b == c*c {
		return true
	}
	return false
}

func main() {
	for a := 1; a <= 500; a++ {
		for b := a; b <= 500; b++ { // b > a
			// c had to be higher to sum(a, b)
			c := 1000 - (a + b)
			if isPythTriplet(a, b, c) {
				fmt.Println(a * b * c)
				break
			}
		}
	}
}
