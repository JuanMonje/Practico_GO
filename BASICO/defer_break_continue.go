package main

import "fmt"

func main() {
	//Defer (hace que X linea sea la ultima en ejecutarse)
	defer fmt.Println("linea con DEFER")
	fmt.Println("linea limpia")

	//Continue y break
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		//continue
		if i == 2 {
			fmt.Println("es 2")
			continue
		}

		//break
		if i == 8 {
			fmt.Println("entró al break")
			break
		}
	}
}
