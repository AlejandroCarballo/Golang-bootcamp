package Car

import (
	"fmt"
)

func main() {
	var myCar car.car
	myCar.Brand = "Ferrari"
	myCar.Year = 2020

	fmt.Println(myCar)

	car.PrintMessage("Funcion Publica")

	car.printMessage("Funcion Privada")

	//	var myOtherCar car.carPrivate
	//	fmt.Println(myOtherCar)
}
