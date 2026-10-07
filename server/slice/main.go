package main

import (
	"fmt"
	"net/http"
	configs "sliceServer/Configs"
	"sliceServer/routes"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	configs.Connect()
	configs.ConnectMySQL()

	routes.OrderRouter()
	routes.UserRoute()

	fmt.Println("server is running at: http://localhost:9999/")
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		panic(err)
	}

}
