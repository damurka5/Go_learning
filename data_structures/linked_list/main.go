package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

func main() {
	node1 := Node{3, nil}
	node2 := Node{5, &node1}
	node3 := Node{7, &node2}
	node4 := Node{10, &node3}
	pointer := &node4
	for pointer != nil {
		fmt.Println(pointer.value)
		pointer = pointer.next
	}
}
