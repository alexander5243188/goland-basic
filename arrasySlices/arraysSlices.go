package main

import (
	"fmt"
)

func main() {
	var arreglo [4]int
	fmt.Println("Arreglo antes de su llenado: ", arreglo)

	fmt.Println("Ingrese el primer valor")
	fmt.Scan(&arreglo[0])

	fmt.Println("Ingrese el segundo valor")
	fmt.Scan(&arreglo[1])

	fmt.Println("Ingrese el tercer valor")
	fmt.Scan(&arreglo[2])

	fmt.Println("Ingrese el cuarto valor")
	fmt.Scan(&arreglo[3])

	fmt.Println("Array despues de su llenado", arreglo)

	fmt.Println("Tamaño del arreglo ", len(arreglo), " elementos")
	fmt.Println("Capacidad del arreglo", cap(arreglo), " elementos")

}
