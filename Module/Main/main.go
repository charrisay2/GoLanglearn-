package main

import (
	"fmt"
	"log"

	functions "Module/Functions"
)

func main() {

	names := []string{"chus cuoi", "chi hang", "tho ngoc"}

	messages, err := functions.NameGreetings(names)

	if err != nil {
		log.Fatal(err)
	}

	for _, message := range messages {
		fmt.Println(message)
	}
}