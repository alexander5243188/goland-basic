package main

import "fmt"

func sumaValores(primerNumero int, segundoNumero int) int {
	return (primerNumero + segundoNumero)
}

func main() {
	//Calculadora básica: Suma dos números ingresados por el usuario.
	var primer_Numero int
	var segundo_Numero int

	fmt.Println("Ingrese el primer numero: ")
	fmt.Scan(&primer_Numero)

	fmt.Println("Ingrese el segundo numero: ")
	fmt.Scan(&segundo_Numero)

	fmt.Println(primer_Numero, "+", segundo_Numero, "=", sumaValores(primer_Numero, segundo_Numero))

}
