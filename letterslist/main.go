package main

import "fmt"

func main() {
	lista := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	listaNova := []string{}
	var escolha int
	fmt.Println("digite o numero de 1 a 26")
	fmt.Scan(&escolha)

	for i := 0; i < escolha; i++ {
		letra := lista[i]
		listaNova = append(listaNova, letra)
	}
	fmt.Println(listaNova)
}
