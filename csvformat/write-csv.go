package csvformat

import (
	"bytes"
	"encoding/csv"
	"io"
	"os"
)

type book struct {
	title string
	author string
}

type books []book

var b = books{
	book{
		author: "F Scott Fitzgerald",
		title: "The Great Gatsby",
	},
	book{
		author: "J D Salinger",
		title: "The Catcher in the Rye",
	},
}

// Conver an array of book => CSV format & Write to "w"
func (b *books) toCsv(w io.Writer) error {
	n := csv.NewWriter(w)

	err := n.Write([]string{"Title", "Author"})
	if err != nil {
		return err
	}

	for _, book := range *b {
		err := n.Write([]string{book.title, book.author})
		if err != nil {
			return err
		}
	}
	n.Flush()

	return nil
}

func writeCsvOutput() error {
	return b.toCsv(os.Stdout)
}

func writeCsvBuffer() (*bytes.Buffer, error) {
	w :=  &bytes.Buffer{}
	err := b.toCsv(w)
	return w, err
}