package main

import (
	"fmt"
)

func main() {
	const (
		um = 1 << iota
		dois
		tres
		quatro
		cinco
		seis
		sete
		oito
	)

	slice := []int{um, dois, tres, quatro, cinco, seis, sete, oito}

	for _, bytes := range slice {
		resultado := fmt.Sprintf("%08b", bytes)
		fmt.Println(resultado, bytes)

	}
}
