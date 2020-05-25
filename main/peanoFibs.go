// +build ignore

// Peano integers are represented by a linked
// list whose nodes contain no data
// (the nodes are the data).
// http://en.wikipedia.org/wiki/Peano_axioms

// This program demonstrates that Go's automatic
// stack management can handle heavily recursive
// computations.
package main

import "fmt"

// Number is a pointer to a Number
type Number *Number

// The arithmetic value of a Number is the
// count of the nodes comprising the list.
// (See the count function below.)

// -------------------------------------
// Peano primitives

func zero() *Number {
	return nil
}

func isZero(x *Number) bool {
	return x == nil
}

func add1(x *Number) *Number {
	e := new(Number)
	*e = x
	return e
}

func sub1(x *Number) *Number {
	return *x
}

func add(x, y *Number) *Number {
	if isZero(y) {
		return x
	}
	return add(add1(x), sub1(y))
}

func mul(x, y *Number) *Number {
	if isZero(x) || isZero(y) {
		return zero()
	}
	return add(mul(x, sub1(y)), x)
}

func fact(n *Number) *Number {
	if isZero(n) {
		return add1(zero())
	}
	return mul(fact(sub1(n)), n)
}

// -------------------------------------
// Helpers to generate/count Peano integers

func gen(n int) *Number {
	if n > 0 {
		return add1(gen(n - 1))
	}
	return zero()
}

func count(x *Number) int {
	if isZero(x) {
		return 0
	}
	return count(sub1(x)) + 1
}

// -------------------------------------
// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// ===========================================================================

// PeanoFact returns the factorial N! of N.
//
// For N > 10 zero is returned as memory would overflow on 32bit architectures
// when using larger values.
func PeanoFact(N int) int {
	if N > 10 {
		N = 0
	} // Max = 10 on 32bit
	return count(fact(gen(N)))
}

// ===========================================================================
// Iterate successive Fibonacci numbers

// peanoFib represents an iterator
type peanoFib struct {
	prev, next *Number
	iter       int
}

// PeanoFibs returns an iterator for successive Fibonacci numbers
// based on Peano numbers.
func PeanoFibs() *peanoFib {
	return &peanoFib{prev: zero(), next: add1(zero())}
}

// Next advances the iterator.
//
// For N > 34 zero is returned as memory would overflow on 32bit architectures
// when using larger values.
func (a *peanoFib) Next() (next bool) {
	if a.iter < 34 { // Max = 34 on 32bit
		a.prev, a.next = a.next, add(a.prev, a.next)
		a.iter++
		return true
	}
	a.prev = zero()
	return
}

// Fib returns the Fibonacci number of the iterator.
func (a *peanoFib) Fib() (fib int) {
	fib = count(a.prev)
	return
}

// ===========================================================================

// showPeanoFact shows how to use the iterator.
func showPeanoFact() {
	fmt.Println("PeanoFact [0..12]")
	for i := 0; i <= 12; i++ {
		fmt.Println(PeanoFact(i))
	}
	fmt.Println("Note: For N > 10 memory would overflow on 32bit architectures")
}

// ===========================================================================

// showPeanoFibs shows how to use the iterator.
func showPeanoFibs() {
	fmt.Println("PeanoFibs [0..33]")
	p := PeanoFibs()
	for p.Next() {
		fmt.Println(p.Fib())
	}
	fmt.Println("Note: For N > 33 memory would overflow on 32bit architectures")
}

// ===========================================================================

func main() {
	showPeanoFact()
	fmt.Println()
	showPeanoFibs()
}
