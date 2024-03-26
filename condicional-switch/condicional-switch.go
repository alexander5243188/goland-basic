package main

import "fmt"

func main() {
	var valor_introducido int

	fmt.Println("Ingrese un numero")
	fmt.Scan(&valor_introducido)

	switch operacion_valor := (valor_introducido % 2); operacion_valor {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")

	}
	limite := valor_introducido
	switch {
	case limite == 9:
		fmt.Println("Ganaste ", limite, " puntos")
	case limite == 13:
		fmt.Println("Ganaste ", limite, " puntos")
	case limite == 17:
		fmt.Println("Perdiste ", limite, " puntos")
	default:
		fmt.Println("Se acabo")
	}
}
