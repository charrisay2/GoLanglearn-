package main

import "fmt"

/*Bài tập rèn luyện sử dụng Array: Viết chương trình quản lý quán cafe (Coffee Shop
Management System) bằng ngôn ngữ Go với yêu cầu sau:
1. Chỉ sử dụng Array (Không sử dụng bất kỳ data structure nào khác)
2. Cho phép thêm mới item, xóa item, sửa item và hiển thị toàn bộ item trong menu*/

func AddDrinks(d [10]string, pos int, drinks string) [10]string {
	d[pos] = drinks
	return d
}
func DeleteDrinks(d [10]string, pos int) [10]string {
	for i := pos; i < len(d)-1; i++ {
		d[i] = d[i+1]
	}
	d[len(d)-1] = ""
	return d
}

func UpdateDrinks(d [10]string, pos int, Newdrink string) [10]string {
	for i := 0; i < len(d)-1; i++ {
		if pos == i {
			d[pos] = Newdrink
		}
	}
	return d
}

func ShowMenu(d [10]string) {
	for _, v := range d {
		fmt.Println(v)
	}
}

func main() {
	Drinks := [10]string{"tra sưa", "ca phe", "banh bo", "banh tieu"}

	Drinks = AddDrinks(Drinks, 7, "suahat")
	Drinks = AddDrinks(Drinks, 5, "suahat")
	Drinks = DeleteDrinks(Drinks, 0)

	Drinks = UpdateDrinks(Drinks, 1, "com tam")
	// fmt.Println(Drinks)
	ShowMenu(Drinks)
}
