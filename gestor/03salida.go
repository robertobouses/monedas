package gestor

import "fmt"

func Salida() int {
	fmt.Println("Indica a que tipo de moneda quiere convertir")
	fmt.Scanln(&Codigo)
	return Codigo
}
