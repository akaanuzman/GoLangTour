package main

import "fmt"

func main() {
	// var names [3]string
	// names[0] = "Ahmet"
	// names[1] = "Mehmet"
	// names[2] = "Veli"
	// names[3] = "Ali" // index out of range

	// names := [4]string{"Ahmet", "Mehmet", "Veli", "Ali"}

	// fmt.Println(names[0:2])

	var names = []string{"Ahmet", "Mehmet", "Veli", "Ali"}

	names = append(names, "Ayşe")
	fmt.Println(names)
}
