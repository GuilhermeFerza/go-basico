// package main

// import "fmt"

// func main() {
// 	texto := "hello"

// 	fmt.Println(texto)
// 	sliceBytes := []byte(texto)
// 	sliceBytes[0] = 'H'

// 	texto = string(sliceBytes)

// 	fmt.Println(texto)

// 	s := "maçã"
// 	// for i := 0; i < len(s); i++ {
// 	// 	fmt.Println(s[i])
// 	// }

// 	for i, char := range s {
// 		fmt.Println(i, char)
// 	}

// 	a := "tentativa"

// 	for i := 0; i < 10000; i++ {
// 		a += "a"
// 	}
// 	fmt.Println(a)
// }

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func constroiString(texto []string) string {
// 	var builder strings.Builder

// 	for i := 0; i < 10; i++ {
// 		builder.WriteString("a")
// 	}

// 	texto1 := builder.String()

// 	fmt.Println(len(texto1))
// 	sliceBytes := []byte(texto1)
// 	texto2 := string(sliceBytes)
// 	fmt.Println(sliceBytes, texto2)
// 	return texto1
// }

// func main() {
// 	texto := []string{"sla da silva"}
// 	resposta := constroiString(texto)
// 	fmt.Println(resposta)
// }

package main

import "fmt"

func contagem(slice []string) map[string]int {
	mapa := make(map[string]int)

	for _, letra := range slice {
		mapa[letra]++
	}
	return mapa
}

func main() {
	s := "salve mano"
	slice := make([]string, 0, 5)

	for _, letra := range s {
		l := string(letra)
		slice = append(slice, l)
	}
	fmt.Println(slice)

	resposta := contagem(slice)
	fmt.Println(resposta)
}
