package main

import (
	"fmt"
)

func main() {
	const pi float64 = 3.1416
	h := 3.1
	fmt.Println(pi)
	fmt.Println(h)

	//Imprimir con formato
	var nombre string = "platzi"
	var cursos int16 = 500
	fmt.Printf("%s tiene mas de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene mas de %v cursos\n", nombre, cursos)

	//concatenar formatos y guardarlas en variable
	var mensaje1 string = fmt.Sprint(nombre, " prueba SIN ", cursos, " f")
	fmt.Println(mensaje1)
	var mensaje2 string = fmt.Sprintf("%s con Sprint %d f", nombre, cursos)
	fmt.Println(mensaje2)

	//Saber que tipo de variable es
	buleano := false
	fmt.Printf("nombre: %T\n", nombre)
	fmt.Printf("crusos: %T\n", cursos)
	fmt.Printf("buleano: %T\n", buleano)
	fmt.Printf("Estado del buleano: %t", buleano)
}
