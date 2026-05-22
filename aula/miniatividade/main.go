/*
Verifica se um Slice tem numeros repetidos utilizando for range, colocando em um map a cada iteracao e verificando se a chave se repete
*/
package main

import "fmt"

func findDuplicate() {
	mapa := make(map[int]struct{})
	problema := []int{1, 2, 3, 4, 4, 6, 7, 8, 9}

	for _, valor := range problema {
		if _, existe := mapa[valor]; existe {
			fmt.Println("o valor ", valor, "ja existe no mapa")
			return
		}
		mapa[valor] = struct{}{}
	}
	fmt.Println("nenhum valor repetido")
}

func main() {

	findDuplicate()
}
