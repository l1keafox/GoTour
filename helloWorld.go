package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	for z := float64(1); z > 0; z++ {
		//fmt.Println((z * z) - x)
		z -= (z*z - x) / (2 * z)

		fmt.Println(z, x, math.Abs(x-z))
	}

	return 0

}

func main() {
	fmt.Println(Sqrt(2))

}
