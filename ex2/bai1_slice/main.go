package bai1slice

func main() {
	menu := []Coffee{}

	menu = addCoffee(menu, Coffee{1, "CaPhe", 25000})
	menu = addCoffee(menu, Coffee{2, "BacXiu", 30000})
	menu = addCoffee(menu, Coffee{3, "TraDao", 25000})

	displayMenu(menu)

	menu = editCoffee(menu, 2, "BacXiuSua", 35000)

	menu = deleteCoffee(menu, 1)

	displayMenu(menu)
}