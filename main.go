package main

import (
	"fmt"

	"strings"
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

// Exercício 4: Closures
//
// Implementar a função fibonacci

// fibonacci é uma "fuction closure" que calcula números
// consecutivos da sequencia de fibonacci
func fibonacci() func() int {
	var indice, anterior, atual, proximo int

	return func() int {
		if indice == 0 {
			indice++
			return 0
		}
		if indice == 1 {
			indice++
			atual = 1
			return 1
		}

		anterior = atual
		atual = proximo
		proximo = anterior + atual
		indice++
		return proximo
	}
}

// Exercício 5: Stringers
//
// Implementar a função de print do tipo IPAddr

// IPAddr tipo que guarda um endereço de IP
type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

// Main
func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
