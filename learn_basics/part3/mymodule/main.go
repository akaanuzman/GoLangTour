package main

import (
	"mymodule/helper"
	"mymodule/helper/rest"
)

func main() {
	helper.Helper()
	rest.Client2()
	// rest.privateFunc() // This will not work
}
