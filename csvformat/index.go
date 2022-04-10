package csvformat

import "fmt"

func RunCsvformat() {
	if err := addMoviesFromText(); err != nil {
		panic(err)
	}

	if err := writeCsvOutput(); err != nil {
		panic(err)
	}

	buff, err := writeCsvBuffer()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Buffer = %v\n", buff.String())
}