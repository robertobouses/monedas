package main

import "fmt"

func (m Moneda) Conversor(Codigo1, Codigo2 int) {
	fmt.Println("Indica que cantidad de la divisa a convertir %s tienes")
	fmt.Scanln(&CantidadDivisa)
	Dolares := CantidadDivisa * ValorDolar1
	Resultado := Dolares / ValorDolar2
	fmt.Println("La divisa %s de la que tenías %v equivale a %v %v", Nombre, CantidadDivisa, Resultado, Nombre)
}
