package filesdirs

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func capitalizer(file1 *os.File, file2 *os.File) error {
	if _, err := file1.Seek(0, io.SeekStart); err != nil {
		return err
	}

	tmp := bytes.NewBuffer([]byte{})
	if _, err := io.Copy(tmp, file1); err != nil {
		return err
	}

	capitalizedStr := strings.ToUpper(tmp.String())
	if _, err := io.Copy(file2, strings.NewReader(capitalizedStr)); err != nil {
		return err
	}

	return nil
}

func tryCaplitalizer() error {
	// Create file1.txt with some content
	file1, err := os.Create("./file1.txt")
	if err != nil {
		return err
	}
	_, err = file1.Write([]byte("Curabitur non nulla sit amet nisl tempus convallis quis ac lectus. Lorem ipsum dolor sit amet, consectetur adipiscing elit."))
	if err != nil {
		return err
	}
	
	// Create file2.txt empty
	file2, err := os.Create("./file2.txt")
	if err != nil {
		return err
	}

	// Call capitalizer on both
	if err := capitalizer(file1, file2); err != nil {
		return err
	}

	// Print result to os.Stdout
	fmt.Print("file1.txt: ")
	file1.Seek(0, io.SeekStart)
	_, err = io.Copy(os.Stdout, file1)
	if err != nil {
		return err
	}
	fmt.Println()

	fmt.Print("file2.txt: ")
	file2.Seek(0, io.SeekStart)
	_, err = io.Copy(os.Stdout, file2)
	if err != nil {
		return err
	}
	fmt.Println()

	// Close both file & remove them
	if err := os.Remove("./file1.txt"); err != nil {
		return err
	}
	if err := os.Remove("./file2.txt"); err != nil {
		return err
	}

	return nil
}