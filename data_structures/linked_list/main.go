package main

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type LinkedList struct {
	head *Node // the first element od the LinkedList
	len  int
}

// adds new node to the LinkedList as a head
func (l *LinkedList) addNode(value int) {
	newNode := Node{
		value: value,
		next:  nil,
	}
	if l.len == 0 {
		l.head = &newNode
		l.len++
		return
	}

	pointer := l.head
	for i := 0; i < l.len; i++ {
		if pointer.next == nil {
			pointer.next = &newNode
			newNode.prev = pointer
			l.len++
			return
		}
		pointer = pointer.next
	}
}

// deletes the first node with a given value
func (l *LinkedList) deleteNode(value int) {
	pointer := l.head
	for pointer != nil {
		if pointer.value == value {
			if pointer.prev == nil {
				l.head = pointer.next
			} else {
				temp := pointer.prev
				temp.next = pointer.next
			}
			pointer = nil
			break
		}
		pointer = pointer.next
	}
}

// inserts new node after given node (first met value)
func (l *LinkedList) insertAfter(value int, newValue int) {
	if l.len == 0 { // this is an empty list
		return
	}

	newNode := Node{
		value: newValue,
	}

	pointer := l.head
	for pointer != nil {
		if pointer.value == value {
			newNode.next = pointer.next
			pointer.next = &newNode
			newNode.prev = pointer
			break
		}
		pointer = pointer.next
	}
}

func (l LinkedList) printList() {
	pointer := l.head
	for pointer != nil {
		fmt.Println(pointer.value)
		pointer = pointer.next
	}
}

func main() {
	link := LinkedList{}
	link.addNode(1)
	link.addNode(2)
	link.addNode(3)
	link.addNode(4)
	link.printList()
	link.insertAfter(2, 35)
	link.printList()
}
