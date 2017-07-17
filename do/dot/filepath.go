// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dot

import (
	"os"
	"path/filepath"
	"strings"
)

func frSlash(path string) string {
	return filepath.Clean(filepath.FromSlash(path))
}

func toSlash(path string) string {
	return filepath.ToSlash(filepath.Clean(path))
}

// filePathGlob
func FilePathGlob(tar Dot, path string) Dot {
	myName := "Glob"

	m, err := filepath.Glob(frSlash(path))
	if tar.SeeError(myName, path, err) {
		return tar
	}

	var res []string
	for _, f := range m {
		res = append(res, toSlash(f))
	}

	tar.UnlockedAdd(FilePath.Id(), res...)
	return tar
}

// Glob with pattern from src and add results to tar below "FilePath:"
func DoFilePathGlob(src, tar Dot) Dot {
	FilePathGlob(tar, src.String())
	return src
}

// Glob with pattern from src and add results to tar below "FilePath:"
func ExecFilePathGlob(d Dot) Dot {
	return DoFilePathGlob(d, d)
}

/*
WalkFunc is the type of the function called for each file or directory
visited by Walk. The path argument contains the argument to Walk as a
prefix; that is, if Walk is called with "dir", which is a directory
containing the file "a", the walk function will be called with argument
"dir/a". The info argument is the os.FileInfo for the named filepath.

If there was a problem walking to the file or directory named by path, the
incoming error will describe the problem and the function can decide how to
handle that error (and Walk will not descend into that directory). If an
error is returned, processing stops. The sole exception is when the function

returns the special value SkipDir. If the function returns SkipDir when
invoked on a directory, Walk skips the directory's contents entirely. If the

function returns SkipDir when invoked on a non-directory file, Walk skips
the remaining files in the containing directory.
*/

// InfoWF returns a WalkFunc, which separates entries found into Dirs, Fils(if ext matches) and Path(other files)
func InfoWF(tar Dot, myName, ext string) filepath.WalkFunc {

	// define walkFunc
	var walkFunc filepath.WalkFunc = func(path string, info os.FileInfo, err error) error {
		// myName := myName
		// ext := ext
		// tar := tar
		if tar.SeeError(myName, path, err) {
			return nil
		}
		var d Dot
		if info.IsDir() {
			d = lookupDot(tar, Dirs.Id())
		} else if strings.HasSuffix(info.Name(), ext) {
			d = lookupDot(tar, Fils.Id())
		} else {
			d = lookupDot(tar, Path.Id())
		}
		lookupDot(d, toSlash(path)).Tag(info)

		return nil
	}
	return walkFunc
}

/*
Walk walks the file tree rooted at root, calling walkFn for each file or
directory in the tree, including root. All errors that arise visiting files
and directories are filtered by walkFn. The files are walked in lexical
order, which makes the output deterministic but means that for very large
directories Walk can be inefficient. Walk does not follow symbolic links.
*/

// Walk down path and add results: Key = FilePath, Val = os.FileInfo
func InfoWalk(tar Dot, dirname, ext string) Dot {
	myName := "InfoWalk"

	dirname = frSlash(dirname) // path to walk
	filepath.Walk(dirname, InfoWF(tar, myName, ext))

	return tar
}

func DoInfoWalk(src, tar Dot) Dot {
	dirname := src.String()
	if ext, ok := vNonEmpty(src, "DoInfoWalk"); ok {
		InfoWalk(tar, dirname, ext)
	}
	return src
}

func ExecInfoWalk(d Dot) Dot {
	return DoInfoWalk(d, d)
}

/*
filepath.Dir(x) (path string)             // path is x less Base(), and cleaned
filepath.Base(x) string                   // base is last element of x, and cleaned
filepath.Join(filepath.Dir(x), filepath.Base(x)) = filepath.Clean(x)
filepath.Ext(x)                           // includes the .dot
filepath.Split(x) (dir, file string)      // x = dir+file, dir may be empty
filepath.Clean(x) (path string)

filepath.SplitList(x) []string            // only filepath . good for PATH or GOPATH

*/
