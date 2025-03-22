package main

import "fmt"

func main() {
	// for loop in Go
	// var nums int = 5
	// for nums > 0 {
	// 	fmt.Println(nums)
	// 	nums--
	// }

	var arr [3]string
	for i := 0; i < 3; i++ {
		fmt.Println("Enter coding language :")
		fmt.Scanln(&arr[i])
	}

	for i := 0; i < 3; i++ {
		fmt.Printf("Coding Languages: %s\n", arr[i])
	}

}
