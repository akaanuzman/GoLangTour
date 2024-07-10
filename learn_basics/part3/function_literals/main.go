package main

func main() {

	add := func(x int, y int) int {
		return x + y
	}

	multiply := func(x int, y int) int {
		return x * y
	}

	addResult, multiplyResult := calculator(10, 5, add, multiply)
	println(addResult, multiplyResult)
}

func calculator(x int, y int, operation func(int, int) int, otherOperation func(int, int) int) (int, int) {
	return operation(x, y), otherOperation(x, y)
}
