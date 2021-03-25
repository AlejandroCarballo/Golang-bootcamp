package main

import "fmt"

type car struct {
	brand   string
	year    int
	seating int
	color   string
	owner   string
}

func main() {
	myCar := car{brand: "Toyota", year: 2018, seating: 4, color: "Rojo", owner: "Alejandro Carballo"}
	fmt.Println("Los datos de mi auto son: ", myCar)

	var myOtherCar car
	myOtherCar.brand = "Honda"
	fmt.Println(myOtherCar)
}
