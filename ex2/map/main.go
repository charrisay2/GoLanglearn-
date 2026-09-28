package main

func main() {
	menu := make(map[int]Coffee)

	addCoffee(menu, 1, Coffee{"CaPhe", 25000})
	addCoffee(menu, 2, Coffee{"BacXiu", 30000})
	addCoffee(menu, 3, Coffee{"TraDao", 25000})

	displayMenu(menu)

	editCoffee(menu, 2, "BacXiuSua", 35000)

	deleteCoffee(menu, 1)

	displayMenu(menu)
}