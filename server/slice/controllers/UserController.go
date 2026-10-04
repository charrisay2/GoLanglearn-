package controllers

import (
	"encoding/json"
	"net/http"
	"sliceServer/models"
	"sliceServer/services"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(services.Users)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	
}
func CreateNewUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	json.NewDecoder(r.Body).Decode(&newUser)
	newUser.ID = len(services.Users) + 1
	services.Users = append(services.Users, newUser)
	json.NewEncoder(w).Encode(services.Users)
}
func UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	
}
func DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	
}
