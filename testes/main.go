package main

import "fmt"

func main() {

	// linhas := 2
	// colunas := 2

	// matriz := make([][]int, linhas)

	// for i := range matriz {
	// 	matriz[i] = make([]int, colunas)
	// }

	// fmt.Println(matriz)

	matriz := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	// for i, valor := range matriz {
	// 	fmt.Println(i, valor)
	// }

	// fmt.Println(len(matriz))

	// for i := 0; i < len(matriz); i++ {
	// 	for j := 0; j < len(matriz[i]); j++ {
	// 		fmt.Println(i, j, "valor:", matriz[i][j])
	// 	}
	// }

	for i, Linha := range matriz {
		for j, Valor := range Linha {
			fmt.Println(i, j, Valor)
		}
	}

	// tamanho := make([]int, 5)
	// tamanho = append(tamanho, 7, 8, 9)
	// fmt.Println(tamanho)
	// for i := 0; i < len(tamanho); i++ {
	// 	tamanho[i] = i + 1
	// }
	// fmt.Println(tamanho)
}
