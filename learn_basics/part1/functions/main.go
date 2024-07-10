package main

func main() {
	// println(add(1, 2))
	// println(subtract(5, 3))
	// println(multiply(2, 3))
	// println(divide(10, 2))
	// // println(calculation(10, 5))
	// total, diff, multiply, divide := calculation(10, 5)
	// println(total, diff, multiply, divide)

	println(sum(1, 2, 3, 4, 5))
	println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}

// func add(a int, b int) int {
// 	return a + b
// }

// func subtract(a int, b int) int {
// 	return a - b
// }

// func multiply(a int, b int) int {
// 	return a * b
// }

// func divide(a int, b int) int {
// 	return a / b
// }

// func calculation(a int, b int) (int, int, int, int) {
// 	return add(a, b), subtract(a, b), multiply(a, b), divide(a, b)
// }

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
