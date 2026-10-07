package services

import (
	configs "sliceServer/Configs"
	"sliceServer/models"

	_ "github.com/joho/godotenv/autoload"
)

func CreateOrder(order models.Order) (int64, error) {
	result, err := configs.MySQLDB.Exec(
		`INSERT INTO orders (user_id, product_name, quantity, price)
VALUES (?, ?, ?, ?)`,
		order.UserID, order.ProductName, order.Quantity, order.Price,
	)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}
