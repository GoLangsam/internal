// +build ignore

// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "fmt"

// ===========================================================================

type ic chan int

func (c ic) next() (v int, ok bool) {
	v, ok = <-c
	return
}

// iter1 looses one item
func (c ic) iter1(many int) ic {
	i := 0
	for v, ok := c.next(); ok && i < many; v, ok = c.next() {
		fmt.Println("#", i, "\t ", v)
		i++
	}
	return c
}

// iter2 looses one item
func (c ic) iter2(many int) ic {
	i := 0
	for v, ok := <-c; ok && i < many; v, ok = <-c {
		fmt.Println("#", i, "\t ", v)
		i++
	}
	return c
}

// iter3 does not loose any item
func (c ic) iter3(many int) ic {
	for i := 0; i < many; i++ {
		if v, ok := c.next(); ok {
			fmt.Println("#", i, "\t ", v)
		} else {
			break
		}
	}
	return c
}

func (c ic) drain() {
	for _ = range c {
	}
}

func ten() (c ic) {
	c = make(chan int)
	go func(c chan<- int) {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}(c)
	return c
}

func main() {

	ten().iter1(2).iter1(2).drain()
	ten().iter2(2).iter2(2).drain()
	ten().iter3(2).iter3(2).drain()
}

// ===========================================================================
