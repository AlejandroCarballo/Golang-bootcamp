package main

import "fmt"

func AreaRectangulo(baseRectangulo, alturaRectangulo int) {

	var areaRectangulo int = baseRectangulo * alturaRectangulo
	fmt.Println(areaRectangulo)

}

func AreaTrapecio(baseMenor, baseMayor, altura int) {
	var areaTrapecio int = (baseMenor + baseMayor) * altura / 2
	fmt.Println(areaTrapecio)
}

func AreaCirculo(radio float64) {
	const pi float64 = 3.14
	var areaCirculo float64 = pi * radio * radio
	fmt.Println(areaCirculo)
}

func main() {

	AreaRectangulo(10, 20)
	AreaTrapecio(15, 7, 25)
	AreaCirculo(50)

}
