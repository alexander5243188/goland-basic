package main

import "fmt"

// Calculadora avanzada: Suma, resta, multiplica y divide dos números ingresados por el usuario.
func suma(primer_numero int, segundo_numero int) int {
	return (primer_numero + segundo_numero)
}
func resta(primer_numero int, segundo_numero int) int {
	return (primer_numero - segundo_numero)
}
func multiplicacion(primer_numero int, segundo_numero int) int {
	return (primer_numero * segundo_numero)
}
func division(primer_valor int, segundo_valor int) float32 {
	return (float32(primer_valor / segundo_valor))
}

func main() {
	var primer_Numero, segundo_Numero int

	fmt.Println("Ingrese el primer numero: ")
	fmt.Scan(&primer_Numero)

	fmt.Println("Ingrese el segundo numero: ")
	fmt.Scan(&segundo_Numero)

	fmt.Println(primer_Numero, "+", segundo_Numero, "=", suma(primer_Numero, segundo_Numero))
	fmt.Println(primer_Numero, "-", segundo_Numero, "=", resta(primer_Numero, segundo_Numero))
	fmt.Println(primer_Numero, "*", segundo_Numero, "=", multiplicacion(primer_Numero, segundo_Numero))
	fmt.Println(primer_Numero, "/", segundo_Numero, "=", division(primer_Numero, segundo_Numero))
}
