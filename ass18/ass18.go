package main

import "fmt"

func main() {
	var w, l, a float32
	var g float32

	fmt.Println("Enter the written test mark")
	fmt.Scan(&w)
	fmt.Println("enter the lab exam mark")
	fmt.Scan(&l)
	fmt.Println("enter the Assignment mark")
	fmt.Scan(&a)
	g = (((w * 70) + (l * 20) + (a * 10)) / 100)
	fmt.Println("Your grade is : ", g)
}
