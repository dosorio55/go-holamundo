package main

import (
	"fmt"
	"github.com/dosorio55/hola-mundo/variables"
	"runtime"
)

func main() {
	variables.ShowIntegers()
	variables.CreatePerson()

	state, newString := variables.ConvertNumberToString(10023216546546)
	fmt.Println(state, newString)

	os := runtime.GOOS

	// this version of the if assaigns the variable in the same line of the condition then creates the condition
	/* 	if os := runtime.GOOS; os == "Linux." {
	   	} else if os == "Windows" {

	   	} */

	switch os {
	case "Linux":
		fmt.Println("Linux")
	case "Windows":
		fmt.Println("windows")
	default:
		fmt.Println("No se reconoce el sistema operativo")
	}
}
