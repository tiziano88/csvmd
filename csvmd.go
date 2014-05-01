package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := csv.NewReader(os.Stdin)

	rs, err := r.ReadAll()
	if err != nil {
		panic(err)
	}

	// Find number of columns.
	colN := 0
	for _, row := range rs {
		if len(row) > colN {
			colN = len(row)
		}
	}

	// Find width of each field.
	width := make([]int, colN)
	for _, row := range rs {
		for j, field := range row {
			if len(field) > width[j] {
				width[j] = len(field)
			}
		}
	}

	// Go through each field and reformat it.
	for i, row := range rs {
		outRow := make([]string, colN)

		for j, field := range row {
			outRow[j] = fmt.Sprintf(" %*s ", width[j], field)
		}
		fmt.Printf("|%s|\n", strings.Join(outRow, "|"))

		// Header row.
		if i == 0 {
			for j, _ := range row {
				outRow[j] = fmt.Sprintf(" %*s ", width[j], strings.Repeat("-", width[j]))
			}
			fmt.Printf("|%s|\n", strings.Join(outRow, "|"))
		}
	}
}
