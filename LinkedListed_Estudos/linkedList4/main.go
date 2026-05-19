package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

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

// Prepend, Print, Delete
type Node struct {
	Data int
	Next *Node
}

type linkedList struct {
	Head *Node
}

var myLink linkedList

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

func menu() {
	var escolha int
	var value int

	for {
		fmt.Println("*******************************************")
		fmt.Println("1-Adicionar No na list*********************")
		fmt.Println("2-Remover No da list***********************")
		fmt.Println("3-Visualizar list**************************")
		fmt.Println("0-Sair*************************************")
		fmt.Println("Escolha sua opcao")
		fmt.Scan(&escolha)
		limparTela()
		switch {
		case escolha == 1:
			fmt.Print("digite o numero que voce deseja adicionar ao No: ")
			fmt.Scan(&value)
			myLink.Prepend(value)
			limparTela()
		case escolha == 2:
			fmt.Print("digite o numero que voce deseja Remover ao No: ")
			fmt.Scan(&value)
			myLink.Prepend(value)
			limparTela()
		case escolha == 3:
			myLink.Print()
		default:
			fmt.Print("opcao invalida!")
			return
		}
	}

}

func main() {
	menu()
}
