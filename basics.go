package main

import (
	"fmt"
	"math"
	"math/rand"
)

// import "fmt"
// import "math/rand"

func add(x int, y int) int {
	return x + y
}

// Other way to declare the function
// func add(x, y int) int {
// 	return x + y
// }

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, java bool
var i, j int = 1, 2

func _main() {
	fmt.Println("Hello, World!")
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println(math.Pi)
	fmt.Println(add(42, 13))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(17))
	fmt.Println(c, python, java)
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
