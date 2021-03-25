package main

import "fmt"

func main() {
	//Declaracion de variables
	helloMessage := "hello"
	worldMessage := "World"

	//Printl
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	//Printf
	nombre := "Deloitte"
	proyecto := 100
	fmt.Printf("%s tiene mas de %d proyectos\n", nombre, proyecto)
	fmt.Printf("%v tiene mas de %d proyectos\v", nombre, proyecto)

	//Sprinff
	message := fmt.Sprintf("%s tiene mas de %d proyectos", nombre, proyecto)

	fmt.Println(message)

	//Tipo de datos
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("helloMessage: %T\n", proyecto)
}
