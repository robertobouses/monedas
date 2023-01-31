package gestor

import "fmt"

type Moneda struct {
	Nombre     string
	ValorDolar int
}

var Codigo int
var InventarioMonedas map[int]Moneda
var opcion string

func (m *Moneda) Formulario() map[int]Moneda {
	InventarioMonedas = make(map[int]Moneda)
	fmt.Println("Has iniciado el formulario para guardar Monedas")
	for {
		fmt.Println("Indica el código de la moneda")
		fmt.Scanln(&Codigo)
		fmt.Println("Indica el nombre de la moneda")
		fmt.Scanln(&m.Nombre)
		fmt.Printf("Indica cuantos %v hacen falta para comprar un Dolar\n", m.Nombre)
		fmt.Scanln(&m.ValorDolar)

		Moneda1 := Moneda{

			Nombre:     m.Nombre,
			ValorDolar: m.ValorDolar,
		}

		InventarioMonedas[Codigo] = Moneda1

		fmt.Println("Quiere añadir más monedas")
		fmt.Scanln(&opcion)

		if opcion == "NO" {
			break
		}
	}
	return InventarioMonedas
}
