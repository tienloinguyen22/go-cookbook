package csvformat

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type movie struct {
	title string
	director string
	year int
}

func readCsv(b io.Reader) ([]movie, error) {
	r := csv.NewReader(b)
	movies := []movie{}

	_, err := r.Read()
	if err != nil && err != io.EOF {
		return nil, err
	}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		year, err := strconv.ParseInt(record[2], 10, 64)
		if err != nil {
			return nil, err
		}

		m := movie{
			title: strings.TrimSpace(record[0]),
			director: strings.TrimSpace(record[1]),
			year: int(year)}
		movies = append(movies, m)
	}
	return movies, nil
}

// Conver a str in CSV format => An array of movie
func addMoviesFromText() error {
	myCsv := `
		Title,Director,Released year
		Guardians of the Galaxy Vol.2,James Gunn,2017
		Star Wars: Episode VIII,Rian Johnson,2017
	`
	b := bytes.NewBufferString(strings.TrimSpace(myCsv))

	movies, err := readCsv(b)
	if err != nil {
		return err
	}

	for _, movie := range movies {
		fmt.Printf("%v\n", movie)
	}
	return nil
}