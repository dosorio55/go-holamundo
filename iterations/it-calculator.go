package iterations

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func IterationsCalculator() string {
	fmt.Println("Ingrese el primer número: ")
	scanner := bufio.NewScanner(os.Stdin)
	var myNumber int
	var error error
	var message string

	for {
		if scanner.Scan() {
			myNumber, error = strconv.Atoi(scanner.Text())

			if error != nil {
				fmt.Println("se ha producido un error" + error.Error())
				continue
			} else {
				break
			}
		}
	}

	for index := 1; index <= 10; index++ {
		result := myNumber * index
		message += fmt.Sprintf("%d x %d = %d\n", myNumber, index, result)
	}

	return message
}
