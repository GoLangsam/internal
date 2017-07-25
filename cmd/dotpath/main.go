// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"

	"github.com/golangsam/do/cli/cancel"
	"github.com/golangsam/internal/container/ccsafe/dotpath"
)

func main() {
	_ = cancel.Canceler() // TODO pass ctx on

	flag.Parse()

	ds := dotpath.FilePathS(flag.Args()...)
	fmt.Println("===============================================================================")
	for _, dp := range ds {
		dp.Print()
		fmt.Println("-------------------------------------------------------------------------------")
	}
}
