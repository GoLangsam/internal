// +build ignore

package main

import "fmt"

func dup3(in <-chan int) (<-chan int, <-chan int, <-chan int) {

	a := make(chan int, 2)
	b := make(chan int, 2)
	c := make(chan int, 2)

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

func fib() <-chan int {
	x := make(chan int, 2)
	a, b, out := dup3(x)
	go func() {
		x <- 0
		x <- 1
		<-a
		for {
			x <- <-a + <-b
		}
	}()
	return out
}

func main() {
	x := fib()
	for i := 0; i < 10; i++ {
		fmt.Println(<-x)
	}
}

// I like it because I only ever declare a single integer variable, and the rest goes through chans.
// This solution is inspired by Haskell's concise "fib = 0:zipWith (+) fib (1:fib)".
