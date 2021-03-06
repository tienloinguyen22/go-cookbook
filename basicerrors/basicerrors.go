package basicerrors

import (
	"errors"
	"fmt"
)

type TypedError struct {
	error
}

func basicErrors() {
	err := errors.New("this is a quick & easy way to create an error")
	fmt.Println("errors.New: ", err)

	err = fmt.Errorf("an error occurred: %s", "something")
	fmt.Println("fmt.Errorf: ", err)

	err = TypedError{errors.New("typed error")}
	fmt.Println("typed error: ", err)
}