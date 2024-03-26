package main

import "fmt"

func main() {
	// defer
	defer fmt.Println("Yo ")
	defer fmt.Println("estoy")
	fmt.Println("programando")
	defer fmt.Println("en")
	fmt.Println("GO")

	//continue, break
	count := 10
	for i := 0; i < count; i++ {
		fmt.Println(i)
		if i == 3 {
			fmt.Println("-", i, "-")
			continue
		}
		if i == 5 {
			fmt.Println("No terminaras el ciclo JAJAJAJa")
			break
		}
	}
}
