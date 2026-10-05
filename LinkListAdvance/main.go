package main

import "fmt"

type Node[T any] struct {
	data T
	next *Node[T]
}

type LinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
}

func (list *LinkedList[T]) Prepend(data T) {
	newNode := &Node[T]{data: data}

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
		return
	}

	newNode.next = list.head
	list.head = newNode
}

func (list *LinkedList[T]) Append(data T) {
	newNode := &Node[T]{data: data}

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
		return
	}

	list.tail.next = newNode
	list.tail = newNode
}

func (list *LinkedList[T]) Search(data T, equal func(T, T) bool) *Node[T] {
	current := list.head

	for current != nil {
		if equal(current.data, data) {
			return current
		}

		current = current.next
	}

	return nil
}

func (list *LinkedList[T]) Delete(data T, equal func(T, T) bool) {
	if list.head == nil {
		return
	}

	if equal(list.head.data, data) {
		list.head = list.head.next

		if list.head == nil {
			list.tail = nil
		}

		return
	}

	current := list.head

	for current.next != nil {
		if equal(current.next.data, data) {
			if current.next == list.tail {
				list.tail = current
			}

			current.next = current.next.next
			return
		}

		current = current.next
	}
}

func (list *LinkedList[T]) PrintList() {
	current := list.head

	for current != nil {
		fmt.Print(current.data, " ")
		current = current.next
	}

	fmt.Println()
}

func (list *LinkedList[T]) Length() int {
	count := 0
	current := list.head

	for current != nil {
		count++
		current = current.next
	}

	return count
}

func (list *LinkedList[T]) Reverse() {
	var prev *Node[T]
	current := list.head

	list.tail = list.head

	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}

	list.head = prev
}

func (list *LinkedList[T]) InsertAt(index int, data T) {
	if index < 0 || index > list.Length() {
		return
	}

	if index == 0 {
		list.Prepend(data)
		return
	}

	if index == list.Length() {
		list.Append(data)
		return
	}

	newNode := &Node[T]{data: data}
	current := list.head

	for i := 0; i < index-1; i++ {
		current = current.next
	}

	newNode.next = current.next
	current.next = newNode
}

func main() {
	var list LinkedList[int]

	list.Prepend(30)
	list.Prepend(20)
	list.Prepend(10)

	fmt.Print("Prepend: ")
	list.PrintList()

	list.Append(40)
	list.Append(50)

	fmt.Print("Append: ")
	list.PrintList()

	fmt.Println("Length:", list.Length())

	result := list.Search(30, func(a, b int) bool {
		return a == b
	})

	if result != nil {
		fmt.Println("Search 30:", result.data)
	} else {
		fmt.Println("Khong tim thay")
	}

	list.Delete(30, func(a, b int) bool {
		return a == b
	})

	fmt.Print("Delete 30: ")
	list.PrintList()

	list.InsertAt(1, 100)

	fmt.Print("InsertAt index 1: ")
	list.PrintList()

	list.Reverse()

	fmt.Print("Reverse: ")
	list.PrintList()

	var stringList LinkedList[string]

	stringList.Append("A")
	stringList.Append("B")
	stringList.Append("C")

	fmt.Print("String List: ")
	stringList.PrintList()

	stringList.Reverse()

	fmt.Print("String Reverse: ")
	stringList.PrintList()
}