package main

type IShippable interface {
	shippingCost() int
}

type Book struct {
	desi int
}

func (b *Book) shippingCost() int {
	return 5 + b.desi*2
}

type Electronic struct {
	desi int
}

func (b *Electronic) shippingCost() int {
	return 10 + b.desi*4
}

type Flower struct {
	desi int
}

func (b *Flower) shippingCost() int {
	return 3 + b.desi*6
}

func main() {
	// var book = Book{desi: 5}
	// println(book.shippingCost())

	// books := []Book{
	// 	{desi: 5},
	// 	{desi: 10},
	// }
	// println(calculateTotalShippingCost(books))

	var products []IShippable = []IShippable{
		&Book{desi: 5},
		&Electronic{desi: 10},
		&Flower{desi: 15},
	}

	println(calculateTotalShippingCost(products))
}

func calculateTotalShippingCost(products []IShippable) int {
	totalCost := 0
	for _, product := range products {
		totalCost += product.shippingCost()
	}
	return totalCost
}
