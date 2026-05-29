/*teste sem significado, so tentando entender mais endereco de memoria */

package main

import (
	"encoding/binary"
	"fmt"
)

func main() {

	b := []byte{0x68, 0x09, 0x00, 0x00}

	novoSlice := make([]byte, 0)
	for _, bytes := range b {
		novoSlice = append(novoSlice, bytes<<1)
	}
	num := int(binary.LittleEndian.Uint32(b))
	novoNum := int(binary.LittleEndian.Uint32(novoSlice))
	fmt.Println(num)
	fmt.Println(novoNum)
	//          b := []byte{0x68, 0x09, 0x00, 0x00}
	resultadoNovo := []byte{0xD0, 0x12, 0x00, 0x00}

	resultado := int(binary.LittleEndian.Uint32(resultadoNovo))
	fmt.Println(resultado)
}

//00 00 12 D0
