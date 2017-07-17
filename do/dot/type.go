// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dot

// Types handled by this package
const (
	Error    Type = "Error" // error
	Dirs     Type = "Dirs"
	Fils     Type = "Fils"
	Path     Type = "Path"
	FileInfo Type = "FileInfo" // os.FileInfo
	FilePath Type = "FilePath" // path/filepath
	Meta     Type = "Meta"     // *template.Template - a parsed MetaTemplate
	Tmpl     Type = "Tmpl"     // *template.Template - a parsed Template(s)
	Exec     Type = "Exec"     // *template.Template - executed
)

var TypeS = []Type{
	Error,
	Dirs,
	Fils,
	Path,
	FileInfo,
	FilePath,
	Meta,
	Tmpl,
	Exec,
}

// Type represents a type handled by this package
type Type string

func (t Type) Id() string {
	return ":" + t.String() + ":"
}

func (t Type) String() string {
	return string(t)
}

func (t Type) Is(v interface{}) bool {
	switch t.String() {
	case "Error":
		return true
	case "Dirs":
		return true
	case "Fils":
		return true
	case "Path":
		return true
	case "FileInfo":
		return IsFileInfo(v)
	case "FilePath":
		return true
	case "Meta":
		return true
	case "Tmpl":
		return true
	case "Exec":
		return true
	default:
		return false
	}
}

func (t Type) In(d Dot) bool {
	switch t.String() {
	case "Error":
		return true
	case "Dirs":
		return true
	case "Fils":
		return true
	case "Path":
		return true
	case "FileInfo":
		return t.Is(d.GetV())
	case "FilePath":
		return true
	case "Meta":
		return t.Is(d.GetV())
	case "Tmpl":
		return t.Is(d.GetV())
	case "Exec":
		return true
	default:
		return false
	}
}

// if FileInfo.Is(d.GetV())
// if FileInfo.In(d)
// TODO: Implement inside *Dot with //  && d.Up().String == t.Id()
type Typer interface {
	Id() string
	String() string
	Is(v interface{}) bool
	In(d Dot) bool
}

var _ Typer = Error
var _ Typer = Dirs
var _ Typer = Fils
var _ Typer = Path
var _ Typer = FileInfo
var _ Typer = FilePath
var _ Typer = Meta
var _ Typer = Tmpl
var _ Typer = Exec
