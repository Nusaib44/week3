package main

import "fmt"

func main() {
	var mark float32
	fmt.Println("enter your mark")
	fmt.Scanln(&mark)
	if mark < 50 {
		fmt.Printf("You Failed")
	} else {
		fmt.Println("You Passed")
	}
}
