// Copyright 2016 Andreas Pannewitz. All rights reserved.

package WikiTemplate

import (
	"encoding/csv"
)

func read(filename string) [][]string {
	infile := Open(filename)
	defer infile.Close()

	r := csv.NewReader(infile)
	r.Comma = csvComma
	r.Comment = csvComment
	r.LazyQuotes = csvLazyQuotes

	records, err := r.ReadAll()
	checkerr(err, "Err on ReadAll")

	return records
}
