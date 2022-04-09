package interfaces

import (
	"bytes"
	"fmt"
)

func RunInterfaces() {
	in := bytes.NewReader([]byte("example"))
	out := &bytes.Buffer{}

	fmt.Print("Stdout on copy = ")
	if err := myCopy(in, out); err != nil {
		panic(err)
	}
	fmt.Println("Out bytes buffer =", out.String())

	fmt.Print("Stdout on pipe = ")
	if err := myPipe(); err != nil {
		panic(err)
	}
}