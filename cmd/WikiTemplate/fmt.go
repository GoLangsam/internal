// Copyright 2016 Andreas Pannewitz. All rights reserved.

package WikiTemplate

import "fmt"

// Println for easy access
func Println(a ...interface{}) (n int, err error) {
	return fmt.Println(a)
}
