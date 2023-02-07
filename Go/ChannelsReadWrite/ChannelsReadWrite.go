package main
import (
	"fmt"
	"time"
)
func main() {
	ch := make(chan string)
	go sell(ch)
	go buy(ch)
	time.Sleep(2 * time.Second)
}
// sending info to the channel
func sell(ch chan string) {
	ch <- "Furniture"
	fmt.Println("Sent data to the channel")
}
// recieve data from the channel
func buy(ch chan string) {
	fmt.Println("Waiting for the data")
	val := <-ch
	fmt.Println("Recieved data--->", val)
}
