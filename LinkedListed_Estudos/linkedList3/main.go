package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type linkedList struct {
	Head *Node
}

func (l *linkedList) Prepend(value int) {
	newNode := &Node{Data: value, Next: l.Head}
	l.Head = newNode
}

func (l linkedList) Print() {
	current := l.Head
	for current != nil {
		fmt.Println(current.Data)
		current = current.Next
	}
}

func (l *linkedList) Delete(value int) {
	if l.Head == nil {
		return
	}

	if l.Head.Data == value {
		l.Head = l.Head.Next
	}

	atual := l.Head
	for atual.Next != nil {
		if atual.Next.Data == value {
			atual.Next = atual.Next.Next
			return
		}
		atual = atual.Next
	}
}

func main() {
	myLink := linkedList{}

	myLink.Prepend(2)
	myLink.Prepend(4)
	myLink.Prepend(7)
	myLink.Prepend(34)
	myLink.Prepend(3)
	myLink.Delete(4)
	myLink.Print()

}
