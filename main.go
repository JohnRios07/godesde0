package main

import (
	"github.com/johnrios07/godesde0/webserver"
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
	}
	numero, texto := ejercicios.ConvNumerico("cdcd")
	fmt.Println(numero)
	fmt.Println(texto)

	teclado.IngresoNumeros()
	iteraciones.Iterar()*/

	//files.GrabaTabla()

	//files.SumaTabla()

	//files.LeoArchivio()

	//arreglos_slice.Capacidad()

	//mapas.MostrarMapas()

	//users.AltaUsuario()

	/*Pedro := new(modelos.Hombre)
	e.HumanosRespirando(Pedro)

	Maria := new(modelos.Mujer)
	e.HumanosRespirando(Maria)*/

	//d.EjemploPanic()

	/*canal1 := make(chan bool)
	go goroutines.MiNombreLentooo("Alexander", canal1)
	defer func() {
		<-canal1
	}()
	fmt.Println("Estoy aquí")*/

	webserver.MiWebServer()

}
