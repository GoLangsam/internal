// +build ignore

// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// ref: https://blog.afoolishmanifesto.com/posts/buffered-channels-in-golang/

// Package main shows two funny ways to generate successive Fibonacci numbers
// using buffered channels.
//
// And it shows how quickly int32 silently overflows.
package main

import (
	"fmt"
	"math/big"
)

// ===========================================================================

func dup3(in <-chan int32) (<-chan int32, <-chan int32, <-chan int32) {

	a := make(chan int32)
	b := make(chan int32, 1)
	c := make(chan int32)

	go func() {
		for {
			x := <-in
			a <- x
			b <- x
			c <- x
		}
	}()
	return a, b, c
}

// ===========================================================================

// IntFibs sends successive Fibonacci numbers
// (starting with zero)
// on the returned channel.
//
// This implementation uses a triple lock-step fan.out:
//
// The first channel (non-buffered)
// is initially advanced/drained in order to get one step
// ahead.
//
// The second channel is 1-buffered in order to balance this.
//
// The third channel (non-buffered) is returned as it emits the results.
func IntFibs() <-chan int32 {
	x := make(chan int32)
	a, b, out := dup3(x)
	go func() {
		x <- 0
		<-a
		x <- 1
		for {
			x <- <-a + <-b
		}
	}()
	return out
}

// I like it because I only ever declare a single integer variable, and the rest goes through chans.
// This solution is inspired by Haskell's concise "fib = 0:zipWith (+) fib (1:fib)".

// Note: above comment is from the original author.

// ===========================================================================

func bigFibs(out chan<- *big.Int) {

	c := make(chan *big.Int, 2)
	c <- big.NewInt(0)
	c <- big.NewInt(1)

	for {
		n := <-c

		c <- big.NewInt(0).Add(n, <-c)
		c <- n

		out <- n
	}
}

// BigFibs sends successive Fibonacci numbers
// (starting with zero)
// on the returned channel.
//
// This implementation uses
// one 2-buffered channel
// to rotate the current and previous Fibonacci numbers.
func BigFibs() <-chan *big.Int {
	cha := make(chan *big.Int)
	go bigFibs(cha)
	return cha
}

// ===========================================================================

func main() {
	i := IntFibs()
	b := BigFibs()

	tab, eol := "\t", "\n"
	fmt.Println("# n", tab, "int", tab, "*big.Int", eol)
	for n := 0; n < 50; n++ { //
		fmt.Println("#", n, tab, <-i, tab, <-b)
	}
	fmt.Println(eol, "Note: on # n > 46 overflow occurs for int32 Fibs")
	fmt.Println(eol, "Note: on # n > 33 overflow occurs for peanoFibs")
}
