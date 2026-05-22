package main

import (
	"fmt"
	"time"
)

func main() {

	inicio := time.Now()
	resultados := make([]int, 0, 9)
	resultados = append(resultados, 40)
	resultados = append(resultados, 50)
	resultados = append(resultados, 60)
	resultados = append(resultados, 40)
	resultados = append(resultados, 50)
	resultados = append(resultados, 60)
	resultados = append(resultados, 40)
	resultados = append(resultados, 50)
	resultados = append(resultados, 60)
	fmt.Println(resultados)

	time.Sleep(1 * time.Millisecond)

	duracao := time.Since(inicio)
	fmt.Println(duracao.Nanoseconds())
}
