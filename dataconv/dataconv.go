package dataconv

import "fmt"

func showConv() {
	a := 24 // int
	b := 2.0 // float64

	// Convert int to float 64
	c := float64(a) + b
	fmt.Printf("c: %v\n", c)

	// Convert to string with Sprintf
	precision := fmt.Sprintf("%.2f", b)
	fmt.Printf("%v - %T\n", precision, precision)
}