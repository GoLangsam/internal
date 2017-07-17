// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"

	"github.com/golangsam/internal/container/ccsafe/dotpath"
)

func main() {
	flag.Parse()

	ds := dotpath.FilePathS(flag.Args()...)
	fmt.Println("===============================================================================")
	for _, dp := range ds {
		dp.Print()
		fmt.Println("-------------------------------------------------------------------------------")
	}
}
