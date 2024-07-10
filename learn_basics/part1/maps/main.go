package main

import "fmt"

func main() {
	// var names map[string]int = make(map[string]int, 0)

	// names["John"] = 41
	// names["Paul"] = 39
	// names["George"] = 40
	// names["Ringo"] = 35

	// println(names["John"])
	// println(names["Bill"]) // if key does not exist in map go return the zero value of the value type

	names := map[string]int{
		"John":   41,
		"Paul":   39,
		"George": 40,
		"Ringo":  35,
	}

	delete(names, "John")

	// For each loop
	for key, value := range names {
		println(key, value)
	}

	var language = "GoLang"

	for _, char := range language {
		fmt.Printf("%c\n", char)
	}
}
