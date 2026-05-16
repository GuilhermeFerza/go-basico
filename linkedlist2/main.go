package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

type Node struct {
	Data int
	Next *Node
}

type linkedList struct {
	Head *Node
}

var value int

func limparTela() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (l *linkedList) Prepend(value int) {
	newNode := &Node{Data: value, Next: l.Head}
	l.Head = newNode

}

func (l linkedList) Print() {
	current := l.Head

	for current != nil {
		fmt.Printf("%d \n", current.Data)
		current = current.Next
	}
}

func main() {

	menu()

}

func menu() {
	myLink := linkedList{}

	for {
		var escolha int
		fmt.Println("1-adicionar a lista linkada")
		fmt.Println("2-printar lista linkada")
		fmt.Println("0-sair")
		fmt.Println("digite opcao...")
		fmt.Scan(&escolha)

		switch {
		case escolha == 1:
			fmt.Println("digite o numero para a lista linkada")
			fmt.Scan(&value)
			myLink.Prepend(value)
			limparTela()
		case escolha == 2:
			limparTela()
			myLink.Print()
		default:
			return
		}
	}

}
