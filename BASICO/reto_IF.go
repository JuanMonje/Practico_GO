package main

import "fmt"

func main() {
	var numero int = 15
	var par int = numero % 2

	if par == 0 {
		fmt.Printf("El numero %d es par\n", numero)
		fmt.Printf("El modulo es %d", par)
	} else {
		fmt.Printf("El numero %d es IMpar\n", numero)
		fmt.Printf("El modulo es %d\n", par)
	}
	//verificar palabra clave
	passOrigin := "xiaomi"
	fmt.Printf("la clave es %t", password(passOrigin))
}

func password(pass string) bool {
	var pass1 string = "beto"
	var pass2 string = "monje"
	var pass3 string = "xiaomi"
	if pass == pass1 {
		fmt.Println("BETO")
		return true
	} else if pass == pass2 {
		fmt.Println("MONJE")
		return true
	} else if pass == pass3 {
		fmt.Println("XIAOMI")
		return true
	} else {
		fmt.Println("ESTA MAL LA CLAVE")
		return false
	}

}
