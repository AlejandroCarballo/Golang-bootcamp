package main

import "fmt"

func main() {
	if isEven(7) {
		fmt.Println("Number is pair")
	} else {
		fmt.Println("Number is not pair")
	}

}

func isEven(num int) bool {
	return num%2 == 0
}
