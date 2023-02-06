package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan int, 3)
	go sell(ch, &wg)
	wg.Wait()
}

func sell(ch chan int, s *sync.WaitGroup) {
	go buy(ch, s)
	fmt.Println("Sent all the data to the channel: ")
	s.Done()
}

func buy(ch chan int, s *sync.WaitGroup) {
	fmt.Println("Sending the data: ")
	fmt.Println("Received the data: ", <-ch)
	s.Done()
}
