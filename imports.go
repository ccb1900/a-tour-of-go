package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("you have some %g problems.", math.Nextafter(2, 3))
}
