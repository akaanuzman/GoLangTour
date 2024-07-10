package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	var data1 string
	var data2 string

	go func() {
		time.Sleep(10 * time.Second)
		ch1 <- "Hello"
	}()

	go func() {
		time.Sleep(5 * time.Second)
		ch2 <- "World"
	}()

	for len(data1) == 0 || len(data2) == 0 {
		select {
		case data1 = <-ch1:
			fmt.Println("Data 1: ", data1)
		case data2 = <-ch2:
			fmt.Println("Data 2: ", data2)
		default:
			fmt.Println("No data came")
		}

		time.Sleep(2 * time.Second)

		// if data1 != "" && data2 != "" {
		// 	break
		// }
	}
}
