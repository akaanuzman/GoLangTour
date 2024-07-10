package main

import (
	"fmt"
)

func main() {

	// var productName string = "Laptop"
	// var quantity int = 10
	// var discount float32 = .5
	// var isInStock bool = true

	// productName = "Mobile"

	// fmt.Println(productName, reflect.TypeOf(productName))
	// fmt.Println(quantity, reflect.TypeOf(quantity))
	// fmt.Println(discount, reflect.TypeOf(discount))
	// fmt.Println(isInStock, reflect.TypeOf(isInStock))

	// productName := "Laptop"
	// fmt.Println(productName)

	// quantity := 10
	// fmt.Println(quantity, reflect.TypeOf(quantity))

	var productName string = "Laptop"
	var quantity int = 10
	var discount float32 = .5
	var isInStock bool = true

	// fmt.Println("Product Name:", productName, "Quantity:", quantity, "Discount:", discount, "Is In Stock:", isInStock)
	// default format için %v kullanılır..
	fmt.Printf("Product Name: %s Quantity: %d Discount: %f Is In Stock: %t\n", productName, quantity, discount, isInStock)
	fmt.Printf("Product Name: %v Quantity: %v Discount: %v Is In Stock: %v\n", productName, quantity, discount, isInStock)
}
