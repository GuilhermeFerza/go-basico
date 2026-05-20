package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isPalindrome(x int) bool {
	var digitos []int

	numeroOriginal := x

	for x > 0 {
		digitos = append(digitos, x%10)
		x /= 10
	}

	var strSlice []string
	for _, num := range digitos {
		strSlice = append(strSlice, strconv.Itoa(num))
	}

	juntos := strings.Join(strSlice, "")

	resultado, err := strconv.Atoi(juntos)
	if err != nil {
		fmt.Println("erro ao converter os caracteres para inteiros", err)
	}
	if resultado == numeroOriginal {
		fmt.Println("verdadeiro")
		return true
	}
	fmt.Println("falso")
	return false

}

func main() {
	x := 121
	isPalindrome(x)
}
