package demos

import (
	"errors"
	"fmt"
)

type InvalidNameError struct {
	Name string
}

func (e *InvalidNameError) Error() string {
	return fmt.Sprintf("invalid name: %s", e.Name)
}

type InvalidAgeError struct {
	Age int
}

func (e *InvalidAgeError) Error() string {
	return fmt.Sprintf("invalid age: %d", e.Age)
}

func RegisterUser(name string, age int) error {
	if name == "" {
		return &InvalidNameError{Name: name}
	}
	if age < 0 {
		return &InvalidAgeError{Age: age}
	}
	return nil
}

func ErrorsDemo() {
	err := RegisterUser("", 25)

	if err == nil {
		return
	}

	var ageErr *InvalidAgeError
	var nameErr *InvalidNameError

	if errors.As(err, &ageErr) {
		fmt.Println("Age error occurred:", ageErr)
	}

	if errors.As(err, &nameErr) {
		fmt.Println("Name error occurred:", nameErr)
	}
}
