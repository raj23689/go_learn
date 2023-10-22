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

	// default values and aliases
	var minElement int
	fmt.Println(minElement)
	fmt.Printf("Variable is of type: %T \n", minElement)

	// implicit type
	var genre = "Anime"
	fmt.Println(genre)

	// no var style
	website := 451515.20
	fmt.Println(website)

}
