package files

import (
	"bufio"
	"fmt"
	"strconv"
	"io/ioutil"
	"github.com/dosorio55/hola-mundo/iterations"
	"os"
)

var filePathName string = "./files/table.txt"

func PromptUser() {
	fmt.Println("que quieres hacer: 1: agregar tabla a archivo, 2: leer archivo")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if scanner.Scan() {
			option, error := strconv.Atoi(scanner.Text())
			if error != nil {
				fmt.Println("se ha producido un error" + error.Error())
				continue
			}

			switch option {
			case 1:
				AddToFile()
				break
			case 2:
				ReadFile()
				break
			default:
				fmt.Println("opcion no valida, por favor introduzca una número entre las opciones")
			}
			break
		} else {
			fmt.Println("se ha producido un error")
			continue
		}
	}
}

func AddToFile() {
	writeToFile := true
	for {
		file, error := os.OpenFile(filePathName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if writeToFile {
			text := iterations.IterationsCalculator()
			if error != nil {
				fmt.Println("se ha producido un error" + error.Error())
				return
			}

			fmt.Fprintln(file, text)
			fmt.Println("se ha escrito en el archivo", text, file)
			writeToFile = askUserForInput()
		} else {
			file.Close()
			break
		}

	}

}

func ReadFile() {
	// file, error := os.OpenFile(filePathName, os.O_RDONLY, 0644)
	// this one reads line by line with bufio and we need to use a for loop
	// file, error := os.Open(filePathName)
	// this ioutil is deprecated and it is better to use the os always,
	// but also in the course says that ioutil reads the whole file which is not recommended
	file, error := ioutil.ReadFile(filePathName)
	if error != nil {
		fmt.Println("se ha producido un error" + error.Error())
		return
	}

	fmt.Println(string(file))
}

func askUserForInput() bool {
	fmt.Println("Quieres seguir agregando números? (s/n)")
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		text := scanner.Text()
		if text == "s" {
			return true
		} else {
			return false
		}
	} else {
		fmt.Println("se ha producido un error")
		return false
	}
}
