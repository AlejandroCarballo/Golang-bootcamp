package main

import "fmt"

func main() {

	// Declaracion de constantes

	const pi float64 = 3.14
	const pi2 = 3.1415
	fmt.Println("pi", pi)
	fmt.Println("pi2", pi2)

	// Declaracion de variables enteras
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	//Zero Values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//Area cuadrado

	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado:", areaCuadrado)

	x := 10
	y := 50
	result := x + y

	//Suma
	fmt.Println("Suma", result)

	result = y - x
	//Resta
	fmt.Println("Resta", result)

	result = y * x
	//Multiplicacion
	fmt.Println("Multicplicacion:", result)

	result = y / x
	//Division
	fmt.Println("Division:", result)

	result = y % x
	//Modulo
	fmt.Println("Modulo:", result)

	//Incremental
	x++
	fmt.Println("Incremental:", x)

	//Decremental
	x--
	fmt.Println("Decremental:", x)

}
