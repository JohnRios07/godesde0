package main

import (
	"fmt"

	"github.com/johnrios07/godesde0/variables"
)

func main() {
	estado, texto := variables.ConviertoATexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)
}
