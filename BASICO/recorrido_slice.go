package main

import (
	"fmt"
	"strings"
)

func palindromo(palabra string) {
	var textReverse string

	for i := len(palabra) - 1; i >= 0; i-- {
		textReverse += string(palabra[i])
		fmt.Println(textReverse)
	}
	if palabra == textReverse {
		fmt.Println("SI es un palindromo")
	} else {
		fmt.Println("NO es un palindromo")
	}
}

func main() {
	var slice []string = []string{"hola", "que", "hace"}
	//recorrido del slice con RANGE
	for i, valor := range slice {
		fmt.Println(i, valor)
	}

	//pone todo en minuscula(strings.ToLower)
	palindromo(strings.ToLower("Ama"))
}
