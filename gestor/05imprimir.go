package gestor

import "fmt"

func Imprimir(InventarioMonedas map[int]Moneda) {
	for codigo, moneda := range InventarioMonedas {
		fmt.Println("Inventario de Divisas:")
		fmt.Printf("Código: %d Moneda: %v\n", codigo, moneda)
	}

}
