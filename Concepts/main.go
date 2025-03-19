package main

import "fmt"

func ifElse() {
	var companyName, companyName2 string
	var role, role2 string
	var salary, salary2 float64

	fmt.Println("Enter comapny name 1:")
	fmt.Scanln(&companyName)

	fmt.Println("Enter your role 1:")
	fmt.Scanln(&role)

	fmt.Println("Enter your salary 1:")
	fmt.Scanln(&salary)

	fmt.Println("Enter comapny name 2:")
	fmt.Scanln(&companyName2)

	fmt.Println("Enter your role 2:")
	fmt.Scanln(&role2)

	fmt.Println("Enter your salary 2:")
	fmt.Scanln(&salary2)

	if salary > salary2 {
		fmt.Println("Salary 1 is greater")
	} else if salary < salary2 {
		fmt.Println("Salary 2 is greater")
	} else {
		fmt.Println("both are equal")
	}
}
