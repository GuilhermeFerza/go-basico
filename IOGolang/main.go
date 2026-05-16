package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var number1 int
var number2 int
var escolhaSwitch int
var total float64

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
	lerNome()
	testarVariaveis()
	escolha()
}

func lerNome() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Digite seu nome completo")
	if scanner.Scan() {
		nomeCompleto := scanner.Text()
		fmt.Println("ola!", nomeCompleto)
	} else {
		fmt.Println("Erro ao scannear")
	}
}

func testarVariaveis() {

	fmt.Println("fale o primeiro numero")
	fmt.Scan(&number1)
	fmt.Println("fale o segundo numero")
	fmt.Scan(&number2)
}

func escolha() {
	fmt.Println("escolha qual operacao voce quer da calculadora")
	fmt.Println("1-adicao")
	fmt.Println("2-subtracao")
	fmt.Println("3-multiplicao")
	fmt.Println("4-divisao")
	fmt.Scan(&escolhaSwitch)

	switch {
	case escolhaSwitch == 1:
		limparTela()
		tipo := "soma"
		adicao(number1, number2, tipo)
	case escolhaSwitch == 2:
		limparTela()
		tipo := "subtracao"
		subtracao(number1, number2, tipo)
	case escolhaSwitch == 3:
		limparTela()
		tipo := "multiplacao"
		multiplicacao(number1, number2, tipo)
	case escolhaSwitch == 4:
		limparTela()
		tipo := "divisao"
		divisao(number1, number2, tipo)
	default:
		fmt.Println("opcao invalida")
	}

}

func adicao(number1 int, number2 int, tipo string) {
	total := number1 + number2
	fmt.Println("o resultado da", tipo, total)
}

func subtracao(number1 int, number2 int, tipo string) {
	total := number1 - number2
	fmt.Println("o resultado da", tipo, total)
}

func multiplicacao(number1 int, number2 int, tipo string) {
	total := number1 * number2
	fmt.Println("o resultado da", tipo, total)
}

func divisao(number1 int, number2 int, tipo string) {
	total := number1 / number2
	fmt.Println("o resultado da", tipo, total)
}
