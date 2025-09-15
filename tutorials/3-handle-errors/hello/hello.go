// https://go.dev/doc/tutorial/create-module

package main

import (
	"fmt"
	"log"
	"example.com/greetings"
)


func main () {

	log.SetFlags(0)
	log.SetPrefix("greetings:")

	message, err := greetings.Hello("")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}