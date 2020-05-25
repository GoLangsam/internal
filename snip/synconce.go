// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package snip

import (
	"image"
	"sync"
)

// Option 1 - on init
// fill iconsI during initialisaton => Concurrencysafe.
var iconsI = map[string]image.Image{
	"spades.png":   loadIcon("spades.png"),
	"hearts.png":   loadIcon("hearts.png"),
	"diamonds.png": loadIcon("diamonds.png"),
	"clubs.png":    loadIcon("clubs.png"),
}

// Concurrencysafe.
func IconI(name string) image.Image { return iconsI[name] }

// ===========================================================================
// Option 2 - sync.Once

var loadIconsOnce sync.Once
var icons map[string]image.Image

func loadIcon(name string) image.Image {
	return nil // just a dummy!
}

// ===========================================================================
func loadIcons() {
	icons = make(map[string]image.Image)
	icons["spades.png"] = loadIcon("spades.png")
	icons["hearts.png"] = loadIcon("hearts.png")
	icons["diamonds.png"] = loadIcon("diamonds.png")
	icons["clubs.png"] = loadIcon("clubs.png")
}

// Concurrencysafe.
func Icon(name string) image.Image {
	loadIconsOnce.Do(loadIcons)
	return icons[name]
}

/*
Each call to Do(loadIcons) locks the mutex and checks the boolean variable.

In the first call, in which the variable is false, "Do"
- calls "loadIcons" and
- sets the variable to true.
Subsequent calls do nothing, but the mutex synchronization ensures that the effects of loadIcons on
memory (specifically, icons) become visible to all goroutines.

Using sync.Once in this way, we can avoid sharing variables with other goroutines
until they have been properly constructed.
*/
