// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"path/filepath"

	"github.com/GoLangsam/do/cli/cancel"
)

func main() {
	_, _ = cancel.WithCancel() // TODO pass ctx on

	flag.Parse()

	for _, arg := range flag.Args() {
		matches, err := filepath.Glob(arg)
		if err != nil {
			fmt.Println("Error:", err.Error())
		} else {
			for _, match := range matches {
				fmt.Println(match)
			}
		}
	}
}
