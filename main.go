package main

import (
	"fmt"

	"github.com/johnrios07/godesde0/ejercicios"
)

func main() {
	/*estado, texto := variables.ConviertoATexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)
	if os := runtime.GOOS; os == "Linux." || os == "OS X." {
		fmt.Println("Esto no  es windows")
	} else {
		fmt.Println("Esto es", os)
	}

	switch os := runtime.GOOS; os {
	case "Linux.":
		fmt.Println("Esto es Linux")
	case "windows":
		fmt.Println("Esto es windows")
	default:
		fmt.Printf("%s \n", os)
	}*/

	numero, texto := ejercicios.ConvNumerico("cdcd")
	fmt.Println(numero)
	fmt.Println(texto)
}
