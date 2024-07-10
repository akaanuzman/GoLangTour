package main

import "fmt"

func main() {
	var customer = Customer{id: 1, name: "John", age: 30, address: Address{city: "New York", district: "Manhattan"}}

	customer.toString()
	customer.changeName("Doe")
	customer.toString()
}

type Customer struct {
	id      int64
	name    string
	age     int
	address Address
}

type Address struct {
	city     string
	district string
}

func (customer *Customer) toString() {
	fmt.Printf("Customer{id: %d, name: %s, age: %d, address: %s\n}", customer.id, customer.name, customer.age, customer.address)
}

func (customer *Customer) changeName(name string) {
	customer.name = name
}
