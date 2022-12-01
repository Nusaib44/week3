package main

import "fmt"

func main() {
	var num, z int
	fmt.Println("enter a number")
	fmt.Scanln(&num)
	for i := 1; i <= 10; i++ {
		z = i * num
		fmt.Println(i, "x", num, "=", z)
	}
}
