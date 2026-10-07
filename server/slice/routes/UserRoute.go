package routes

import (
	"net/http"
	"sliceServer/controllers"
)

func UserRoute() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			controllers.GetAllUsers(w, r)
			return
		}
		controllers.CreateNewUser(w, r)
	})
	// http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
	// 	switch r.Method {
	// 	case "GET":
	// 		controllers.GetUserById(w, r)
	// 	case "PUT":
	// 		controllers.UpdateUserByID(w, r)
	// 	case "DELETE":
	// 		controllers.DeleteUserByID(w, r)
	// 	}

	// })
}
