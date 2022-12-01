package main

import "fmt"

func main() {
	var arr1 [10]int
	var arr2 [10]int
	var limit int
	fmt.Println("Enter the limit of array")
	fmt.Scanln(&limit)
	fmt.Println("Enter the values of array1")
	for i := 0; i < limit; i++ {
		fmt.Scan(&arr1[i])
	}
	fmt.Println("Enter the values of array2")
	for i := 0; i < limit; i++ {
		fmt.Scan(&arr2[i])
	}
	for i := 0; i < limit; i++ {
		arr1[i] = arr1[i] + arr2[i]
		arr2[i] = arr1[i] - arr2[i]
		arr1[i] = arr1[i] - arr2[i]
	}
	fmt.Print("Arrays after swapping\nArray 1   ")
	for i := 0; i < limit; i++ {
		fmt.Print(arr1[i], " ,")
	}
	fmt.Print("\nArray 2   ")
	for i := 0; i < limit; i++ {
		fmt.Print(arr2[i], " ,")
	}
}
