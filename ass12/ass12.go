package main

import "fmt"

func main() {
	var a [10]int
	var limit, temp int
	fmt.Print("Enter the size of array")
	fmt.Scan(&limit)
	fmt.Print("Enter the values of array")
	for i := 0; i < limit; i++ {
		fmt.Scan(&a[i])
	}
	for i := 0; i < limit; i++ {
		for j := 0; j < limit; j++ {
			if a[i] > a[j] {
				temp = a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	fmt.Print("Sorted array")
	for i := 0; i < limit; i++ {
		fmt.Print(a[i], " ")
	}
}
