package main

import (
	"fmt"

	"github.com/Samanthapuri/gopractice/greetings"
)

var name string

func main() {

	name = "Vijay"

	message := greetings.Hello(name)
	fmt.Println(message)

}
