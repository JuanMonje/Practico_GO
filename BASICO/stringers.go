package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) String() string {
	return fmt.Sprintf("Tengo %d GB RAM, %d GB Disco y es una %s", myPC.ram, myPC.disk, myPC.brand)
}

func main() {
	myPC := pc{ram: 16, disk: 100, brand: "msi"}

	fmt.Println(myPC)

	//otras maneras de visualisar datos de Structs
	fmt.Printf("%+v\n", myPC) //quitar la funcion String()
	fmt.Printf("%#v\n", myPC)
}
