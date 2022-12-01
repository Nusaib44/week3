package main

import "fmt"

func main() {
	var num int
	flag := 0
	fmt.Println("Enter a number")
	fmt.Scanln(&num)
	if num == 1 || num == 0 {
		flag = 1
	} else {
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = 1
				break

			}
		}
	}
	if flag == 0 {
		fmt.Println("Number is prime")
	} else {
		fmt.Println("Number is not a prime")
	}

}
