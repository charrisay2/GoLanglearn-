package main

import "fmt"

type Coffee struct {
	Name  string
	Price float64
}

func addCoffee(menu map[int]Coffee, id int, coffee Coffee) {
	menu[id] = coffee
}

func deleteCoffee(menu map[int]Coffee, id int) {
	delete(menu, id)
}

func editCoffee(menu map[int]Coffee, id int, name string, price float64) {
	if _, exists := menu[id]; exists {
		menu[id] = Coffee{
			Name:  name,
			Price: price,
		}
	}
}

func displayMenu(menu map[int]Coffee) {

	for id, coffee := range menu {
		fmt.Printf("ID: %d | Tên: %s | Giá: %.0f\n",
			id,
			coffee.Name,
			coffee.Price,
		)
	}
}