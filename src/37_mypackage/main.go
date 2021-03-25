package main

import (
	"fmt"
	"golang-bootcamp/src/37_mypackage/car"
)

func main() {
	var myCar car.Car
	myCar.Brand = "Ferrari"
	myCar.Year = 2020

	fmt.Println(myCar)

	car.PrintMessage("Funcion Publica")

	//	var myOtherCar car.carPrivate
	//	fmt.Println(myOtherCar)
}
