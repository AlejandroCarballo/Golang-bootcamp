package main

import "fmt"

type pc struct {
	ram  int
	os   string
	disk int
}

func (myPC pc) string() string {
	return fmt.Sprintf("Tengo %d GB de RAM, %d GB de Disco y es un sistema %s", myPC.ram, myPC.disk, myPC.os)
}

func main() {

	myPc := pc{ram: 16, os: "Linux", disk: 100}

	fmt.Println(myPc)

}
