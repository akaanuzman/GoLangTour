package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func printTypeAndValue(value any) {
	// %T is the type of the value
	// %v is the value in a default format
	fmt.Printf("Type: %T Value: %v\n", value, value)
}

func __main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)

	printTypeAndValue(ToBe)
	printTypeAndValue(MaxInt)
	printTypeAndValue(z)

	var ll int
	var f float64
	var b bool
	var s string
	// %q is a single-quoted character literal safely escaped with Go syntax.
	fmt.Printf("%v %v %v %q\n", ll, f, b, s)
}
