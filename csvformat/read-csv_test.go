package csvformat

import (
	"bytes"
	"strings"
	"testing"
)

func TestReadCsv(t * testing.T) {
	t.Run("Base case", func(t *testing.T) {
		myCsv := `
			Title,Director,Released year
			Guardians of the Galaxy Vol.2,James Gunn,2017
			Star Wars: Episode VIII,Rian Johnson,2017
		`
		b := bytes.NewBufferString(strings.TrimSpace(myCsv))

		movies, err := readCsv(b)
		if err != nil {
			t.Errorf("readCsv() error = %v, expectedError = false", err)
		}
		if len(movies) != 2 {
			t.Errorf("readCsv() missing item, expected len(movies) = 2")
		}
		if movies[0].title != "Guardians of the Galaxy Vol.2" || movies[0].director != "James Gunn" || movies[0].year != 2017 {
			t.Errorf("readCsv() wrong item data, expected Guardians of the Galaxy Vol.2 | James Gunn | 2017. Retrieve %v", movies[0])
		}
		if movies[1].title != "Star Wars: Episode VIII" || movies[1].director != "Rian Johnson" || movies[1].year != 2017 {
			t.Errorf("readCsv() wrong item data, expected Star Wars: Episode VIII | Rian Johnson | 2017. Retrieve %v", movies[1])
		}
	})
}