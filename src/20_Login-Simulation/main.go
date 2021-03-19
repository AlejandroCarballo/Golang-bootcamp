package main

import "fmt"

func main() {

	fmt.Println("El estado es:", login("Alejandro", "constraseña no es segura"))

}

func login(username, password string) bool {

	const usernameStorage string = "Alejandro"
	const passwordStorage string = "constraseña no es segura"

	if username == usernameStorage && password == passwordStorage {
		return true
	} else {
		return false
	}

}
