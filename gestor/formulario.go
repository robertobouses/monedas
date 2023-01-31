package gestor

import "fmt"

var InventarioMonedas map[int]Moneda
var opcion string
InventarioMonedas=make(map[Codigo]Moneda)

func (m *Moneda) Formulario() []Moneda {
	fmt.Println("Has iniciado el formulario para guardar Monedas")
	for {
		fmt.Println("Indica el código de la moneda")
		fmt.Scanln(&m.Codigo)
		fmt.Println("Indica el código de la moneda")
		fmt.Scanln(&m.Nombre)
		fmt.Println("Indica el código de la moneda")
		fmt.Scanln(&m.ValorDolar)

		Moneda1 := Moneda{
			Codigo:     m.Codigo,
			Nombre:     m.Nombre,
			ValorDolar: m.ValorDolar,
		}

		InventarioMonedas = append(InventarioMonedas, Moneda1)

		fmt.Println("Quiere añadir más monedas")
		fmt.Scanln(&opcion)

		if opcion == "NO" {
			break
		}
	}
	return InventarioMonedas
}
