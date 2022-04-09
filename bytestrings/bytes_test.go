package bytestrings

import (
	"testing"
)

func TestWorkWithBuffer(t *testing.T) {
	t.Run("Base case", func(t *testing.T) {
		err := workWithBuffer()
		if err != nil {
			t.Errorf("workWithBuffer error %v\n", err)
		}
	})
}