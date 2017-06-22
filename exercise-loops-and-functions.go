package main

import (
	"fmt"
	"math"
)

/**
牛顿法求解平方根
 */

func main() {
	fmt.Println(Sqrt(3))
	fmt.Println(math.Sqrt(3))
}

func Sqrt(x float64) float64 {
	z := x / 2
	for i := 0; i < 100; i++ {
		t := z
		z = z - (z*z-x)/2/z

		if math.Abs(t-z) < 0.000000001 {
			return z
		}
		fmt.Println(i, z)
	}

	return z
}
