package main

// Adicionado o import do rsc.io/quote
import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
}
