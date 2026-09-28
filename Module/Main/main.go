package main

import (
	"fmt"
)

func nextID() func() int  {
	value := 0
	return func() int {
		value++
		return value
	}
}

func multiplier(factor int) func(int)int  {
	
}
func main() {
	count := nextID()
	fmt.Println(count())

}