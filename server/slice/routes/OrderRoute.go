package routes

import (
	"net/http"
	"sliceServer/controllers"
)

func OrderRouter() {
	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// controllers.GetAllOrders(w, r)
		case http.MethodPost:
			controllers.CreateOrder(w,r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// http.HandleFunc("/orders/", func(w http.ResponseWriter, r *http.Request) {
	// 	switch r.Method {
	// 	case http.MethodGet:
	// 		controllers.GetOrderByID(w, r)
	// 	case http.MethodPut:
	// 		controllers.UpdateOrderByID(w, r)
	// 	case http.MethodDelete:
	// 		controllers.DeleteOrderByID(w, r)
	// 	default:
	// 		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	// 	}
	// })
}
