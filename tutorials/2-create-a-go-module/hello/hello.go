package main

import (
	"fmt"

	"example.com/greetings"
)


func main () {
	message:= greetings.Hello("Gdays")
	fmt.Println(message)
}