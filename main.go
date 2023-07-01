package main

import (
	// "github.com/dosorio55/hola-mundo/users"
	"fmt"
	// I can add a name to name the package and abreviate it
	in "github.com/dosorio55/hola-mundo/interfaces"
	"github.com/dosorio55/hola-mundo/models"
)

func main() {
	human := new(models.Person)

	aliveHuman := in.MyHuman(human)

	fmt.Println(aliveHuman)

	// users.CreateUser()
}
