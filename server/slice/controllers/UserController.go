package controllers

import (
	"encoding/json"

	"net/http"
	"sliceServer/models"
	"sliceServer/services"
	"strconv"
	"strings"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	_ = services.LoadUsers()
	json.NewEncoder(w).Encode(services.Users)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	endPoints := strings.Split(r.URL.Path, "/")
	id, _ := strconv.Atoi(endPoints[2])
	for _, user := range services.Users {
		if id == user.ID {
			json.NewEncoder(w).Encode(user)
			return
		}
	}
}
func CreateNewUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	json.NewDecoder(r.Body).Decode(&newUser)
	newUser.ID = len(services.Users) + 1
	services.Users = append(services.Users, newUser)
	json.NewEncoder(w).Encode(services.Users)
}
func UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	endPoints := strings.Split(r.URL.Path, "/")
	id, _ := strconv.Atoi(endPoints[2])
	var updateUser models.User
	json.NewDecoder(r.Body).Decode(&updateUser)

	for i, user := range services.Users {
		if user.ID == id {
			updateUser.ID = user.ID
			services.Users[i] = updateUser
			json.NewEncoder(w).Encode(user)
			return
		}
	}

}
func DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	endPoints := strings.Split(r.URL.Path, "/")
	id, _ := strconv.Atoi(endPoints[2])

	for i, users := range services.Users {
		if users.ID == id {
			services.Users = append(services.Users[:i], services.Users[i+1:]...)
			json.NewEncoder(w).Encode(users)
			return
		}
	}
}
