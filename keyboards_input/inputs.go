package keyboardsinput

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func KeyboardsInput() {
	fmt.Println("Ingrese el primer número a sumar: ")
	var error error
	var input1 int
	var input2 int
	// the os.Stdin is the keyboard input, it stands for standard input
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		input1, error = strconv.Atoi(scanner.Text())

		if error != nil {
			fmt.Println("se ha producido un error" + error.Error())
		}
		fmt.Println("El primer número es: input")
	} else {
		fmt.Println("No se ingresó ningún número")
		return
	}

	fmt.Println("Ingrese el segundo número a sumar: ")

	if scanner.Scan() {
		input2, error = strconv.Atoi(scanner.Text())

		if error != nil {
			fmt.Println("se ha producido un error" + error.Error())
		}
	} else {
		fmt.Println("No se ingresó ningún número")
		return
	}

	fmt.Println("El resultado de la suma es: ", input1+input2)
}
