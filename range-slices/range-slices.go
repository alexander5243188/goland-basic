package main

import (
	"fmt"
	"strings"
)

// funcion palindromo
func palabraAlReves(palabra string) string {
	var palabraReves string
	for i := len(palabra) - 1; i >= 0; i-- {
		palabraReves += string(palabra[i])
	}
	return palabraReves
}

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

	// Palindromo
	// Declaramos la variable
	var texto string

	// Ingresar el usuario la palabra
	fmt.Println("Ingresa una palabra :")

	// Almacenamos la informacion
	fmt.Scan(&texto)

	//Convertimos la palabra a minusculas
	texto = strings.ToLower(texto)

	// comprobamos, si lo que retorna la funcion es la misma palabra
	if palabraAlReves(texto) == texto {
		fmt.Println("Ok, es palindromo: ", (palabraAlReves(texto)), "=", texto)
	} else {
		fmt.Println("No es palindromo", palabraAlReves(texto), "=", texto)
	}

}
