package demos

import "fmt"

type Monster struct {
	Flavor string
}

func (m *Monster) Scare() {
	fmt.Println("Boo! I'm a", m.Flavor, "monster!")
}

func PointersDemo() {
	number := 10
	direction := &number

	fmt.Println("Value of number:", number)
	fmt.Println("Address of number:", direction)
	fmt.Println("Value at address stored in direction:", *direction)

	monster := Monster{Flavor: "Khaos"}
	monster.Scare()
}
