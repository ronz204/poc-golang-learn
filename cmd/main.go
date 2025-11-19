package main

import (
	"fmt"
	"poc-golang-learn/src/helpers"
)

func main() {
	min, max := 1, 100

	random := helpers.RandomInt(min, max)
	fmt.Printf("Random number: %d\n", random)
}
