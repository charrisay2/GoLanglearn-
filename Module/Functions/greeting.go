package functions

import (
	"errors"
	"fmt"
)

func Hello(name string) (string,error) {
	if name == "" {
		return "", errors.New("Name cannot be NULL.")
	}
	message := fmt.Sprintln(RandomFomat,name)

	// không có null chỉ có nill
	return message,nil
}



