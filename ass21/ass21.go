package main

import "fmt"

func main() {
	var a [10]int
	var m [10]int
	var size int
	fmt.Print("Enter the size of Array ")
	fmt.Scanln(&size)
	fmt.Println("Enter the values of array")
	for i := 0; i < size; i++ {
		fmt.Scan(&a[i])
	}

	for i := 0; i < size; i++ {
		m[i] = a[i] * a[i+1]
	}

	for i := 0; i < size-1; i++ {
		fmt.Print(m[i], "  ")
	}
}
