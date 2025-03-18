package main

import (
	"fmt"
)

// by default false
var status bool

func main() {
	// Explicite declaration
	const name string = "Preyansh"
	var age int = 23
	status1 := true

	fmt.Println("Name: ", name)
	fmt.Println("Age: ", age)

	name2 := "Khanna"
	age2 := 23

	fmt.Println("Name2:", name2)
	fmt.Println("Age2:", age2)
	fmt.Println("Status:", status)
	fmt.Println("Status1:", status1)
}
