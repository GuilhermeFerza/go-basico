/*
	conta quantas vezes um objeto aparece em um slice
*/

// package main

// import "fmt"

// func quantCounter(vetor []string) {

// 	mapa := make(map[string]int)
// 	counter := 0
// 	for _, modelo := range vetor {
// 		if _, existe := mapa[modelo]; existe {
// 			counter++
// 			mapa[modelo] = counter
// 		} else {
// 			counter = 1
// 			mapa[modelo] = 1
// 		}
// 	}
// 	fmt.Println(mapa)
// }

// func main() {
// 	vetor := []string{"iphone", "iphone", "iphone", "android", "windows", "raspberry", "windows"}
// 	quantCounter(vetor)
// }

package main

import "fmt"

func quantCounter(vetor []string) {
	mapa := make(map[string]int)
	for _, nome := range vetor {
		mapa[nome]++
	}

	fmt.Println(mapa)
}

func main() {

	vetor := []string{"sol", "sol", "lua", "lua", "lua", "terra", "marte"}
	quantCounter(vetor)
}

/*
O motivo de mapa[modelo]++ funcionar é que, em Go, quando a chave ainda não existe em um map[int], o valor padrão retornado é 0. Então:

mapa["iphone"]++

equivale a:

mapa["iphone"] = mapa["iphone"] + 1

*/
