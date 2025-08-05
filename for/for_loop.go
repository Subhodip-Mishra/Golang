package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now() // record start time

	for i := 0; i < 300000; i++ {
		fmt.Println(i)
	}

	elapsed := time.Since(start) // calculate elapsed time
	fmt.Printf("Execution time: %s\n", elapsed)
}
