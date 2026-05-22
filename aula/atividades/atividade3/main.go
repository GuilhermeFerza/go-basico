/*
Dado um slice nums, retorne o elemento majoritário. O elemento majoritário é aquele que aparece na lista mais da metade das vezes (frequência maior que o tamanho do slice dividido por 2). Você pode assumir que ele sempre existe.

Entrada: nums = [2, 2, 1, 1, 1, 2, 2]

Saída: 2

Como pensar em Go: O map[int]struct{} não serve aqui, porque agora você precisa contar. Use um map[int]int onde a chave é o número e o valor é a frequência dele. O Go facilita muito isso: se uma chave não existe, o valor inicia em 0. Portanto, você pode simplesmente fazer mapa[num]++ dentro do loop sem precisar verificar se a chave já existia antes.
*/

package main

import "fmt"

func majori(vetor []int) int {
	mapa := make(map[int]int)
	metade := len(vetor) / 2
	for _, num := range vetor {
		mapa[num]++
		if mapa[num] > metade {
			return num
		}
	}
	return 0
}

func main() {
	vetor := []int{2, 2, 1, 1, 1, 2, 2}
	resposta := majori(vetor)
	fmt.Println(resposta)
}
