package main

import "fmt"

func main() {
	var i uint8 = 0
	for i < 10 {
		fmt.Println(i)
		i++

		//continue
		if i == 2 {
			fmt.Println("El numero que sigue es par!!")
			continue
		}

	}

}
