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

	var smallFloat float32 = 255.4555444432654432
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)
}
