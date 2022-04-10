package filesdirs

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func operate() error {
	if err := os.Mkdir("./filesdirs/example", os.FileMode(0755)); err != nil {
		return err
	}

	if err := os.Chdir("./filesdirs/example"); err != nil {
		return err
	}

	file, err := os.Create("./text.txt")
	if err != nil {
		return err
	}

	// Write content to file
	value := []byte("Donec sollicitudin molestie malesuada. Curabitur non nulla sit amet nisl tempus convallis quis ac lectus.")
	count, err := file.Write(value)
	if err != nil {
		return err
	}
	if count != len(value) {
		return errors.New("incorrect length returned from write")
	}
	if err := file.Close(); err != nil {
		return err
	}

	// Read content from file
	file, err = os.Open("./text.txt")
	if err != nil {
		return err
	}
	io.Copy(os.Stdout, file)
	fmt.Println()
	if err := file.Close(); err != nil {
		return err
	}

	// Clean up
	if err := os.Chdir(".."); err != nil {
		return err
	}
	if err := os.RemoveAll("./example"); err != nil {
		return err
	}

	return nil
}