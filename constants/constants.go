package main

import "fmt"

func main() {
	// const name string = "golang"
	// const age = 20

	const (
		port = 3000
		host = "localhost"
	)

	fmt.Println(port, host)
}