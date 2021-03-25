package main

import (
	"fmt"
	"sync"
)

func say(text string, wg *sync.WaitGroup) {

	defer wg.Done()
	fmt.Println(text)

}
func main() {
	var wg sync.WaitGroup

	fmt.Println("hello")
	wg.Add(1)
	go say("World", &wg)

	wg.Wait()

}
