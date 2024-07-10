package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// wg := sync.WaitGroup{}
	// wg.Add(2)

	// go func() {
	// 	defer wg.Done()
	// 	println("Hello")
	// }()
	// go func() {
	// 	defer wg.Done()
	// 	println("World")
	// }()

	// wg.Wait()

	// println("End of main function")

	startTime := time.Now()

	wg := sync.WaitGroup{}

	wg.Add(3)

	go func() {
		defer wg.Done()
		fmt.Println("Start of the f1")
		time.Sleep(2 * time.Second)
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Start of the f2")
		time.Sleep(3 * time.Second)
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Start of the f3")
		time.Sleep(4 * time.Second)
	}()

	wg.Wait()

	fmt.Println("Time taken: ", time.Since(startTime))
}
