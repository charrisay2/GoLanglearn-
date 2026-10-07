package controllers

import (
	"encoding/json"

	"net/http"
	"sliceServer/models"
	"sliceServer/services"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(services.GetAllUsers())
}

//	func GetUserById(w http.ResponseWriter, r *http.Request) {
//		endPoints := strings.Split(r.URL.Path, "/")
//		id, _ := strconv.Atoi(endPoints[2])
//		for _, user := range services.Users {
//			if id == user.ID {
//				json.NewEncoder(w).Encode(user)
//				return
//			}
//		}
//	}
func CreateNewUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	json.NewDecoder(r.Body).Decode(&newUser)
	json.NewEncoder(w).Encode(services.Create(newUser))
}

// // func UpdateUserByID(w http.ResponseWriter, r *http.Request) {
// 	endPoints := strings.Split(r.URL.Path, "/")
// 	id, _ := strconv.Atoi(endPoints[2])
// 	var updateUser models.User
// 	json.NewDecoder(r.Body).Decode(&updateUser)

// 	for i, user := range services.Users {
// 		if user.ID == id {
// 			updateUser.ID = user.ID
// 			services.Users[i] = updateUser
// 			json.NewEncoder(w).Encode(user)
// 			return
// 		}
// 	}

// }
// func DeleteUserByID(w http.ResponseWriter, r *http.Request) {
// 	endPoints := strings.Split(r.URL.Path, "/")
// 	id, _ := strconv.Atoi(endPoints[2])

// 	for i, users := range services.Users {
// 		if users.ID == id {
// 			services.Users = append(services.Users[:i], services.Users[i+1:]...)
// 			json.NewEncoder(w).Encode(users)
// 			return
// 		}
// 	}
// }
