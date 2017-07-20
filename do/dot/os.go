// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dot

import (
	"os"
)

// OsFriendly - interface exported for go doc only
type OsFriendly interface {
	FileInfoIsDir(d Dot) bool
	FileInfoIsFile(d Dot) bool
	IsFileInfo(d Dot) bool
}

// var _ OsFriendly = New("Interface satisfied? :-)")

// IsFileInfo determines, if we have an `os.FileInfo`
func IsFileInfo(v interface{}) bool {
	switch v.(type) {
	case os.FileInfo:
		return true
	default:
		return false
	}
}

// ToFileInfo casts to `os.FileInfo`, if possible
func ToFileInfo(v interface{}) (os.FileInfo, bool) {
	switch v := v.(type) {
	case os.FileInfo:
		return v, true
	default:
		return nil, false
	}
}

// FileInfoIsDir casts to `os.FileInfo`, if possible, and determines IsDir()
func FileInfoIsDir(d Dot) bool {
	fi, ok := ToFileInfo(d.GetV())
	switch {
	case !ok:
		return false
	default:
		return fi.IsDir()
	}
}

// FileInfoIsFile casts to `os.FileInfo`, if possible, and determines !IsDir()
func FileInfoIsFile(d Dot) bool {
	fi, ok := ToFileInfo(d.GetV())
	switch {
	case !ok:
		return false
	default:
		return !fi.IsDir()
	}
}

/*
go doc os | grep string
found:

var Args []string					=> template.TempFunc "cmdline"

func Getwd() (dir string, err error)			=> template.TempFunc
func Chdir(dir string) error
	func Chmod(name string, mode FileMode) error
	func Chown(name string, uid, gid int) error
	func Chtimes(name string, atime time.Time, mtime time.Time) error

func Environ() []string					=> template.TempFunc
func ExpandEnv(s string) string
func Getenv(key string) string
func Setenv(key, value string) error
func LookupEnv(key string) (string, bool)
func Unsetenv(key string) error

func Expand(s string, mapping func(string) string) string
	Expand replaces ${var} or $var in the string based on the mapping function.
	For example, os.ExpandEnv(s) is equivalent to os.Expand(s, os.Getenv).

func Hostname() (name string, err error)		=> template.TempFunc
	func Lchown(name string, uid, gid int) error

func Mkdir(name string, perm FileMode) error
func MkdirAll(path string, perm FileMode) error
	func NewSyscallError(syscall string, err error) error

func Remove(name string) error
func RemoveAll(path string) error

func Rename(oldpath, newpath string) error

func Link(oldname, newname string) error
func Symlink(oldname, newname string) error
func Readlink(name string) (string, error)

func TempDir() string					=> template.TempFunc
	func Truncate(name string, size int64) error

*File
    func Create(name string) (*File, error)
    func NewFile(fd uintptr, name string) *File
    func Open(name string) (*File, error)
    func OpenFile(name string, flag int, perm FileMode) (*File, error)
    func Lstat(name string) (FileInfo, error)
    func Stat(name string) (FileInfo, error)
    func StartProcess(name string, argv []string, attr *ProcAttr) (*Process, error)
*/
