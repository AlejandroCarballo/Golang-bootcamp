package main

import "fmt"

func main() {

	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
	}
}
