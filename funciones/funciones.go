package main

import "fmt"

func fecha(anio int, mes string) {
	fmt.Println(mes, " del ", anio)
}

func sumaDosValores(valorUno, valorDos int) int {
	suma := valorUno + valorDos
	return suma
}

// doble retorno
func dobleReturn(primerValor int) (valorPrimerRetorno int, valorSegundoRetorno string) {
	return primerValor, "Programa mas porfis"
}
func main() {
	var primerValor, segundoValor int

	fmt.Println("Ingrese el primer valor ")
	fmt.Scan(&primerValor)

	fmt.Println("Ingrese el segundo valor ")
	fmt.Scan(&segundoValor)

	sumatotal := sumaDosValores(primerValor, segundoValor)
	fmt.Println(primerValor, " + ", segundoValor, " = ", sumatotal)

	fecha(2024, "Marzo")

	// doble retorno
	dobleRetornoEntero, dobleRetornoString := dobleReturn(69)
	fmt.Println(dobleRetornoEntero)
	fmt.Println(dobleRetornoString)

	// omitir una variable  mediante _
	_, otroDobleRetornoString := dobleReturn(66)
	fmt.Println("Omitimos el valor entero ;)", otroDobleRetornoString)
}
