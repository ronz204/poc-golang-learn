package demos

import "fmt"

type Contact struct {
	Email string
	Phone string
}

type Customer struct {
	Name    string
	Contact Contact
	Active  bool
}

func (c *Customer) Display() {
	fmt.Println("Name:", c.Name)
	fmt.Println("Active:", c.Active)
	fmt.Println("Email:", c.Contact.Email)
	fmt.Println("Phone:", c.Contact.Phone)
}

func StructuringDemo() {
	customer1 := Customer{
		Name: "Alice",
		Contact: Contact{
			Email: "alice@example.com",
			Phone: "123-456-7890",
		},
		Active: true,
	}

	customer1.Display()
}
