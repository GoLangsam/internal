// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.package dotpath

package dotpath

const (
	Dot       = `.`
	SingleDot = Dot // just another name
	DoubleDot = Dot + Dot
	TripleDot = Dot + Dot + Dot
	Empty     = ``
)

func init() { // some paranoid sanity checks ;-)
	if len(Empty) != 0 {
		panic("My empty '" + Empty + "' has non-zero length!")
	}

}
