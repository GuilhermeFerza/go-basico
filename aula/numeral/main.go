package main

import "fmt"

func numeral(s string) int {

	valores := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0

	for i := 0; i < len(s); i++ {
		valorAtual := valores[s[i]]

		if i+1 < len(s) && valorAtual < valores[s[i+1]] {
			total -= valorAtual
		} else {
			total += valorAtual
		}
	}

	return total
}

func main() {

	s := "IV"
	resultado := numeral(s)
	fmt.Println(resultado)
}
