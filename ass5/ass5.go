package main

import "fmt"

func main() {
	var mark float32

	fmt.Println("Enter the mark")
	fmt.Scanln(&mark)
	if mark < 50 {
		fmt.Println(" You were failed")
	} else {
		fmt.Print("Your grade is ")
	}

	if mark > 49 && mark < 60 {
		fmt.Println("E")
	} else if mark > 59 && mark < 70 {
		fmt.Println("D")
	} else if mark > 69 && mark < 80 {
		fmt.Println("C")
	} else if mark > 79 && mark < 90 {
		fmt.Println("B")
	} else if mark > 89 {
		fmt.Println("A")
	}

}
