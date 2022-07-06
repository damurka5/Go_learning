package main

import "fmt"

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
	new_node := Node{
		value: value,
		next:  nil,
	}
	if l.len == 0 {
		l.head = &new_node
		l.len++
		return
	}

	pointer := l.head
	for i := 0; i < l.len; i++ {
		if pointer.next == nil {
			pointer.next = &new_node
			new_node.prev = pointer
			l.len++
			return
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
}
