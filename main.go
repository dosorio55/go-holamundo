package main

import (
	"fmt"

	"github.com/dosorio55/hola-mundo/goRoutines"
)

func main() {
	channel1 := make(chan bool)
	defer func() {
		<-channel1
	}()
	go goRoutines.GoRoutines(channel1)

	fmt.Println("Hola Mundo!")

}
