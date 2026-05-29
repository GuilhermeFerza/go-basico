package main

import (
	"encoding/binary"
	"fmt"
)

func main() {

	b := []byte{0x20, 0xA1, 0x07, 0x00}

	num := int(binary.LittleEndian.Uint32(b))

	fmt.Println(num)

}
