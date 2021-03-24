package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPc pc) ping() {
	fmt.Println(myPc.brand, "pong")

}

func (myPc *pc) duplicateRame() {
	myPc.ram = myPc.ram * 2

}

func (myPc *pc) duplicateDisk() {
	myPc.ram = myPc.disk * 2

}

func main() {

	var a int = 50
	b := &a

	fmt.Println(*b, b)

	*b = 100
	fmt.Println(a)

	myPc := pc{ram: 16, disk: 200, brand: "linux"}
	fmt.Println(myPc)

	myPc.ping()

	fmt.Println(myPc)
	myPc.duplicateDisk()
	myPc.duplicateRame()

	fmt.Println(myPc)

}
