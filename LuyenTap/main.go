package main

import "fmt"

type inter interface {
	~int
}

func test[T inter](x T) {
	fmt.Printf("%d",x)
}
func main() {
	type MyType int
	var x MyType = 100

	test(x)
}