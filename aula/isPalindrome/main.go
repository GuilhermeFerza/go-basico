// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func isPalindrome(x int) bool {

// 	resultadoConv := make([]int, 0)
// 	str := strconv.Itoa(x)
// 	for _, rune := range str {
// 		char := string(rune)
// 		num, _ := strconv.Atoi(char)
// 		resultadoConv = append(resultadoConv, num)
// 	}
// 	fmt.Println(resultadoConv)
// 	for i := 0; i < len(resultadoConv); i++ {
// 		pointer1 := resultadoConv[i]
// 		for j := len(resultadoConv) - 1; j > 0; j-- {
// 			pointer2 := resultadoConv[j]
// 			if pointer1 != pointer2 {
// 				return false
// 			}
// 		}
// 	}
// 	return false
// }

// func main() {

// 	var x int
// 	x = 121
// 	resultado := isPalindrome(x)
// 	fmt.Println(resultado)
// }

// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func isPalindrome(x int) bool {

// 	str := strconv.Itoa(x)

// 	left := 0
// 	right := len(str) - 1

// 	for left < right {
// 		if str[left] != str[right] {
// 			return false
// 		}

// 		left++
// 		right--
// 	}

// 	return true
// }

// func main() {

// 	x := 1231
// 	resultado := isPalindrome(x)
// 	fmt.Println(resultado)
// }

package main

import "fmt"

func main() {
	str := "texto"
	fmt.Println(string(str[3]))
}
