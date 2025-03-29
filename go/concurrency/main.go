package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	// ch := make(chan int64)
	loop()
	loop()

	// <-ch

	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s\n", elapsed)
}

func loop() {
	count := int64(0)
	for i := range int64(1e9) {
		count += i
	}
	// ch <- count
}
