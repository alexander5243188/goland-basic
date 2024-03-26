package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
)

func main() {
	var numero_usuario int
	var numero_aleatorio int

	// Genera un número aleatorio entre 0 y 10
	numero_aleatorio = rand.Intn(10)

	fmt.Println("Ingrese un valor entre 1 y 10 :")
	fmt.Scan(&numero_usuario)

	if numero_usuario == numero_aleatorio {
		fmt.Println("Felicidades")
		fmt.Println("Usuario: ", numero_usuario, " ", numero_aleatorio, ": Ordenador")
	} else {
		fmt.Println("Usuario: ", numero_usuario, " ", numero_aleatorio, ": Ordenador")
	}

	// Un numero es par o impar
	if (numero_usuario % 2) == 0 {
		fmt.Println("Es par el numero ingresado", numero_usuario)
	} else {
		fmt.Println("Es impar el numero ingresado", numero_usuario)
		texto_numero()
	}
}

func texto_numero() {
	//convertir texto a numero
	fmt.Printf("Ingrese un valor .....")

	valor_leido := ""
	fmt.Scan(&valor_leido)

	value, err := strconv.Atoi(valor_leido)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Valor convertido a entero: ", value)
	}

}
