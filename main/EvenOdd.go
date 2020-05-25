// +build ignore

package main

import "fmt"

func main() {

	// i%2 == 0 versus i&1 == 0
	for i := 0; i < 24; i++ {
		if i&1 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}
}
