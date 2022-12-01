package main

import "fmt"

func main() {
	var limit int
	fmt.Println("Enter a limit")
	fmt.Scanln(&limit)
	sum := 0
	for i := 0; i <= limit; i++ {
		if i%2 == 1 {
			sum = sum + i
		}
	}
	fmt.Println("sum of odd number is", sum)
}
