package main

import "fmt"

func main() {
	// defer
	defer fmt.Println("Yo ")
	defer fmt.Println("estoy")
	fmt.Println("programando")
	defer fmt.Println("en")
	fmt.Println("GO")
}
