// +build ignore

package main

import "fmt"
import "strings"

// https://FsharpForFunAndProfit.com/posts/roman-numeral-kata/

// ============================================================================

type Roman string

func tallyMarks(n int) Roman {
	var b strings.Builder
	for i := 0; i < n; i++ {
		b.WriteString("I")
	}
	return Roman(b.String())
}

func (s Roman) Replace(old, new string) Roman {
	return Roman(strings.Replace(string(s), old, new, -1))
}

func ToRoman(n int) Roman {
	return tallyMarks(n).
		// deci = 10 unit
		Replace("IIIIIIIIII", "X").
		Replace("IIIII", "V").
		Replace("VV", "X").
		// cent = 10 deci = 100 unit
		Replace("XXXXXXXXXX", "C").
		Replace("XXXXX", "L").
		Replace("LL", "C").
		// mille = 10 cent = 100 deci = 1000 unit
		Replace("CCCCCCCCCC", "M").
		Replace("CCCCC", "D").
		Replace("DD", "M").

		// optional substitutions
		Replace("IIII", "IV").
		Replace("VIV", "IX").
		Replace("XXXX", "XL").
		Replace("LXL", "XC").
		Replace("CCCC", "CD").
		Replace("DCD", "CM")
}

// ============================================================================

func biQuinaryDigits(place int, unit, five, ten string, arabic int) string {

	digit := arabic % (10 * place) / place
	switch digit {
	case 0:
		return ""
	case 1:
		return unit
	case 2:
		return unit + unit
	case 3:
		return unit + unit + unit
	case 4:
		return unit + five
	case 5:
		return five
	case 6:
		return five + unit
	case 7:
		return five + unit + unit
	case 8:
		return five + unit + unit + unit
	case 9:
		return unit + ten
	default:
		panic("single digit expected")
	}
}

func ToRomanII(n int) string {
	units := biQuinaryDigits(1, "I", "V", "X", n)
	tens := biQuinaryDigits(10, "X", "L", "C", n)
	hundreds := biQuinaryDigits(100, "C", "D", "M", n)
	thousands := biQuinaryDigits(1000, "M", "?", "!", n)

	return thousands + hundreds + tens + units
}

// ============================================================================

func main() {
	for _, n := range []int{946, 1956, 2022, 3497} {
		fmt.Println(n, "\t", ToRoman(n), "\t", ToRomanII(n))
	}
}
