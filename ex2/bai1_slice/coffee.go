package bai1slice

import "fmt"

type Coffee struct {
	ID    int
	Name  string
	Price float64
}

func addCoffee(menu []Coffee, coffee Coffee) []Coffee {
	menu = append(menu, coffee)
	return menu
}

func deleteCoffee(menu []Coffee, id int) []Coffee {
	for i, coffee := range menu {
		if coffee.ID == id {
			menu = append(menu[:i], menu[i+1:]...)
			fmt.Println("Đã xóa món!")
			return menu
		}
	}

	fmt.Println("Không tìm thấy món!")
	return menu
}

func editCoffee(menu []Coffee, id int, name string, price float64) []Coffee {
	for i := range menu {
		if menu[i].ID == id {
			menu[i].Name = name
			menu[i].Price = price
			fmt.Println("Đã sửa món!")
			return menu
		}
	}

	fmt.Println("Không tìm thấy món!")
	return menu
}

func displayMenu(menu []Coffee) {
	fmt.Println("\n===== MENU =====")

	if len(menu) == 0 {
		fmt.Println("Menu đang trống!")
		return
	}

	for _, coffee := range menu {
		fmt.Printf("ID: %d | Tên: %s | Giá: %.0f\n",
			coffee.ID,
			coffee.Name,
			coffee.Price,
		)
	}
}