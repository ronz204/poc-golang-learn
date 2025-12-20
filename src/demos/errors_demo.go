package demos

import "errors"

var (
	ErrInvalidName = errors.New("invalid name provided")
	ErrInvalidAge  = errors.New("invalid age provided")
)

func RegisterUser(name string, age int) error {
	if name == "" {
		return ErrInvalidName
	}
	if age < 0 {
		return ErrInvalidAge
	}
	return nil
}

func ErrorsDemo() {
	err := RegisterUser("", 25)

	if errors.Is(err, ErrInvalidName) {
		println("Error: The name provided is invalid.")
	}

	if errors.Is(err, ErrInvalidAge) {
		println("Error: The age provided is invalid.")
	}
}
