package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error

func TablaDeMultiplicar() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese número 1: ")
	if scanner.Scan() {
		numero, err = strconv.Atoi(scanner.Text())
		if err != nil {
			TablaDeMultiplicar()
		}

		for i := 1; i <= 10; i++ {
			fmt.Printf("%d x %d = %d \n", numero, i, numero*i)
		}
	}
}
