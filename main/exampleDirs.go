// +build ignore

// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
	"path"
	"text/tabwriter"

	dirs "github.com/GoLangsam/internal/container/ccsafe/dotpath"
)

var _ = os.DevNull
var _ = path.ErrBadPattern
var _ = tabwriter.AlignRight

/*
func Example() {
	paths := []string{
			"a/",
			"a/c",
			"a//c",
			"a/c/.",
			"a/c/b/..",
			"/../a/c",
			"/../a/b/../././/c",
			"/a/b/c/bar.css/css",
			"/a/b/c/bar.css/.css",
			"/a/b/c/bar.css/",
			"/a/b/c/bar.css/.",
			"/a/b/c/bar.css/..",
			"/a/b/c/bar.css/../.",
			"/a/b/c/../bar.css",
			"/a/b/c/../bar.css/",
			"/a/b/c/../bar.css/.",
			"/a/b/c/../bar.css/..",
		"/a/b/c/../bar.css/.../.",
		"/a/b/c/.../bar.css/..../..",
		"/a/b/c/../.../bar.css/...../...",
		"/a/b/c/..../..../bar.css/....../....",
		"x...;y/z/.../..;/a/b/c/../bar.css/......./.....",
	}
	_ = tabwriter.FilterHTML                   // Ignore html tags and treat entities (starting with '&' and ending in ';') as single characters (width = 1).
	_ = tabwriter.StripEscape                  // Strip Escape characters bracketing escaped text segments instead of passing them through unchanged with the text.
	_ = tabwriter.AlignRight                   // Force right-alignment of cell content. Default is left-alignment.
	_ = tabwriter.DiscardEmptyColumns          // Handle empty columns as if they were not present in the input in the first place.
	_ = tabwriter.TabIndent                    // Always use tabs for indentation columns (i.e., padding of leading empty cells on the left) independent of padchar.
	_ = tabwriter.Debug                        // Print a vertical bar ('|') between columns (after formatting). Discarded columns appear as zero-width columns ("||").
	_ = tabwriter.AlignRight | tabwriter.Debug // Flags can be combined

	t := tabwriter.NewWriter(os.Stdout, 8, 0, 1, ' ', tabwriter.Debug)
	_ = t
	for _, p := range paths {
		d := dirs.New(p)
		_ = d
		println("")
		dirs.PrintS(d)
		_, _ = path.Split(p)
		// dir, fil := path.Split(p)
		// println("Dir:", dir, "\tFile:", fil)
	}

	// Output:
}
*/

func ExampleDirSplitter() {
	paths := []string{
		/*
			"a/",
			"a/c",
			"a//c",
			"a/c/.",
			"a/c/b/..",
			"/../a/c",
			"/../a/b/../././/c",
			"/a/b/c/bar.css/css",
			"/a/b/c/bar.css/.css",
			"/a/b/c/bar.css/",
			"/a/b/c/bar.css/.",
			"/a/b/c/bar.css/..",
			"/a/b/c/bar.css/../.",
			"/a/b/c/../bar.css",
			"/a/b/c/../bar.css/",
			"/a/b/c/../bar.css/.",
			"/a/b/c/../bar.css/..",
		*/
		"/a/b/c/../bar.css .../.",
		"/a/b/c/../bar.css / .../.",
		"/a/b/c..../../bar.css..../..",
		"/a/b/c/..../../bar.css/..../..",
		"/a/b/c../bar.css...../...",
		"/a/b/c/../bar.css/...../...",
		"/a/b/c/.../bar.css....../....",
		"/a/b/c/...../bar.css/....../....",
	}
	for _, p := range paths {
		l := dirs.PathSplitter(p)
		fmt.Println("Path:", p, " => ", l)
	}
}

func ExampleDotPath() {
	paths := []string{
		"/a/b/c/../bar.css .../.",
		"/a/b/c/../bar.css / .../.",
		"/a/b/c/../bar.css..../..",
		"/a/b/c/../bar.css/..../..",
		"/a/b/c/../bar.css...../...",
		"/a/b/c/../bar.css/...../...",
		"/a/b/c/../bar.css....../....",
		"/a/b/c/../bar.css/....../....",
	}

	for _, p := range paths {
		dp := dirs.NewDotPath(p)
		fmt.Println("====================================")
		fmt.Println("List:", p, " => ", dp.String())
		fmt.Println("Path:", p, " => ", dp.Path())
		fmt.Println("Look:", p, " => ", dp.RecursePathS())
		fmt.Println("Exec:", p, " => ", dp.InspectPathS())
	}
}
func main() {
	ExampleDotPath()
}
