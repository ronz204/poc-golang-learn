package demos

type Contact struct {
	Email string
	Phone string
}

type Customer struct {
	Name string
	Contact
	Active bool
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

	println("Customer Name:", customer1.Name)
	println("Is Active:", customer1.Active)
}
