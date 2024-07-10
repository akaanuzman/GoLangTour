package main

func main() {

	// index := 0

	// for index < 5 {
	// 	println(index)
	// 	index++
	// }

	// for i := 0; i < 5; i++ {
	// 	println(i)
	// }

	// index := 0
	// var names = []string{"John", "Paul", "George", "Ringo"}

	// for index < len(names) {
	// 	println(names[index])
	// 	index++
	// }

	for i := 0; i < 11; i++ {
		if i == 5 {
			break
		}
		if i == 3 {
			continue
		}
		println(i)
	}
}
