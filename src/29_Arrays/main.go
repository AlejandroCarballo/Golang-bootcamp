package main

import "fmt"

func main() {
	//Array
	var array [4]int
	array[0] = 100
	array[1] = 10
	fmt.Println(array, len(array), cap(array))

}
