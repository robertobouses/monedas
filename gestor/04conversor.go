package gestor

import "fmt"

var CantidadDivisa int

func Conversor(Codigo1, Codigo2 int, InventarioMonedas map[int]Moneda) {
	_, ok := InventarioMonedas[Codigo1]
	//fmt.Println(moneda)
	if ok {

		fmt.Printf("Indica que cantidad de la divisa a convertir %s tienes\n", InventarioMonedas[Codigo1].Nombre)
		fmt.Scanln(&CantidadDivisa)
		Dolares := CantidadDivisa / InventarioMonedas[Codigo1].ValorDolar
		Resultado := Dolares * InventarioMonedas[Codigo2].ValorDolar
		fmt.Println("")
		fmt.Printf("%v %s equivalen a %v %s", CantidadDivisa, InventarioMonedas[Codigo1].Nombre, Resultado, InventarioMonedas[Codigo2].Nombre)
	} else {
		fmt.Println("Indica una moneda que sí esté en el inventario de monedas")

	}
}
