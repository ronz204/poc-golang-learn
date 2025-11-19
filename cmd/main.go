package main

import (
	"fmt"
	"poc-golang-learn/src/helpers"
)

func main() {
	random := helpers.RandomInt(1, 10)
	fmt.Printf("Random number: %d\n", random)
}
