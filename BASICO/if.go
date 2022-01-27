package main

import "fmt"

func main() {
	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("es 1")
	} else {
		fmt.Println("No es 1")
	}

	//operador AND
	if valor1 == 2 && valor2 == 2 {
		fmt.Println("Es verdad")
	} else {
		fmt.Println("no es verdad")
	}

	//Operador OR
	if valor1 == 0 || valor2 == 3 {
		fmt.Println("Es verdad OR")
	} else {
		fmt.Println("Es mentira OR")
	}
}
