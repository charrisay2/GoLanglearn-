package controllers

import (
	"encoding/json"

	"net/http"
	"sliceServer/models"
	"sliceServer/services"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var o models.Order
	if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	if !services.UserExistByID(o.UserID) {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	id, err := services.CreateOrder(o)
	if err != nil {
		http.Error(w, "cannot create order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]any{
		"message":  "order created successfully",
		"order_id": id,
	})
}
