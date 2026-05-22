/*
2. Interseção de Dois Arrays
Você recebe dois slices de inteiros, nums1 e nums2. Retorne um novo slice contendo apenas os números exclusivos que aparecem em ambos (sem duplicatas na resposta).

Entrada: nums1 = [4, 9, 5], nums2 = [9, 4, 9, 8, 4]

Saída: [9, 4] ou [4, 9] (a ordem não importa)

Como pensar em Go: Esta é a aplicação perfeita para o seu map[int]struct{}. Varra o primeiro slice e coloque todos os números no map. Depois, varra o segundo slice verificando se o número atual existe no map. Se existir, adicione-o ao slice de resultados (usando append) e delete a chave do map para garantir que ele não seja adicionado novamente caso se repita no nums2.
*/

package main

import "fmt"

func intersecArrays(nums1 []int, nums2 []int) []int {

	resposta := make([]int, 0, 5)
	mapa := make(map[int]struct{})

	for _, num := range nums1 {
		mapa[num] = struct{}{}
	}

	for _, num := range nums2 {
		if _, existe := mapa[num]; existe {
			resposta = append(resposta, num)
			delete(mapa, num)
		}
	}
	return resposta
}

func main() {

	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	resultados := intersecArrays(nums1, nums2)
	fmt.Println(resultados)
}
