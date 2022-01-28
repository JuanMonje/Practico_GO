package main

import (
	"fmt"
	"paquete"
)

//modulos de librerias
func main() {
	var myCar paquete.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 2020
	fmt.Println(myCar)

	paquete.PrintMessage()
}
