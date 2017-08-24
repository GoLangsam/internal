// Copyright 2016 Andreas Pannewitz. All rights reserved.

package WikiTemplate

import (
	"strings"
)

func unBracket(in, bo, bc string) string {
	return strings.TrimLeft(strings.TrimRight(in, bc), bo)
}

func split(s, d string) []string {
	return strings.Split(s, d)
}

func splitEmoji(s, d string) []string {
	var r []string
	for _, w := range strings.Split(s, d) {
		if w != "" {
			r = append(r, unBracket(w, "[", "]"))
		}
	}
	return r
}
