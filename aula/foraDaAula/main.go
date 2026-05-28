//testando o Iota e o byteshifting

package main

import (
	"fmt"
)

func main() {

	const (
		um = 1 << iota
		dois
		tres
		quatro
	)

	message := fmt.Sprintf("%b", dois)

	fmt.Println(message)

}
