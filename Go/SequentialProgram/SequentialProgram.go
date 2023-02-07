package main

import (
	"fmt"
	"time"
)

func CalculateSquare(i int) {
	time.Sleep(1 * time.Second)
	fmt.Println(i * i)
}

func main() {
	start := time.Now()
	for i := 1; i <= 1000; i++ {
		CalculateSquare(i)

	}
	elapsed := time.Since(start)
	fmt.Println("Elapsed Time :", elapsed)

}
