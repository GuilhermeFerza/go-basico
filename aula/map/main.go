package main

import (
	"fmt"
)

func main() {

	// mapa := make(map[string]int)
	// mapa["inicio"] = 1_000_000 //underline serve como o ponto para melhor visibilidade de numeros grandes
	// fmt.Println(mapa["inicio"])

	// mapa := make(map[int]int)
	// mapa[1] = 1
	// mapa[2] = 2
	// mapa[3] = 3
	// mapa[4] = 4
	// mapa[5] = 5

	// target := 3

	// for i := 1; i < len(mapa)+1; i++ {
	// 	if _, existe := mapa[i]; existe {
	// 		if mapa[i] == target {
	// 			fmt.Println("o numero, ", i, " existe no map")
	// 		}
	// 	} else {
	// 		mapa[i] = target
	// 	}

	// }
	// fmt.Println(mapa)

	mapa := make(map[int]struct{})

	mapa[1] = struct{}{}
	mapa[2] = struct{}{}
	mapa[3] = struct{}{}

	target := 4

	if _, existe := mapa[target]; existe {
		fmt.Println("a chave", target, "existe e tem valor: ")
	} else {
		fmt.Println("a chave", target, "nao existe")
	}

}
