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

	    // A slice of names.
    names := []string{"Gladys", "Samantha", "Darrin"}

    // Request greeting messages for the names.
    messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}