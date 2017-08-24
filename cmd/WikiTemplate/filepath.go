// Copyright 2016 Andreas Pannewitz. All rights reserved.

package WikiTemplate

import (
	"path/filepath"
)

func fullname(pathname, basename string) string {
	return filepath.Join(pathname, basename)
}
func asCSV(basename string) string {
	return basename + myExtCSV
}

func asTpl(basename string) string {
	return basename + myExtTpl
}

func asWiki(basename string) string {
	return basename + myExtWiki
}
