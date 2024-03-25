package main

import "fmt"

func main() {
	var primer_numero, segundo_numero int

	fmt.Println("Ingresa el primer numero: ")
	fmt.Scan(&primer_numero)

	fmt.Println("Ingresa el segundo numero: ")
	fmt.Scan(&segundo_numero)

	suma := primer_numero + segundo_numero
	fmt.Println("Suma ", suma)

	resta := primer_numero - segundo_numero
	fmt.Println("Resta ", resta)

	multiplicacion := primer_numero * segundo_numero
	fmt.Println("Multiplicacion ", multiplicacion)

	division := primer_numero / segundo_numero
	fmt.Println("Division ", division)

	modulo := primer_numero % segundo_numero
	fmt.Println("Residuo", modulo)

	primer_numero++
	fmt.Println("Incremento ", primer_numero)

	segundo_numero--
	fmt.Println("Decreemnto ", segundo_numero)

}
