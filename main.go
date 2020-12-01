package main

import (
	"fmt"

	"io"

	"strings"

	"golang.org/x/tour/tree"

	"image"

	"image/color"

	"sync"
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

// Exercício 9: Images
//
// Criar o tipo Image para construir uma imagem

// Image é o tipo que definde uma imagem
type Image struct{}

// ColorModel é a função que devolve o tipo de cor da imagem
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds retorna o retângulo contendo a imagem
func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 500, 250)
}

// At retorna a cor na posicao x, y informada
func (i Image) At(x, y int) color.Color {
	cor := uint8(x + y/2)
	return color.RGBA{cor, cor, 255, 255}
}

// Exercício 10: Equivalent Binary Trees
//
// Verificar se duas árvores binárias são equivalentes

// Walk caminha pela árvore enviando os valores contidos nos nodos
// para o canal
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same devolve true se ambas as árvores forem iguais
func Same(t1, t2 *tree.Tree) bool {
	canal1 := make(chan int, 10)
	canal2 := make(chan int, 10)

	go Walk(t1, canal1)
	go Walk(t2, canal2)

	for i := 0; i < 10; i++ {
		um, dois := <-canal1, <-canal2

		if um != dois {
			return false
		}
	}
	return true
}

// Exercício 11: Web Crawler
//
// Buscar URLs usando paralelismo

// Fetcher returns the body of URL and
// a slice of URLs found on that page.
type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

// Cache é uma estrutura que guarda a informação
// de quais URLs foram visitadas. Possui segurança de
// paralelismo (Mutex)
type Cache struct {
	tabela map[string]bool
	mux    sync.Mutex
}

// foiVisitado verifica se a URL está presente no
// cache de URLs visitadas
func (v *Cache) foiVisitado(url string) bool {
	v.mux.Lock()
	defer v.mux.Unlock()

	_, ocupado := v.tabela[url]

	if !ocupado {
		v.tabela[url] = true
		return false
	}
	return true
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, visitados *Cache, espera *sync.WaitGroup) {
	defer espera.Done()

	if depth <= 0 || visitados.foiVisitado(url) {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		espera.Add(1)
		Crawl(u, depth-1, fetcher, visitados, espera)
	}

	return
}

// Main
func main() {
	visitados := Cache{tabela: make(map[string]bool)}
	var espera sync.WaitGroup

	espera.Add(1)
	Crawl("https://golang.org/", 4, fetcher, &visitados, &espera)
	espera.Wait()
}
