package main

import "fmt"

func main() {
	sum := 1

	for sum < 511 {
		sum += sum
	}

	fmt.Println(sum)
}
