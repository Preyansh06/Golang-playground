package main

import (
	"fmt"
)

func main() {
	var companyName string
	var role string
	var salary float64

	fmt.Println("Enter comapny name:")
	fmt.Scanln(&companyName)

	fmt.Println("Enter your role:")
	fmt.Scanln(&role)

	fmt.Println("Enter your salary:")
	fmt.Scanln(&salary)

	println("Company name:", companyName)
	println("Role:", role)
	fmt.Printf("Salary: %.2f\n", salary)

}
