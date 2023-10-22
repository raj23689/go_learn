package main

import "fmt"

func main() {
	var userName string = "Raj"
	fmt.Println(userName)
	fmt.Printf("Variable is of type: %T \n", userName)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)
}
