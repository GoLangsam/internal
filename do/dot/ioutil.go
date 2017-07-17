// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dot

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var Perm os.FileMode = 0644 // default os.FileMode

// func ioutil.ReadFile(filename string) ([]byte, error)
func TagReadFile(tar Dot, filename string) Dot {
	myName := "TagReadFile"

	filename = filepath.FromSlash(filename)
	data, err := ioutil.ReadFile(filename)

	if !tar.SeeError(myName, filename, err) {
		tar.Tag(string(data))
	}
	return tar
}

func DoTagReadFile(src, tar Dot) Dot {
	TagReadFile(tar, src.String())
	return src
}

func ExecTagReadFile(d Dot) Dot {
	return TagReadFile(d, d.String())
}

// func ioutil.WriteFile(filename string, data []byte, perm os.FileMode) error
func WriteFileData(tar Dot, filename string, data []byte) Dot {
	myName := "WriteFileData"

	err := ioutil.WriteFile(filename, data, Perm)
	_ = tar.SeeError(myName, filename+": "+tar.String(), err)
	return tar
}

// func ioutil.WriteFile(filename string, data []byte, perm os.FileMode) error
func WriteFileFromValue(tar Dot, filename string) Dot {
	myName := "WriteFile"
	_ = myName
	if value, ok := (tar.GetV()).(string); ok {
		//	if value, ok := vNonEmpty(tar, myName); ok { // is V a string?
		WriteFileData(tar, frSlash(filename), []byte(value))
	}

	return tar
}

func DoWriteFileFromValue(src, tar Dot) Dot {
	WriteFileFromValue(tar, src.String())
	return src
}

func ExecWriteFileFromValue(d Dot) Dot {
	return DoWriteFileFromValue(d, d)
}

// readDir

func readDir(tar Dot, myName, dirname string) ([]os.FileInfo, bool) {
	dirname = frSlash(dirname)

	dirs, err := ioutil.ReadDir(dirname)
	if tar.SeeError(myName, dirname, err) {
		return []os.FileInfo{}, false
	} else {
		return dirs, true
	}
}

// ReadAllDir: Read all IsDir from dirname and Tag FileInfo, and recurse
// Skips ".git" or other dot nonsense.
// func ioutil.ReadDir(dirname string) ([]os.FileInfo, error)
func ReadAllDirs(tar Dot, dirname string) Dot {
	myName := "ReadDir"

	if dirs, ok := readDir(tar, myName, dirname); ok {
		for _, info := range dirs {
			tar.Tag(info)
			if info.IsDir() && !strings.HasPrefix(info.Name(), ".") { // No .git or other dot nonsense please.
				ReadAllDirs(lookupDot(tar, info.Name()), filepath.Join(dirname, info.Name()))
			}
		}
	}
	return tar
}

func DoReadAllDirs(src, tar Dot) Dot {
	dirname := src.String()
	ReadAllDirs(tar, dirname)
	return src
}

func ExecReadAllDirs(d Dot) Dot {
	return DoReadAllDirs(d, d)
}

// ReadDirFils: Read all !IsDir from dirname and Tag FileInfo
// func ioutil.ReadDir(dirname string) ([]os.FileInfo, error)
func ReadDirFils(tar Dot, dirname string) Dot {
	myName := "ReadDirFils"

	if dirs, ok := readDir(tar, myName, dirname); ok {
		for _, info := range dirs {
			if !info.IsDir() {
				c := lookupDot(tar, info.Name())
				c.Tag(info)
			}
		}
	}
	return tar
}

func DoReadDirFils(src, tar Dot) Dot {
	dirname := src.String()
	ReadDirFils(tar, dirname)
	return src
}

func ExecReadDirFils(d Dot) Dot {
	return DoReadDirFils(d, d)
}

// ReadDir: Read dirname and Tag FileInfo
// func ioutil.ReadDir(dirname string) ([]os.FileInfo, error)
func ReadDir(tar Dot, dirname string) Dot {
	myName := "ReadDir"

	if dirs, ok := readDir(tar, myName, dirname); ok {
		for _, info := range dirs {
			c := lookupDot(tar, info.Name())
			c.Tag(info)
		}
	}
	return tar
}

func DoReadDir(src, tar Dot) Dot {
	dirname := src.String()
	ReadDir(tar, dirname)
	return src
}

func ExecReadDir(d Dot) Dot {
	return DoReadDir(d, d)
}

/*
io/ioutil Discard
var Discard io.Writer = devNull(0)
    Discard is an io.Writer on which all Write calls succeed without doing
    anything.
*/

/*
func TempDir(dir, prefix string) (name string, err error)
func TempFile(dir, prefix string) (f *os.File, err error)
*/
