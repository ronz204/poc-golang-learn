package demos

func Operate(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

func CallbacksDemo() {
	add := func(x, y int) int {
		return x + y
	}

	result := Operate(3, 4, add)
	println("Addition result:", result)
}
