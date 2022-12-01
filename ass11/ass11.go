package main

import "fmt"

func main() {
	var arr1 [10]int
	var limit int
	fmt.Println("Enter the size of array")
	fmt.Scan(&limit)
	fmt.Println("Enter the values of array")
	count := 0
	for i := 0; i < limit; i++ {
		fmt.Scan(&arr1[i])
		if arr1[i]%2 == 0 {
			count++
		}
	}
	fmt.Println("Number of even numbers in the given array is :", count)
}
