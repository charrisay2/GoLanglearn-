package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w,"chu cuoi ngoi goc cay da")
}
func main() {
	http.HandleFunc("/",helloHandler)
	fmt.Println("server is running at: http://localhost:9999/")
	err := http.ListenAndServe(":9999",nil)
	if err != nil {
		panic(err)
	}



}