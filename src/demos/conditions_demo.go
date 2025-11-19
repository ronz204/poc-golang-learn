package demos

type User struct {
	Name   string
	Passwd string
}

func ConditionsDemo() {

	if age := 18; age > 18 {
		println("Adult")
	}

	if temperature := 30; temperature > 25 {
		println("It's a hot day")
	} else {
		println("It's a cold day")
	}

	user1 := User{Name: "admin", Passwd: "1234"}

	if user1.Name == "admin" && user1.Passwd == "1234" {
		println("Access granted")
	}

	isWeekend, isHoliday, isRaining := true, false, false

	if isWeekend || isHoliday {
		println("You can relax today")
	}

	if !isRaining {
		println("You can go for a walk")
	}
}
