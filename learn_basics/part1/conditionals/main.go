package main

func main() {

	// var isTrue bool = true

	// if isTrue {
	// 	println("isTrue is true")
	// } else {
	// 	println("isTrue is false")
	// }

	var a, b, c int = 10, 20, 30

	if a >= b && a >= c {
		println("a is the largest")
	} else if b >= a && b >= c {
		println("b is the largest")
	} else {
		println("c is the largest")
	}
}
