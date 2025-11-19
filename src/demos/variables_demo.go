package demos

import "fmt"

func VariablesDemo() {

	int8Var := 127
	fmt.Println("int8:", int8Var)

	int16Var := 32767
	fmt.Println("int16:", int16Var)

	int32Var := 2147483647
	fmt.Println("int32:", int32Var)

	int64Var := 9223372036854775807
	fmt.Println("int64:", int64Var)

	uint8Var := 255
	fmt.Println("uint8:", uint8Var)

	uint16Var := 65535
	fmt.Println("uint16:", uint16Var)

	uint32Var := 4294967295
	fmt.Println("uint32:", uint32Var)

	float32Var := 3.4028235e+38
	fmt.Println("float32:", float32Var)

	float64Var := 1.7976931348623157e+308
	fmt.Println("float64:", float64Var)

	isActive := true
	fmt.Println("bool:", isActive)

	character := 'R'
	fmt.Println("rune:", character)

	country := "USA"
	fmt.Println("string:", country)
}
