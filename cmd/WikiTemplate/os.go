// Copyright 2016 Andreas Pannewitz. All rights reserved.

package WikiTemplate

import (
	"os"
)

var myHomePath = fullname(os.Getenv("priv"), "!!Wiki!!")

// Open returns a R-ONLY file 'basename' from myRoot
func Open(basename string) *os.File {
	infile, err := os.Open(fullname(myPathRaw, basename))
	checkerr(err, "Err on Open")
	return infile
}

// Make returns a new RW file 'basename' inside myRoot
func Make(basename string) *os.File {
	outfile, err := os.Create(fullname(myPathNew, basename))
	checkerr(err, "Err on Create")
	_, err = outfile.Write(Utf8BOM)
	checkerr(err, "Err on writing UTF-8 BOM")
	return outfile
}
