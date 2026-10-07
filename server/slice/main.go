package main

import (
	"fmt"
	"net/http"
	configs "sliceServer/Configs"
	"sliceServer/routes"
)

func main() {
	configs.Connect()
	routes.UserRoute()
	fmt.Println("server is running at: http://localhost:9999/")
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		panic(err)
	}

}
