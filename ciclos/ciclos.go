package main

import "fmt"

func main() {
	// for condicional
	for i := 0; i < 6; i++ {
		fmt.Println("-", i)
	}

	// for while
	contador := 0
	for contador < 3 {
		fmt.Println("--", contador)
		contador++
	}
}
