// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dotpath

import (
	"fmt"
)

func (dp *DotPath) Print() {
	fmt.Println("Got:\t", dp.String())
	rp := dp.RecursePathS()
	fmt.Println("=> Path:\t", dp.Path())
	if len(rp) > 0 {
		fmt.Println("=> Name:\t", dp.PathName())
		fmt.Println("Recurse:\t", rp)
	}
	wd := dp.WaydownPathS()
	if len(wd) > 0 || dp.Path() != dp.PathBase() {
		fmt.Println("=> Base:\t", dp.PathBase())
		fmt.Println("WayDown:\t", dp.PathS())
	}
}
