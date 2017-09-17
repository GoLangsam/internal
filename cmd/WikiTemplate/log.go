// Copyright 2016 Andreas Pannewitz. All rights reserved.

package WikiTemplate

import (
	"log"
)

func checkerr(err error, txt string) {
	if err != nil {
		log.Fatal(txt+": ", err)
	}
}
