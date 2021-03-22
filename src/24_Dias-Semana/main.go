package main

import "fmt"

func main() {
	var dia int8 = 4
	switch dia {
	case 1:
		fmt.Println("Hoy es lunes")
	case 2:
		fmt.Println("Hoy es martes")
	case 3:
		fmt.Println("Hoy es miercoles")
	case 4:
		fmt.Println("Hoy es jueves")
	case 5:
		fmt.Println("Hoy es viernes")
	case 6:
		fmt.Println("Hoy es sabado")
	case 7:
		fmt.Println("Hoy es domindo")

	default:
		fmt.Println("Este no es un dia valido de la semana")

	}
}
