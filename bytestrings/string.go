package bytestrings

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func searchString() {
	str := "this is a test"

	// returns true because s contains the word this
	fmt.Printf("Contains 'this': %v\n", strings.Contains(str, "this"))

	// returns true because s contains the letter a would also match if it contained b or c
	fmt.Printf("Contains any 'abc': %v\n", strings.ContainsAny(str, "abc"))
}

func modifyString() {
	str := "simple string"

	// Split by space. Return an array
	fmt.Printf("Split: %v\n", strings.Split(str, " "))

	// Title case
	fmt.Printf("Title: %v\n", strings.Title(str))

	// Remove all spaces
	fmt.Printf("Trim spaces: %v\n", strings.TrimSpace("     " + str + "    "))
}

func myStrReader() {
	str := "simple stringn"
	strReader := strings.NewReader(str)
	io.Copy(os.Stdout, strReader)
}