package main

import (
	"fmt"
	"math"
)

// Sqrt is a function that takes arguments for the numbers to be calculated sqrt and returns the respective values
func Sqrt(x float64) float64 {
	z := 1.0
	prev := z

	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		if math.Abs(z-prev) < 1e-18 {
			break
		}

		prev = z
	}

	return z
}

func main() {
	fmt.Println("Output of operation performed using Sqrt: ", Sqrt(2))
	fmt.Println("Output of operation performed using math.Sqrt: ", math.Sqrt(2))
}
