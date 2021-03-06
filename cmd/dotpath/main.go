// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"

	"github.com/GoLangsam/do/cli/cancel"
	"github.com/GoLangsam/internal/container/ccsafe/dotpath"
)

func main() {
	_, _ = cancel.WithCancel() // TODO pass ctx on

	flag.Parse()

	ds := dotpath.FilePathS(flag.Args()...)
	fmt.Println("===============================================================================")
	for _, dp := range ds {
		dp.Print()
		fmt.Println("-------------------------------------------------------------------------------")
	}
}
