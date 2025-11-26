package demos

import (
	"fmt"
)

func CompositeDemo() {
	numbers := []int{10, 20, 30, 40, 50}
	println(len(numbers))
	println(numbers[0])

	countries := []string{"USA", "MEX", "CAN"}
	countries = append(countries, "CR")

	println(len(countries))
	println(countries[3])

	products := map[string]float64{
		"Laptop": 999.99,
		"Phone":  499.49,
		"Tablet": 299.29,
	}

	if price, exists := products["Phone"]; exists {
		fmt.Printf("Phone price: %.2f\n", price)
	}
}
