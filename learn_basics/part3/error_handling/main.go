package main

import (
	"fmt"
)

func main() {
	// var x int

	// var y float32

	// var pointer *int

	// fmt.Println(x, y, pointer)

	// var pointer *int

	// if pointer == nil {
	// 	fmt.Println("pointer is nil")
	// } else {
	// 	fmt.Println("pointer is not nil")
	// }

	// fileData, err := os.ReadFile("sample.txt")

	// if err != nil {
	// 	fmt.Println("Error reading file", err)
	// 	return
	// }
	// fmt.Println(string(fileData))

	// result, err := divide(10, 0)
	// if err != nil {
	// 	fmt.Println("Error dividing", err)
	// 	return
	// }
	// fmt.Println(result)

	productService := ProductService{}

	product, err := productService.Add(Product{id: 0, name: "Product 1", price: 10})
	if err != nil {
		fmt.Println("Error adding product", err)
		return
	}
	fmt.Println(product)
}

// func divide(x int, y int) (int, error) {
// 	if y == 0 {
// 		return 0, errors.New("Cannot divide by zero")
// 	}
// 	return x / y, nil
// }

type Product struct {
	id    int
	name  string
	price int
}

type ValidationError struct {
	code    string
	message string
}

func (ve *ValidationError) Error() string {
	return fmt.Sprintf("Error code: %s, message: %s", ve.code, ve.message)
}

type ProductService struct{}

func (ps *ProductService) Add(product Product) (*Product, error) {
	if product.id == 0 {
		return nil, &ValidationError{code: "INVALID_ID", message: "Invalid product id"}
	}
	if product.name == "" {
		return nil, &ValidationError{code: "INVALID_NAME", message: "Invalid product name"}
	}
	if product.price == 0 {
		return nil, &ValidationError{code: "INVALID_PRICE", message: "Invalid product price"}
	}

	return &Product{id: product.id, name: product.name, price: product.price}, nil
}
