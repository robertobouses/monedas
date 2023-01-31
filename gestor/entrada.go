package gestor

import "fmt"

type Moneda struct {
	Codigo     int
	Nombre     string
	ValorDolar int
}

func (m *Moneda) Entrada() int {
	fmt.Println("Indica que tipo de moneda tiene")
	fmt.Scanln(&m.Codigo)
	return m.Codigo
}
