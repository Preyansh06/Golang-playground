package main

import "fmt"

func main() {
	// for loop in Go
	// var nums int = 5
	// for nums > 0 {
	// 	fmt.Println(nums)
	// 	nums--
	// }

	// var arr [3]string
	// for i := 0; i < 3; i++ {
	// 	fmt.Println("Enter coding language :")
	// 	fmt.Scanln(&arr[i])
	// }

	// for i := 0; i < 3; i++ {
	// 	fmt.Printf("Coding Languages: %s\n", arr[i])
	// }

	var arr [3]int
	fmt.Println("Enter 3 numbers:")

	for i := 0; i < 3; i++ {
		fmt.Scanf("%d\n", &arr[i])
	}

	maxNum := arr[0]

	for i := 0; i < 3; i++ {
		if maxNum < arr[i] {
			maxNum = arr[i]
		}
	}

	fmt.Println("The highest number entered is: ", maxNum)

}
