package basicerrors

import "testing"

func TestCustomError(t *testing.T) {
	t.Run("Base case", func(t *testing.T) {
		c := CustomError{
			Result: "mine",
		}
		r := "There was an error; mine was the result"
		if c.Error() != r {
			t.Errorf("c.Error(). Expect: %v. Receive: %v", r, c.Error())
		}
	})
}