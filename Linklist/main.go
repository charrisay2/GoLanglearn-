package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (list *LinkedList) Prepend(data int) {
	newNode := &Node{data: data}

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
		return
	}

	newNode.next = list.head
	list.head = newNode
}

func (list *LinkedList) Append(data int) {
	newNode := &Node{data: data}

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
		return
	}

	list.tail.next = newNode
	list.tail = newNode
}

func (list *LinkedList) Search(data int) *Node {
	current := list.head

	for current != nil {
		if current.data == data {
			return current
		}

		current = current.next
	}

	return nil
}

func (list *LinkedList) Delete(data int) {
	if list.head == nil {
		return
	}

	if list.head.data == data {
		list.head = list.head.next

		if list.head == nil {
			list.tail = nil
		}

		return
	}

	current := list.head

	for current.next != nil {
		if current.next.data == data {
			if current.next == list.tail {
				list.tail = current
			}

			current.next = current.next.next
			return
		}

		current = current.next
	}
}

func (list *LinkedList) PrintList() {
	current := list.head

	for current != nil {
		fmt.Print(current.data, " ")
		current = current.next
	}

	fmt.Println()
}

func (list *LinkedList) Length() int {
	count := 0
	current := list.head

	for current != nil {
		count++
		current = current.next
	}

	return count
}

func main() {
	var list LinkedList

	list.Prepend(30)
	list.Prepend(20)
	list.Prepend(10)

	fmt.Print("Sau khi Prepend: ")
	list.PrintList()

	list.Append(40)
	list.Append(50)

	fmt.Print("Sau khi Append: ")
	list.PrintList()

	fmt.Println("Length:", list.Length())

	result := list.Search(30)

	if result != nil {
		fmt.Println("Tim thay:", result.data)
	} else {
		fmt.Println("Khong tim thay")
	}

	list.Delete(30)

	fmt.Print("Sau khi Delete 30: ")
	list.PrintList()

	fmt.Println("Length:", list.Length())
}
