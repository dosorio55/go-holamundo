package variables

import (
	"fmt"
	"time"
	"strconv"
)

var Name string
var State bool
var Salary float32
var Date time.Time

func CreatePerson () {
	Name = "Daniel"
	State = true
	Salary = 1000.50
	Date = time.Now()

	fmt.Println("Name:", Name)
	fmt.Println("State:", State)
	fmt.Println("Salary:", Salary)
	fmt.Println("Date:", Date)
}

func ConvertNumberToString (number int) (bool,string) {
	texto := strconv.Itoa(number)

	return true, texto
}
