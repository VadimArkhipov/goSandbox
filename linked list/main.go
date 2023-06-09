package main

import "fmt"

type node struct {
	value int
	next  *node
}

type List struct {
	head *node
}

func (l *List) append(item int) {
	if l.head == nil {
		l.head = &node{item, nil}
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		nextNode := node{item, nil}
		current.next = &nextNode
	}
}

func (l *List) remove(value int) {
	if l.head == nil {
		return
	}

	if l.head.value == value {
		l.head = l.head.next
		return
	}

	current := l.head
	for current.next != nil && current.next.value != value {
		current = current.next
	}

	if current.next != nil {
		current.next = current.next.next
	}
}

func (l List) printValues() {
	current := l.head
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}

func main() {
	list := List{}

	list.append(1)
	list.append(2)
	list.append(3)
	list.append(4)
	list.append(5)

	list.remove(1)

	list.printValues()

}
