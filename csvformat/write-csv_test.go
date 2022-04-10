package csvformat

import (
	"strings"
	"testing"
)

func TestWriteCsv(t * testing.T) {
	t.Run("Base case", func(t *testing.T) {
		buff, err := writeCsvBuffer()
		if err != nil {
			t.Errorf("writeCsv() error = %v, expectedError = false", err)
		}
		
		expectedResult := strings.TrimSpace(`Title,Author
The Great Gatsby,F Scott Fitzgerald
The Catcher in the Rye,J D Salinger`)
		result := strings.TrimSpace(buff.String())
		if result != expectedResult {
			t.Errorf("writeCsv() wrong result. Expected: %v. Received: %v", expectedResult, result)
		}
	})
}