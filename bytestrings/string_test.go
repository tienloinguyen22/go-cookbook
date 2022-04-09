package bytestrings

import (
	"testing"
)

func TestSearchString(t *testing.T) {
	t.Run("Base case", func(t *testing.T) {
		workWithBuffer()
	})
}

func TestModifyString(t *testing.T) {
	t.Run("Base case", func(t *testing.T) {
		modifyString()
	})
}

func TestMyStrReader(t *testing.T) {
	t.Run("Base case", func(t *testing.T) {
		myStrReader()
	})
}