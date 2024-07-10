package main

func main() {
	// // Declaring a pointer
	// var p *int

	// // Declaring a variable
	// var a int = 42

	// // Assigning the address of a to p
	// p = &a

	// // Printing the value of a
	// println(a)

	// // Printing the address of a
	// println(p)

	// // Printing the value of a using the pointer p
	// println(*p)

	// // Changing the value of a using the pointer p
	// *p = 21

	// // Printing the value of a
	// println(a)

	// // Printing the value of a using the pointer p
	// println(*p)

	// var a = 10
	// var b int

	// var p *int

	// p = &a
	// b = a

	// *p = 20

	// println(a, b)

	var a = 10
	println(a)
	add12(a)
	println(a)

	add12p(&a)
	println(a)

}

func add12(a int) {
	a += 12
}

func add12p(a *int) {
	*a += 12
}
