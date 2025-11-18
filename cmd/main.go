package main

import (
	"fmt"
	"poc-golang-learn/src/helpers"
)

func main() {
	randomNum := helpers.RandomInt(1, 100)
	fmt.Printf("Random number: %d\n", randomNum)
}
