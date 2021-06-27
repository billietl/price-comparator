package main

import (
	"fmt"
	"gitlab.com/billietl/price-comparator/store"
)

func main() {
	fmt.Println("Hello world !")
	store.Hello()
	store.Load("foo")
}
