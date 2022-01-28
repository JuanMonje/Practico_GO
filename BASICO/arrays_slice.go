package main

import "fmt"

func main() {
	//Array
	var arreglo [4]int
	//								tamano del arreglo// capacidad maxima del arreglo
	fmt.Println(arreglo, len(arreglo), cap(arreglo))

	//Slice (es lo mismo pero sin definir el maño del arreglo)
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(slice, len(slice), cap(slice))

	//Slicing
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[3:5])
	fmt.Println(slice[4:])

	//Appened
	newSlice := []int{9, 10, 11, 12}
	//textSlice := []string{"beto", "monje", "xiaomi"}
	//slice = append(slice, 8)
	slice = append(slice, newSlice...)
	fmt.Println(slice)

}
