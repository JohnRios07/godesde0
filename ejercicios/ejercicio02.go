package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error
var texto string

func TablaDeMultiplicar() string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese número 1: ")
	if scanner.Scan() {
		numero, err = strconv.Atoi(scanner.Text())
		if err != nil {
			TablaDeMultiplicar()
		}
		fmt.Printf("\n")
		for i := 1; i <= 10; i++ {
			texto += fmt.Sprintf("%d x %d = %d \n", numero, i, numero*i)
		}
	}

	return texto
}
