package main

import (
	"fmt"
	"math"
)

func divisors(n int) int {
	var c int = 0

	for i := int(1); i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			c = c + 2
		}
	}
	return c
}

func main() {
	var i, triangle int = 0, 1
	var nextTriangle int

	for {
		nextTriangle = i + triangle
		if divisors(nextTriangle) > 500 {
			fmt.Println(nextTriangle)
			break
		}
		i = nextTriangle
		triangle++
	}

}
