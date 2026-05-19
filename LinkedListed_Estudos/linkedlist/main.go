package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) Prepend(value int) {
	newNode := &Node{Data: value, Next: l.Head}

	l.Head = newNode
}

func (l *LinkedList) Print() {
	current := l.Head

	for current != nil {
		fmt.Printf("%d \n", current.Data)
		current = current.Next
	}
	return
}

func main() {
	myLink := &LinkedList{}

	myLink.Prepend(2)
	myLink.Prepend(1)
	myLink.Prepend(3)

	myLink.Print()
}
