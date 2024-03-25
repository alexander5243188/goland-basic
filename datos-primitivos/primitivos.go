package main

import "fmt"

func main() {
	fmt.Println("Datos primitivos en GO")

	// Enteros
	var temperatura int8 = 15
	fmt.Println("Temperatura ", temperatura)

	// Flotantes
	var pi float64 = 3.14159
	fmt.Println("Pi ", pi)

	//Booleanos
	var esSoleado bool = false
	fmt.Println("Es un tia soleado ", esSoleado)

	// Cadenas de texto
	var mensaje string = "!Hola mundo"
	fmt.Println(mensaje)

	// Caracter
	var letra byte = 'A'
	fmt.Println(letra)

	// Enteros sin signo
	var cantidad uint16 = 100
	fmt.Println("Cantidad ", cantidad)

	// Complejo
	var numeroComplejo complex64 = 3 + 4i
	fmt.Println(numeroComplejo)
}
