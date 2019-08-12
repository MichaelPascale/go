package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	if x == 0 {
		return 0
	}

	z := float64(1)

	for i := 0; i < 25; i++ {
		z -= (z*z - x) / (2 * z)
	}

	return z
}

func main() {
	for n := 0; n < 10; n++ {
		fmt.Println(Sqrt(float64(n)))
	}
}
