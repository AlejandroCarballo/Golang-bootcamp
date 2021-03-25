package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Alejandro"] = 34
	m["Melisa"] = 35

	//fmt.println(m)

	//Recorrer map
	for index, value := range m {
		fmt.Println(index, value)
	}
}
