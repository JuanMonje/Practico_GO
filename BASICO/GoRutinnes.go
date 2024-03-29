package main

import (
	"fmt"
	"sync"
)

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}

func main() {
	var wg sync.WaitGroup

	fmt.Println("Hellow")
	wg.Add(1)

	go say("world", &wg)
	wg.Wait()

	//time.Sleep(time.Second*1)

	//Funcion anonima
	go func() {
		fmt.Println("Adios")
	}()
}
