package main

import (
	"fmt"

	"github.com/dosorio55/hola-mundo/ejercicios"
	// "github.com/dosorio55/hola-mundo/variables"
	// "runtime"
)

func main() {
	myString, message := ejercicios.ReturnValues("10")
	fmt.Println(myString, message)
}
