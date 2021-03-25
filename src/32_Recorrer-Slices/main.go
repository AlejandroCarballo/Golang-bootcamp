package main

import "fmt"

func main() {

	slice := []string{"Hola", "GO"}
	for _, valor := range slice {
		fmt.Println(valor)
	}
}
