package main

import "fmt"

func main() {
	myChannel := make(chan int)

	go func() {
		data := <-myChannel
		fmt.Println("First Go Routine data: ", data)
	}()

	go func() {
		data := <-myChannel
		fmt.Println("Second Go Routine data: ", data)
	}()

	myChannel <- 42
	fmt.Println("End of main function")
}
