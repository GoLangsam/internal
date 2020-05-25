// +build ignore

package main

import (
	"fmt"
	// "math"
	"strconv"
	"strings"
)

var input = "3 4 2 * 1 5 - 2 3 ^ ^ / +"

func pow(a, n int) (p int) {
	p = 1
	for i := 1; i < n; i++ {
		p *= a
	}
	return
}

func main() {
	fmt.Printf("For postfix %q\n", input)
	fmt.Println("\nToken            Action            Stack")
	var stack []int
	for _, tok := range strings.Fields(input) {
		action := "Apply op to top of stack"
		TOP, NIP := len(stack)-1, len(stack)-2
		switch tok {
		case "+":
			stack[NIP] += stack[TOP]
			stack = stack[:TOP]
		case "-":
			stack[NIP] -= stack[TOP]
			stack = stack[:TOP]
		case "*":
			stack[NIP] *= stack[TOP]
			stack = stack[:TOP]
		case "/":
			stack[NIP] /= stack[TOP]
			stack = stack[:TOP]
		case "^":
			stack[NIP] = pow(stack[NIP], stack[TOP])
			stack = stack[:TOP]
		default:
			action = "Push num onto top of stack"
			f, _ := strconv.Atoi(tok)
			stack = append(stack, f)
		}
		fmt.Printf("%3s    %-26s  %v\n", tok, action, stack)
	}
	fmt.Println("\nThe final value is", stack[0])
}
