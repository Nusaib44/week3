package main

import "fmt"

func main() {
	var x int
	var y, z float32
	fmt.Println("Enter two numbers")
	fmt.Scanln(&x)
	fmt.Scanln(&y)

	z = y + float32(x)
	fmt.Println("Sum of numbers are", z)

}
