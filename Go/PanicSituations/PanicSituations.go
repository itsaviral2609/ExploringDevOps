package main

func main() {
	ch := make(chan int, 10)
	ch <- 10
	ch <- 11
	close(ch)
	ch <- 12
	close(ch)
}
