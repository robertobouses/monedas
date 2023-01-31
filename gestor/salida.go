package gestor

import "fmt"

func (m *Moneda) Salida() int {
	fmt.Println("Indica a que tipo de moneda quiere convertir")
	fmt.Scanln(&m.Codigo)
	return m.Codigo
}
