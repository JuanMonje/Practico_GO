package main

import "fmt"

func main() {
	//condicion clasica del FOR
	for i := 20; i >= 0; i-- {
		fmt.Println(i)
	}

	//FOR como while
	contador := 0
	for contador <= 10 {
		fmt.Println("FOR while: ", contador)
		contador++
	}

	//FOR forever, ciclo FOR infinito
	contadorInfinito := 0
	for {
		fmt.Println(contadorInfinito)
		contadorInfinito++
	}
}
