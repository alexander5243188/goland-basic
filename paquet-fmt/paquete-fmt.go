package main

import "fmt"

func main() {
	mes := "Marzo"
	dia := 25

	fmt.Printf("hoy es %s %d del año 2024 \n", mes, dia)

	// en caso de no conocer que tipo de valor tendra
	fmt.Printf("%v %v, sera un buen día?\n", dia, mes)

	// guardar el mensaje en una variable
	mensaje := fmt.Sprintf("%v %v, sera un buen día?\n", dia, mes)
	fmt.Println(mensaje, " yo creo que si !!!!")

	// conocer el tipo de variable
	fmt.Printf("%T \n", mes)
	fmt.Printf("%T \n", dia)
}
