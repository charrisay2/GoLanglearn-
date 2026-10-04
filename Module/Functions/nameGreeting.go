package functions

import "fmt"

func NameGreetings(names []string) ([]string, error) {

	messages := []string{}

	for _, name := range names {
		message := fmt.Sprintf("Hi, %s", name)
		messages = append(messages, message)
	}

	return messages, nil
}
