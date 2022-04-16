package basicerrors

import "fmt"

func RunBasicerrors() {
	basicErrors()

	err := myCustomError()
	fmt.Println("custom error: ", err)
}