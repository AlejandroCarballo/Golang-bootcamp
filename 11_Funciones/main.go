package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)

}
func returnValue(a int) int {
	return a * 2

}

func doubleReturn(a int) (c, d int) {
	return a, a * 2

}

func main() {
	normalFunction("hola mundo")
	tripleArgument(1, 3, "hola mundo")

	value := returnValue(2)
	fmt.Println("value", value)

	value1, _ := doubleReturn(2)

	fmt.Println("Value1:", value1)
}
