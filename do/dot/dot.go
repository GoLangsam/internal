// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dot

import (
	"github.com/GoLangsam/do/ats"
)

// Dot defines what we need from *dot.Dot
type Dot interface {
	//	Lock()	 // sync.Locker
	//	Unlock()	// sync.Locker

	String() string // fmt.Stringer

	// from Tag
	//	K() string	// no need - we use String()
	//	V() string	// no need - Value as string
	GetV() interface{}   // Get Value
	Tag(val interface{}) // Set Value

	// from StringMap
	UnlockedGet(key string) (interface{}, bool)
	UnlockedAdd(key string, val ...string) (interface{}, bool)

	SeeError(myName, myThing string, err error) bool
	SeeNotOk(myName, myThing string, ok bool, complain string) bool
}

// helpers

// access child

func lookupDot(d Dot, key string) Dot {
	any, ok := d.UnlockedGet(key)
	if d.SeeNotOk("UnlockedGet for key=", key, ok, " returned false?!?") {
		panic("UnlockedGet for key=" + key + " returned false?!?")
	}
	dot, ok := any.(Dot)
	if d.SeeNotOk("UnlockedGet", key, ok, "What? No dot returned?!?") {
		panic("UnlockedGet: No dot returned for key=" + key + "?!?")
	}
	return dot
}

// Value-handlers
func v(d Dot) string {
	return ats.GetString(d.GetV())
}

func vNonEmpty(d Dot, myName string) (string, bool) {
	value := v(d)
	switch {
	case value == "":
		d.SeeNotOk(myName, d.String(), false, "my Value must not be empty!")
		return "", false
	default:
		return value, true
	}
}
