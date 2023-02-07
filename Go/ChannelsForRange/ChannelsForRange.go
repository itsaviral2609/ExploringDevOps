package main

import (
	"fmt"
)

/*
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan int)
	go sell(ch, &wg)
	go buy(ch, &wg)
	wg.Wait()
}

func sell(ch chan int, wg *sync.WaitGroup) {
	ch <- 1
	ch <- 2
	fmt.Println("Sending all the data: ")
	close(ch)
	wg.Done()
}

func buy(ch chan int, wg *sync.WaitGroup) {
	fmt.Println("Waiting for the data")
	for val := range ch {
		fmt.Println("Value received from channel", val)
	}
	wg.Done()
}
*/

func main() {
	ch := make(chan int, 5)
	ch <- 10
	ch <- 12
	close(ch)

	for val := range ch {
		fmt.Println(val)
	}
}
