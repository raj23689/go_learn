package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza:")

	// Comma, ok || err ok
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for your input: ", input)
	fmt.Printf("Type of this input is %T: ", input)
}
