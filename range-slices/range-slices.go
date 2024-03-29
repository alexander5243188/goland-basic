package main

import "fmt"

func main() {
	arreglo_dias := []string{"Lunes", "Martes", "Miercoles", "Jueves"}

	fmt.Println("Recorrido con indice")
	for i, valor := range arreglo_dias {
		fmt.Println(i, valor)
	}
	fmt.Println("Recorrido sin indice: ;)")
	for _, valor := range arreglo_dias {
		fmt.Println(valor)
	}

	fmt.Println("Solo conocer el indice ;)")
	for i := range arreglo_dias {
		fmt.Println(i)
	}
}
