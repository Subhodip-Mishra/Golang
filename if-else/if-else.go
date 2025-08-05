package main

import "fmt"

func main() {
	age := 15

	if age >= 18 {
		fmt.Println("person is adult")
	} else {
		fmt.Println("Person is not adult")
	}
	
	// age := 11

	if age >= 18 {
		fmt.Println("person is adult")
	} else if age >= 12 {
		fmt.Println("perosn is teeneger")
	} else {
		fmt.Println("person is a kid")
	}

	var role = "admin"
	var hasPermission = true

	if role == "admin" && hasPermission {
		fmt.Println("yes")
	}

	// we can declare a variable inside if	 
	if age := 20; age >= 18 {
		fmt.Println("peron is adult", age)
	} else if age >= 12 {
		fmt.Println("peron is teeneger", age)
	}

	// go does not have ternary, you will have to use normal if else
}