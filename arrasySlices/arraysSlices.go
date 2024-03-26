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

	estructura_datos_dinamica := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Estructura de datos dinamica", estructura_datos_dinamica)
	fmt.Println("Tamaño de la estructura dinamica ", len(estructura_datos_dinamica))
	fmt.Println("Capacidad dela estructura dinamica ", cap(estructura_datos_dinamica))
	fmt.Println("[0]", estructura_datos_dinamica[0])
	fmt.Println("[:3]", estructura_datos_dinamica[:3])
	fmt.Println("[2:7]", estructura_datos_dinamica[2:7])
	fmt.Println("[5:8]", estructura_datos_dinamica[5:8])

	//Append
	estructura_datos_dinamica = append(estructura_datos_dinamica, 66)
	fmt.Println("Aplicando append", estructura_datos_dinamica)
}
