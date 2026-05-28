package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

var scanner = bufio.NewScanner(os.Stdin)
var Lista = make(map[string]int)

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

func main() {

	inicioExec()

}

func inicioExec() {

	fmt.Println("Digite seu nome para comecarmos a aplicacao: ")

	if scanner.Scan() {
		nomeCompleto := scanner.Text()
		page2(nomeCompleto)
	} else {
		fmt.Println("Erro ao scannear")
	}
}

func page2(nomeCompleto string) {
	fmt.Println("ola ", nomeCompleto)
	fmt.Println("iniciando programa...")
	time.Sleep(2 * time.Second)
	limparTela()
	menu()
}

func menu() {

	var escolha int

	for {
		fmt.Println("1-Adicionar item a lista do mercado")
		fmt.Println("2-Remover item da lista do mercado")
		fmt.Println("3-Ver itens do mercado")
		fmt.Println("0-Sair")
		fmt.Scan(&escolha)

		switch {
		case escolha == 1:
			addLista(Lista)
			limparTela()
		case escolha == 2:
			delLista(Lista)
			limparTela()
		case escolha == 3:
			verLista(Lista)
			limparTela()
		case escolha == 0:
			return
		default:
			fmt.Println("opcao invalida")
			time.Sleep(2 * time.Second)
		}
	}

}

func addLista(Lista map[string]int) {
	var nomeItem string
	var quantItem int
	fmt.Println("adicione o nome do item: ")
	fmt.Scan(&nomeItem)
	nomeCorreto := strings.ToLower(nomeItem)
	fmt.Println("quantos items? ")
	fmt.Scan(&quantItem)

	Lista[nomeCorreto] = quantItem

}

func verLista(Lista map[string]int) bool {
	num := 0
	fmt.Println(Lista)

	for {
		fmt.Println("continuar?")
		fmt.Println("digite 1: ")
		fmt.Scan(&num)
		if num == 1 {
			return true
		} else {
			limparTela()
			fmt.Println("tecla incorreta")
		}
	}

}

func delLista(Lista map[string]int) {
	item := ""
	escolha := ""
	fmt.Println(Lista)
	fmt.Println("digite qual item voce deseja excluir? ")
	fmt.Scan(&item)
	strings.ToLower(item)
	fmt.Println("confirmar exclusao: S/N...")
	fmt.Scan(&escolha)
	strings.ToLower(escolha)

	for {
		switch {
		case escolha == "ss" || escolha == "s":
			delete(Lista, item)
			return
		case escolha == "nn" || escolha == "n":
			limparTela()
			delLista(Lista)
		default:
			fmt.Println("opcao invalida")
		}

	}

}
