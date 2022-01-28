package main

import "fmt"

func main() {
	/* modulo := 4 % 2
	switch modulo { */
	switch modulo := 4 % 2; modulo {
	case 0:
		fmt.Println("es par")
	default:
		fmt.Println("es IMpar")
	}

	//Switch sin definir variable en Switch
	valor := 50
	switch {
	case valor > 100:
		fmt.Println("Es mayor que 100")
	case valor < 0:
		fmt.Println("Es menor que cero")
	default:
		fmt.Println("No es condicion")
	}

}
