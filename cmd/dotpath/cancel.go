// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"os"
)

//!+1
var cancel = make(chan struct{})

func init() {
	// Cancel traversal when input is detected.
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		close(cancel)
	}()
}

// usage:
//	select {
//	case <-cancel: // abort (after drain, if need)
//

// convenience:
func Cancelled() bool {
	select {
	case <-cancel:
		return true
	default:
		return false
	}
}
