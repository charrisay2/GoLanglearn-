package main

import "fmt"

func Add(items *[10]string, count *int, item string) {
    if *count >= len(items) {
        fmt.Println("San pham da day")
        return
    }

    items[*count] = item
    *count++
}

func Delete(items *[10]string, count *int, index int) {
    if index < 0 || index >= *count {
        fmt.Println("Vi tri khong hop le")
        return
    }

    for i := index; i < *count-1; i++ {
        items[i] = items[i+1]
    }

    items[*count-1] = ""
    *count--
}

func Update(items *[10]string, count int, index int, newItem string) {
    if index < 0 || index >= count {
        fmt.Println("Vi tri khong hop le")
        return
    }

    items[index] = newItem
}

func Display(items [10]string, count int) {
    for i := 0; i < count; i++ {
        fmt.Println(i, items[i])
    }
}

func main() {
    var items [10]string
    count := 0

    Add(&items, &count, "Coffee")
    Add(&items, &count, "Tea")
    Add(&items, &count, "Cake")

	Add(&items, &count, "chocolate")

    fmt.Println("Danh sach ban dau:")
    Display(items, count)

    Update(&items, count, 0, "Milk")

    fmt.Println("Sau khi sua:")
    Display(items, count)

    Delete(&items, &count, 0)

    fmt.Println("Sau khi xoa:")
    Display(items, count)
}