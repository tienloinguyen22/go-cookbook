package bytestrings

import (
	"bytes"
	"io"
	"io/ioutil"
)

func myBuffer(str string) *bytes.Buffer {
	return bytes.NewBufferString(str)
}

func toString(r io.Reader) (string, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(b), nil
}