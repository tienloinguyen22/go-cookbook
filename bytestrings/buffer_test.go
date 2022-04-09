package bytestrings

import (
	"bytes"
	"strings"
	"testing"
)

func TestMyBuffer(t *testing.T) {
	t.Run("Base case", func(t *testing.T) {
		str := "my test string"
		buff := myBuffer(str)
		if buff.String() != str {
			t.Errorf("myBuffer error, expected %v\n", bytes.NewBufferString(str))
		}
	})
}

func TestToString(t *testing.T) {
	t.Run("Base case", func(t *testing.T) {
		str := "my test string"
		r := strings.NewReader(str)
		result, err := toString(r)
		if err != nil {
			t.Errorf("toString error %v\n", err)
		}
		if result != str {
			t.Errorf("toString error, receive %v, expect %v\n", result, str)
		}
	})
}