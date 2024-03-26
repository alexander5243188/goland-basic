package main

import "fmt"

func main() {
	var valor_uno int16 = 1
	var valor_dos int16 = 2

	// Operadores de comparacion
	fmt.Println("Operadores de comparación")
	fmt.Println(valor_uno, "==", valor_dos, valor_uno == valor_dos)
	fmt.Println(valor_uno, "!=", valor_dos, valor_uno != valor_dos)
	fmt.Println(valor_uno, "<", valor_dos, valor_uno < valor_dos)
	fmt.Println(valor_uno, ">", valor_dos, valor_uno > valor_dos)
	fmt.Println(valor_uno, "<=", valor_dos, valor_uno <= valor_dos)
	fmt.Println(valor_uno, ">=", valor_dos, valor_uno >= valor_dos)

	// Operador AND
	fmt.Println("Operador AND")
	fmt.Println((valor_uno == valor_uno), "&&", (valor_dos == valor_dos), (valor_uno == valor_uno) && (valor_dos == valor_dos))
	fmt.Println((valor_uno == valor_uno), "&&", (valor_uno == valor_dos), (valor_uno == valor_uno) && (valor_uno == valor_dos))
	fmt.Println((valor_uno == valor_dos), "&&", (valor_uno == valor_uno), (valor_uno == valor_dos) && (valor_uno == valor_uno))
	fmt.Println((valor_uno == valor_dos), "&&", (valor_uno == valor_dos), (valor_uno == valor_dos) && (valor_uno == valor_dos))

	// Operador OR
	fmt.Println("Operador OR")
	fmt.Println((valor_uno == valor_uno), "||", (valor_dos == valor_dos), (valor_uno == valor_dos) || (valor_dos == valor_dos))
	fmt.Println((valor_uno == valor_uno), "||", (valor_uno == valor_dos), (valor_uno == valor_uno) || (valor_uno == valor_dos))
	fmt.Println((valor_uno == valor_dos), "||", (valor_dos == valor_dos), (valor_uno == valor_dos) || (valor_dos == valor_dos))
	fmt.Println((valor_uno == valor_dos), "||", (valor_uno == valor_dos), (valor_uno == valor_dos) || (valor_uno == valor_dos))

	// Operador NOT
	fmt.Println("Operador NOT")
	fmt.Println("!", (valor_uno == valor_uno), "=", !(valor_uno == valor_uno))
	fmt.Println("!", (valor_uno == valor_dos), "=", !(valor_uno == valor_dos))
}
