package main

import "fmt"

func main() {
	var p int
	var r, n, si float32

	fmt.Println("Enter the principle amount")
	fmt.Scanln(&p)
	fmt.Println("Enter the Interest rate")
	fmt.Scanln(&r)
	fmt.Println("Enter the number of years ")
	fmt.Scanln(&n)
	si = ((float32(p) * r * n) / 100)
	fmt.Println("Your Simple Interest is", si)
}
