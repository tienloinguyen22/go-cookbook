package dataconv

import "fmt"

func checkType(s interface{}) {
	switch s.(type) {
	case string:
		fmt.Println("It's a string")
	case int:
		fmt.Println("It's an int")
	default:
		fmt.Println("Not sure what it is ...")
	}
}

func myInterfaces() {
	checkType("test")
	checkType(1)
	checkType([]byte{})

	var i interface{}
	i = "test"

	if val, ok := i.(string); ok {
		fmt.Println("val is ", val)
	}

	if _, ok := i.(int); !ok {
		fmt.Println("uh ok. glad we handled this")
	}
}