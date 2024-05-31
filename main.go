package main

import (
	"fmt"

	"github.com/YamanaShop/welpen"
)

func main() {
	string1 := welpen.Bark()
	string2 := welpen.Barks()
	fmt.Println(string1)
	fmt.Println(string2)
	fmt.Println(welpen.Bark())
	fmt.Println(welpen.Barks())
	fmt.Println(welpen.BigBark())
	fmt.Println(welpen.BigBarks())
	fmt.Println(welpen.BigBarks())
	fmt.Println(welpen.From12())
}
