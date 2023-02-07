package main

import "fmt"

func main() {
	go func() {
		fmt.Println("This is a anaonymous Function")
	}()
	time.Sleep(1 * time.Second)
}
