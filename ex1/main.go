package main

import "fmt"

const MAX = 100

func addItem(ids *[MAX]int, names *[MAX]string, prices *[MAX]float64, count *int) {
	if *count >= MAX {
		fmt.Println("Menu đã đầy!")
		return
	}

	fmt.Print("Nhập mã món: ")
	fmt.Scan(&ids[*count])

	fmt.Print("Nhập tên món: ")
	fmt.Scan(&names[*count])

	fmt.Print("Nhập giá món: ")
	fmt.Scan(&prices[*count])

	*count++

	fmt.Println("Thêm món thành công!")
}

func displayItems(ids [MAX]int, names [MAX]string, prices [MAX]float64, count int) {
	if count == 0 {
		fmt.Println("Menu đang trống!")
		return
	}

	fmt.Println("\nMENU QUÁN CAFE")
	fmt.Printf("%-10s %-20s %-10s\n", "Mã", "Tên món", "Giá")

	for i := 0; i < count; i++ {
		fmt.Printf("%-10d %-20s %.0f VND\n", ids[i], names[i], prices[i])
	}
}

func deleteItem(ids *[MAX]int, names *[MAX]string, prices *[MAX]float64, count *int) {
	if *count == 0 {
		fmt.Println("Menu đang trống!")
		return
	}

	var id int
	fmt.Print("Nhập mã món cần xóa: ")
	fmt.Scan(&id)

	position := -1

	for i := 0; i < *count; i++ {
		if ids[i] == id {
			position = i
			break
		}
	}

	if position == -1 {
		fmt.Println("Không tìm thấy món!")
		return
	}

	for i := position; i < *count-1; i++ {
		ids[i] = ids[i+1]
		names[i] = names[i+1]
		prices[i] = prices[i+1]
	}

	*count--

	fmt.Println("Xóa món thành công!")
}

func updateItem(ids [MAX]int, names *[MAX]string, prices *[MAX]float64, count int) {
	if count == 0 {
		fmt.Println("Menu đang trống!")
		return
	}

	var id int
	fmt.Print("Nhập mã món cần sửa: ")
	fmt.Scan(&id)

	position := -1

	for i := 0; i < count; i++ {
		if ids[i] == id {
			position = i
			break
		}
	}

	if position == -1 {
		fmt.Println("Không tìm thấy món!")
		return
	}

	fmt.Print("Nhập tên món mới: ")
	fmt.Scan(&names[position])

	fmt.Print("Nhập giá mới: ")
	fmt.Scan(&prices[position])

	fmt.Println("Sửa món thành công!")
}

func main() {
	var ids [MAX]int
	var names [MAX]string
	var prices [MAX]float64

	count := 0
	choice := 0

	for choice != 5 {
		fmt.Println("\n COFFEE SHOP")
		fmt.Println("1. Thêm món")
		fmt.Println("2. Xóa món")
		fmt.Println("3. Sửa món")
		fmt.Println("4. Hiển thị menu")
		fmt.Println("5. Thoát")
		fmt.Println("")

		fmt.Print("Nhập lựa chọn: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addItem(&ids, &names, &prices, &count)

		case 2:
			deleteItem(&ids, &names, &prices, &count)

		case 3:
			updateItem(ids, &names, &prices, count)

		case 4:
			displayItems(ids, names, prices, count)

		case 5:
			fmt.Println("Đã thoát chương trình!")

		default:
			fmt.Println("Lựa chọn không hợp lệ!")
		}
	}
}