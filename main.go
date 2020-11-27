package main

import (
	"fmt"

	"strings"

	"golang.org/x/tour/wc"
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

// Exercício 2: Slices
//
// Implementar a função Pic

// Pic é uma Função para criar uma imagem através
// de uma função
func Pic(dx, dy int) [][]uint8 {
	saida := make([][]uint8, dy)
	for i := range saida {
		saida[i] = make([]uint8, dx)
		for j := range saida {
			saida[i][j] = uint8((i + j) / 2)
		}
	}
	return saida
}

// Exercício 3: Maps
//
// Implementar a função WordCount

// WordCount conta o número de palavras em uma string s
// passada como argumneto
func WordCount(s string) map[string]int {
	mapa := make(map[string]int)
	palavras := strings.Fields(s)

	for _, pal := range palavras {
		mapa[pal]++
	}
	return mapa
}

// MAIN
func main() {
	wc.Test(WordCount)
}
