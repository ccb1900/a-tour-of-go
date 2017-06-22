package main

import (
	"math"
	"fmt"
)

func main() {
	fmt.Println(
		powelse(3, 2, 10),
		powelse(3, 3, 20),
	)
}

func powelse(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >=%g\n", v, lim)
	}

	return lim
}
