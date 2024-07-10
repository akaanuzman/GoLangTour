package main

func main() {
	// Defer is a keyword that allows you to defer the execution of a function until the surrounding function returns.
	// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
	// defer println("First Defer")
	// defer println("Second Defer")
	// defer println("Thirth Defer")

	// println("World")

	// defer println("First Defer")

	// var condition bool = true

	// if condition {
	// 	return
	// }

	// defer println("Second Defer") // Dead code

	// x := 10
	// y := 20

	// defer println("x:", x, "y:", y)

	// x = 100
	// y = 200

	// println("x:", x, "y:", y)

	var hasError bool = true

	defer cleanup()

	if hasError {
		panic("An error occurred")
	}
}

func cleanup() {
	println("Cleaning up worked...")
}
