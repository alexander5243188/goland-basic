package main

import "fmt"

func main() {
	// Declaracion de constantes
	const cambioDolar float64 = 6.94
	const cambioEuro = 7.1

	fmt.Println("Cambio de dolar: ", cambioDolar)
	fmt.Println("Cambio de euro: ", cambioEuro)

	// Declaracion de variables
	anio := 2024
	var mes string = "marzo"
	var dia int

	fmt.Println(anio, mes, dia)

	//Asignacion de valor por defecto
	var entero int
	var decimal float64
	var cadena string
	var booleano bool

	fmt.Println(entero, decimal, cadena, booleano)
}
