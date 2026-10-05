package services

import (
	"encoding/json"
	"os"
	"sliceServer/models"
)

var Users []models.User

const filePath = "Datas/UserData.json"

func LoadUsers() error {
	file, _ := os.Open(filePath)
	defer file.Close()
	return json.NewDecoder(file).Decode(&Users)
}
func SaveUsers() error {
	file, _ := os.Create(filePath)
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	return encoder.Encode(Users)
}
