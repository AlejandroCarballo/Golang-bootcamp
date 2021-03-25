package car

import (
	"fmt"
)

//My Public Car
type Car struct {
	Brand string
	Year  int
}
type CarPrivate struct {
	brand string
	year  int
}

// PrinteMessage es para imprimir un mensaje
func PrintMessage(text string) {
	fmt.Println(text)
}

func printMessage(text string) {
	fmt.Println(text)
}
