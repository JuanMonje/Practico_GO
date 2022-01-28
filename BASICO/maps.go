package main

import "fmt"

func main() {
	/* m := make(map[string]int)
	m["jose"] = 14
	m["pepito"] = 20
	m["sebastian"] = 26 */

	//es lo mismo que lo de arriba
	m := map[string]int{
		"jose":      14,
		"pepito":    20,
		"sebastian": 26,
	}

	fmt.Println(m)
	//Recorrer map
	for i := range m {
		fmt.Println(i)
	}

	//Encontrar X valor
	valor, err := m["jose"]
	fmt.Println(valor, err)
}
