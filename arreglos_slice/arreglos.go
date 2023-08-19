package arreglos_slice

import "fmt"

var tabla [10]int = [10]int{10, 0, 59}
var matriz [20][30]int

func MustroArreglos() {
	tabla[7] = 33
	tabla[2] = 54

	tabla2 := [10]string{"pablo", "Juan"}
	fmt.Println(tabla)
	fmt.Println(tabla2)

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}

	matriz[7][24] = 15

	fmt.Println(matriz)
}
