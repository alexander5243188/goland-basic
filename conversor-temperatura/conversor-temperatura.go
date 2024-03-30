package main

import "fmt"

// Conversor de temperatura: Convierte grados Celsius a Fahrenheit y viceversa.

func conversor_grados_celsius(grados_fahrenheit float32) float32 {
	//°C = (°F -32) *(5/9) O °C = (°F -32) / 1.8
	celsius := ((grados_fahrenheit - 32) / 1.8)
	return celsius
}
func conversor_grados_fahrenheit(grados_celsius float32) float32 {
	//°F = (°C × 1.8) + 32
	fahrenheit := (grados_celsius * 1.8) + 32
	return fahrenheit
}

//func grados_fahrenheit(grados float32) float32 {}

func main() {
	//var grados_celsius float32
	var opcion int

	fmt.Println("Seleccione: 1 Grados Celsius, 2 Grados Fahrenheit")
	fmt.Scan(&opcion)

	switch {

	case opcion == 1:

		var lectura_grados_celsius float32
		fmt.Println("Ingrese los grados Celosius:")
		fmt.Scan(&lectura_grados_celsius)

		if lectura_grados_celsius >= -273.15 && lectura_grados_celsius <= 100 {
			fmt.Println(lectura_grados_celsius, "°C", conversor_grados_fahrenheit(lectura_grados_celsius), "°F")
		} else {
			fmt.Println("Ingrese datos correctos.")
		}
	case opcion == 2:

		var lectura_grados_fahrenheit float32
		fmt.Println("Ingrese los grados Fahrenheit:")
		fmt.Scan(&lectura_grados_fahrenheit)

		if lectura_grados_fahrenheit >= -32 && lectura_grados_fahrenheit <= 212 {
			fmt.Println(lectura_grados_fahrenheit, "°F", conversor_grados_celsius(lectura_grados_fahrenheit), "°C")
		} else {
			fmt.Println("Ingrese datos correctos")
		}
	}
}
