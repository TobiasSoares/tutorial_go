package main

import (
	"fmt"
	"io"

	"os"

	"strings"
)

// Exercício 1: Loops e funções
//
// Implementar a função de raiz quadrada

// Sqrt calcula a raiz quadrada do número x passado como
// argumneto
/*
func Sqrt(x float64) float64 {
	z := float64(1)
	for i := 0; i < 10; i++ {
		fmt.Println(z)
		z -= (z*z - x) / (2 * z)
	}
	return z
}
*/

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

// String é a função para correta impressão do valor de IPAddr
func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

// Exercício 6: Errors
//
// Modificar a função Sqrt() anterior para retornar um erro

// Sqrt calcula a raiz quadrada de x se ele for postivo e
// retorna um erro se ele for negativo
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := float64(1)
	for i := 0; i < 10; i++ {
		fmt.Println(z)
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

// ErrNegativeSqrt é o tipo do erro
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

// Exercício 7: Readers
//
// Criar um tipo Reader que emite uma quantidade infinita de 'A'

// MyReader emite uma quantidade infinita de 'A'
type MyReader struct{}

func (mr MyReader) Read(destino []byte) (int, error) {
	lidos := 0
	for i := 0; i < len(destino); i++ {
		destino[i] = 'A'
		lidos++
	}
	return lidos, nil
}

// Exercício 8: rot13Reader
//
// Criar um tipo Reader que implemente o rot13

type rot13Reader struct {
	r io.Reader
}

func (rot13 rot13Reader) Read(destino []byte) (int, error) {
	lidos, eof := rot13.r.Read(destino)
	var novaLetra byte

	for i := 0; i < len(destino); i++ {
		if destino[i] < 65 || destino[i] > 123 ||
			destino[i] > 90 && destino[i] < 97 {
			continue
		}

		switch {
		case destino[i] >= 96:
			novaLetra = destino[i] + 13
			if novaLetra > 122 {
				novaLetra = (novaLetra % 122) + 96
			}
			destino[i] = novaLetra

		case destino[i] >= 65:
			novaLetra = destino[i] + 13
			if novaLetra > 90 {
				novaLetra = (novaLetra % 90) + 65
			}
			destino[i] = novaLetra
		}
	}

	return lidos, eof
}

// Main
func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
