package main

import (
	"fmt"
)

func main() {
	// myChannel := make(chan int, 2)

	// myChannel <- 1
	// myChannel <- 2
	// myChannel <- 3 // deadlock

	// data := <-myChannel
	// fmt.Println("Data: ", data)
	// fmt.Println("End of main function")

	// myChannel <- 1
	// myChannel <- 2
	// fmt.Println(<-myChannel)

	// myChannel <- 3

	// fmt.Println(<-myChannel)
	// fmt.Println(<-myChannel)

	// fmt.Println(<-myChannel)

	// messages := make(chan string, 2)

	// go func() {
	// 	data1 := <-messages
	// 	fmt.Println("Data 1: ", data1)
	// 	data2 := <-messages
	// 	fmt.Println("Data 2: ", data2)
	// }()

	// messages <- "Hello"
	// messages <- "World"
	// messages <- "123123"

	// time.Sleep(1 * time.Second)
	// fmt.Println("End of main function")

	bufferedChannel := make(chan int, 4)

	go func() {
		for i := 1; i <= 10; i++ {
			bufferedChannel <- i
			fmt.Println("Sent data: ", i)
		}
		close(bufferedChannel)
	}()

	for data := range bufferedChannel {
		fmt.Println("Received data: ", data)
	}
}
