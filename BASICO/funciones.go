package main

import "fmt"

func normalFunction(mensaje string) {
	fmt.Println(mensaje)
}

func tripleArgumento(a, b int, c string) {
	fmt.Println(a, b, c)
}
func returnValue(a int) int {
	return a * 2
}
func dobleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	fmt.Println("Hola2")
	fmt.Println("Hola3")
	normalFunction("Hola en NormalFunction")
	tripleArgumento(3, 5, "Hola en tripleFunction")
	fmt.Printf("el resultado es: %d\n", returnValue(2))
	var valor1, _ int = dobleReturn(4)
	fmt.Printf("el resultado 1 es: %d y el resultado 2 es: _", valor1)
}
