package interfaces

import (
	"fmt"
	"io"
	"os"
)

func myCopy(in io.ReadSeeker, out io.Writer) error {
	// Write to "out" & "stdout"
	writer := io.MultiWriter(out, os.Stdout)
	if _, err := io.Copy(writer, in); err != nil {
		return err
	}
	in.Seek(0, 0)

	buf := make([]byte, 64)
	if _, err := io.CopyBuffer(writer, in, buf); err != nil {
		return err
	}

	fmt.Println()
	return nil
}