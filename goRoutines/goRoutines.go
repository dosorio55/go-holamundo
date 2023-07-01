package goRoutines

import "time"

func GoRoutines(channel1 chan bool) {
	myArray := []string{"a", "b", "c", "d", "e", "f", "g", "h"}

	for index, letra := range myArray {
		time.Sleep(300 * time.Millisecond)
		println(index, letra)
	}

	channel1 <- true
}