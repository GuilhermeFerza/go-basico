// // /*teste sem significado, so tentando entender mais endereco de memoria */

// // package main

// // import (
// // 	"encoding/binary"
// // 	"fmt"
// // )

// // func main() {

// // 	b := []byte{0x68, 0x09, 0x00, 0x00}

// // 	novoSlice := make([]byte, 0)
// // 	num := int(binary.LittleEndian.Uint32(b))
// // 	num << 1
// // 	novoNum := int(binary.LittleEndian.Uint32(novoSlice))
// // 	fmt.Println(num)
// // 	fmt.Println(novoNum)
// // 	//          b := []byte{0x68, 0x09, 0x00, 0x00}
// // 	resultadoNovo := []byte{0xD0, 0x12, 0x00, 0x00}

// // 	resultado := int(binary.LittleEndian.Uint32(resultadoNovo))
// // 	fmt.Println(resultado)
// // }

// // //00 00 12 D0

// /*teste sem significado, so tentando entender mais endereco de memoria */

// package main

// import (
// 	"encoding/binary"
// 	"fmt"
// )

// func main() {

// 	b := []byte{0x68, 0x09, 0x00, 0x00}

// 	num := int(binary.LittleEndian.Uint32(b))
// 	byteShifted := num >> 1
// 	fmt.Println(byteShifted)

// 	var bytes byte = 1

// 	novoBytes := bytes << 1
// 	numByte := byte(binary.)
// 	fmt.Println(novoBytes)

// }

// //00 00 12 D0

package main

import (
	"fmt"
)

func main() {

	num := 255
	alt := fmt.Sprintf("%08b", num)
	fmt.Println(alt)

}
