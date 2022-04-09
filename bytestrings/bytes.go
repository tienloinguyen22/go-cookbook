package bytestrings

import (
	"bufio"
	"bytes"
	"fmt"
)

func workWithBuffer() error {
	str := "it's easy to encode unicode into a byte array"
	bytesBuff := myBuffer(str)
	fmt.Println("Bytes buff .String(): " + bytesBuff.String())

	originStr, err := toString(bytesBuff)
	if err != nil {
		return err
	}
	fmt.Println("Using toString(): " + originStr)

	r := bytes.NewReader([]byte(str))
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Print(scanner.Text() + "|")
	}
	fmt.Println()

	return nil
}