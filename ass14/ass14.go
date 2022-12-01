package main

import "fmt"

func main() {
	var a [10][10]int
	var b [10][10]int
	var c [10][10]int
	var limit int
	fmt.Println("Enter the size of array")
	fmt.Scan(&limit)
	fmt.Println("Enter the values of array1")
	for i := 0; i < limit; i++ {
		for j := 0; j < limit; j++ {
			fmt.Scan(&a[i][j])
		}
	}
	fmt.Println("Enter the values of array2")
	for i := 0; i < limit; i++ {
		for j := 0; j < limit; j++ {
			fmt.Scan(&b[i][j])
		}
	}

	fmt.Println("Sum of 2 arrays is:")
	for i := 0; i < limit; i++ {
		for j := 0; j < limit; j++ {
			c[i][j] = a[i][j] + b[i][j]
			fmt.Print(b[i][j], " ")
		}
		fmt.Print("\n")
	}
}
