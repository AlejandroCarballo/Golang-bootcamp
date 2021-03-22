package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Alejandro"] = 34
	m["Melisa"] = 35

	//fmt.println(m)

	//Recorrer map
	value, ok := m["Alejandro"]
	fmt.Println(value, ok)
}
