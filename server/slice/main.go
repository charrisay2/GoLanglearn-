package main

import (
	"fmt"
	"net/http"
	"sliceServer/routes"
	"sliceServer/services"
)

func main() {
	routes.UserRoute()
	services.LoadUsers()
	fmt.Println("server is running at: http://localhost:9999/")
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		panic(err)
	}

}
