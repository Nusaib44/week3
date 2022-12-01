package main

import "fmt"

func main() {

	var word string
	var length int
	flag := 0
	fmt.Print("Enter a Word : ")
	fmt.Scan(&word)
	length = len(word)

	for i := 0; i < length; i++ {
		if word[i] != word[length-i-1] {
			flag = 1
			break
		}
	}
	if flag == 0 {
		fmt.Println("Entered word is Palindrome")
	} else {
		fmt.Println("Entered word is not Palindrome")
	}
}
