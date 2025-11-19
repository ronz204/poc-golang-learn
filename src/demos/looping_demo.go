package demos

import "fmt"

func LoopingDemo() {

	for i := 1; i <= 5; i++ {
		fmt.Printf("Iteration %d\n", i)
	}

	fruits := []string{"Apple", "Banana", "Cherry"}
	for index, fruit := range fruits {
		fmt.Printf("Fruit %d: %s\n", index, fruit)
	}

	ages := map[string]int{
		"Alice":   30,
		"Bob":     25,
		"Charlie": 35,
	}
	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}

	text := "Golang"
	for index, char := range text {
		fmt.Printf("   [%d] %c (Unicode: %d)\n", index, char, char)
	}
}
