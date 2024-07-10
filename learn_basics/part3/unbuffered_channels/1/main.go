package main

import "fmt"

func main() {
	myChannel := make(chan string)
	done := make(chan bool)

	// go func() {
	// 	myChannel <- "Hello"
	// }()
	// // Blocking operation
	// message, isOpen := <-myChannel

	// fmt.Println(message, isOpen)

	go func() {
		message := <-myChannel
		fmt.Println(message)
		done <- true
	}()

	go func() {
		myChannel <- "Hello World"
	}()

	<-done
	fmt.Println("End of main function")
}
