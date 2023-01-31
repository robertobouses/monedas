package gestor

import "fmt"

func Entrada() int {
	fmt.Println("Indica que tipo de moneda tiene")
	fmt.Scanln(&Codigo)
	return Codigo
}
