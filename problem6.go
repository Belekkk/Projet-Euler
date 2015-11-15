package main

import (
	"fmt"
	"math"
)

func main() {
	var sumSquares int = 0

	for i := 0; i <= 100; i++ {
		sumSquares += int(math.Pow(float64(i), 2))
	}

	var sum int = 0

	for i := 0; i <= 100; i++ {
		sum += i
	}
	squareSum := int(math.Pow(float64(sum), 2))

	diff := squareSum - sumSquares
	fmt.Println(sumSquares, squareSum, diff)
}
