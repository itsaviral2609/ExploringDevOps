package main

import (
	"fmt"
	"time"
)

func main() {
	go start()
	time.Sleep(1 * time.Second)

}

func start() {
	go process()
	fmt.Println("Starting Process!")
}

func process() {
	fmt.Println("Process xyz")
}
