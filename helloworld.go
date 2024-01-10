package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Chris")
	fmt.Println(message)
}
