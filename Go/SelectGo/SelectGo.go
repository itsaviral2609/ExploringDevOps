package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go goOne(ch1)
	go goTwo(ch2)
	select {
	case val1 := <-ch1:
		fmt.Println(val1)
		break
		fmt.Println("After break")

	case val2 := <-ch2:
		fmt.Println(val2)

	default:
		fmt.Println("Default will be executed!")

	}
	time.Sleep(1 * time.Second)
}
func goOne(ch1 chan string) {
	ch1 <- "Channel1"
}
func goTwo(ch2 chan string) {
	ch2 <- "Channel2"
}
