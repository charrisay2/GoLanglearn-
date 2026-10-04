package routes

import (
	"net/http"
	"sliceServer/controllers"
)

func UserRoute() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			controllers.GetAllUsers(w,r)
			return
		}
		controllers.CreateNewUser(w,r)
	})
}
