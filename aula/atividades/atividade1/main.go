/*
Você recebe um slice de inteiros nums e um número inteiro target. Você precisa retornar os índices (as posições no slice, não os valores) dos dois números que somados resultam no target.

Entrada: nums = [2, 7, 11, 15], target = 9

Saída: [0, 1] (pois nums[0] + nums[1] == 9)

Como pensar em Go: Um loop duplo funcionaria, mas é lento. Use um map[int]int onde a chave é o número do slice e o valor é o índice dele. Para cada número que você visitar no for range, calcule o "complemento" (target - num). Se o complemento já estiver no map, você achou o par!


*/

// package main

// import "fmt"

// func twoSums(nums []int, target int) {

// 	var cache int
// 	mapa := make(map[int]int)
// 	for i, num := range nums {
// 		mapa[num] = i
// 		cache = target - num
// 		if _, existe := mapa[cache]; existe {
// 			fmt.Println(mapa[cache], i)
// 		}
// 	}

// }

// func main() {
// 	target := 19
// 	nums := []int{2, 7, 11, 15}
// 	twoSums(nums, target)
// }

package main

import "fmt"

// No LeetCode, a função deve retornar []int
func twoSum(nums []int, target int) []int {
	mapa := make(map[int]int)

	for i, num := range nums {
		complemento := target - num

		// 1. PRIMEIRO checamos se o complemento já existe no map
		if indiceComplemento, existe := mapa[complemento]; existe {
			// Se existe, retornamos o índice do complemento salvo antes, e o índice atual 'i'
			return []int{indiceComplemento, i}
		}

		// 2. DEPOIS, se não achou, salvamos o número atual e seu índice no map
		mapa[num] = i
	}

	// Retorno padrão caso não encontre (o LeetCode diz que sempre haverá 1 resposta)
	return nil
}

func main() {
	target := 9
	nums := []int{2, 7, 11, 15}
	resultado := twoSum(nums, target)
	fmt.Println("Os índices são:", resultado)
}
