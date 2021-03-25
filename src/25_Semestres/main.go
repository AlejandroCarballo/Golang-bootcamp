package main

import "fmt"

func main() {
	var mes int8 = 4
	switch {
	case mes <= 3:
		fmt.Println("Primer trimestre")
	case mes > 3 && mes <= 6:
		fmt.Println("Segundo trimestre")
	case mes > 6 && mes <= 9:
		fmt.Println("Tercer trimestre")
	case mes > 9 && mes <= 12:
		fmt.Println("Cuarto trimestre")
	default:
		fmt.Println("Ese no es un valor valido, el maximo valor es 12.")

	}
}
