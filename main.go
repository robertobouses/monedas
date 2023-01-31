package main

import (
	"github.com/robertobouses/monedas/gestor"
)

func main() {
	var InventarioMonedas map[int]gestor.Moneda

	valor := gestor.Moneda{}

	InventarioMonedas = valor.Formulario()
	gestor.Imprimir(InventarioMonedas)
	codigo1 := gestor.Entrada()
	codigo2 := gestor.Salida()
	gestor.Conversor(codigo1, codigo2, InventarioMonedas)
}
