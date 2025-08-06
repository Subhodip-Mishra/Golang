package main

import (
	"fmt"
	"time"
)

func main() {
	// simple switch
	i := 1

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("other")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend")
	case time.Monday:
		fmt.Println("its monday")
	case time.Tuesday:
		fmt.Println("its tuesday")
	case time.Wednesday:
		fmt.Println("its wednesday")
	case time.Thursday:
		fmt.Println("its thusday")
	case time.Friday:
		fmt.Println("its friday")
	default:
		fmt.Println("it's workday")
	}

	
}
