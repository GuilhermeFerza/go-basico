// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {

// 	var construtor strings.Builder

// 	construtor.Grow(10000)

// 	for i := 0; i < 10000; i++ {
// 		construtor.WriteByte('a')
// 	}

// 	resultado := construtor.String()
// 	fmt.Println(resultado)
// }

// package main

// import "fmt"

// func isPalindrome(x []string) bool {

// 	left := 0
// 	right := len(x) - 1

// 	for left < right {
// 		if x[left] != x[right] {
// 			return false
// 		}
// 		left++
// 		right--
// 	}
// 	return true
// }

// func main() {

// 	// x := []string{"b", "a", "n", "a", "n", "a"}
// 	x := []string{"b", "a", "n", "a", "b"}

// 	resultado := isPalindrome(x)
// 	fmt.Println(resultado)
// }

// package main

// import "fmt"

// func twoSum(slice []int, target int) []int {

// 	complemento := 0
// 	mapa := make(map[int]int)

// 	for i, num := range slice {

// 		complemento = target - num

// 		if iComplemento, existe := mapa[complemento]; existe {
// 			return []int{iComplemento, i}
// 		}
// 		mapa[num] = i
// 	}
// 	return nil
// }

// func main() {

// 	slice := []int{2, 7, 11, 15}

// 	target := 22

// 	resultados := twoSum(slice, target)

// 	fmt.Println(resultados)

// }

// package main

// import "fmt"

// func inverterString(s string) string {

// 	b := []byte(s)

// 	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
// 		b[i], b[j] = b[j], b[i]
// 	}
// 	return string(b)
// }

// func main() {

// 	s := "Guilherme"

// 	resultados := inverterString(s)
// 	fmt.Println(resultados)
// }

// package main

// import "fmt"

// func RomanToNum(s string) int {

// 	mapa := map[byte]int{
// 		'I': 1,
// 		'V': 5,
// 		'X': 10,
// 		'L': 50,
// 		'C': 100,
// 		'D': 500,
// 		'M': 1000,
// 	}

// 	total := 0

// 	for i := 0; i < len(s); i++ {
// 		atual := mapa[s[i]]

// 		if i+1 < len(s) && atual < mapa[s[i+1]] {
// 			total -= atual
// 		} else {
// 			total += atual
// 		}
// 	}

// 	return total

// }

// func main() {

// 	s := "V"

// 	resultados := RomanToNum(s)

// 	fmt.Println(resultados)

// }

// package main

// import "fmt"

// func twoSum(vetor []int, target int) []int {

// 	complemento := 0
// 	mapa := make(map[int]int)

// 	for i, num := range vetor {
// 		complemento = target - num
// 		if indexComplemento, existe := mapa[complemento]; existe {
// 			return []int{indexComplemento, i}
// 		}
// 		mapa[num] = i
// 	}

// 	return nil

// }

// func main() {

// 	target := 18
// 	vetor := []int{2, 7, 9, 11, 15}

// 	fmt.Println(twoSum(vetor, target))
// }

package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(valores int) bool {

	s := strconv.Itoa(valores)
	valor := make([]int, 0, 5)

	for _, num := range s {
		valor = append(valor, int(num-'0'))
	}

	left := 0
	right := len(valor) - 1

	for left < right {
		if valor[left] != valor[right] {
			return false
		} else {
			left++
			right--

		}

	}
	return true
}

func main() {

	valor := 12321
	fmt.Println(isPalindrome(valor))
}
