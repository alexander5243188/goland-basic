package main
import {
	"fmt"
	"os"
}
funv main(){
	envVar := os.Getenv("MY_ENV_VAR")
	if envVar == "" {
		fmt.Println("La variable de entorno HOME no está definida.")
	} else {
		fmt.Printf("El valor de la variable de entorno HOME es: %s\n", envVar)
	}

	file, err := os.Create("archivo.txt")
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}
	defer file.Close()
}