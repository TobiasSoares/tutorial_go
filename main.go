package main

import (
	"fmt"
)

// Exercício 1: Loops e funções
//
// Implementar a função de raiz quadrada

// Sqrt calcula a raiz quadrada do número x passado como
// argumneto
func Sqrt(x float64) float64 {
	z := float64(1)
	for i := 0; i < 10; i++ {
		fmt.Println(z)
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	Sqrt(400)
}
