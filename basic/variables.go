package main

import "fmt"

func main() {
	fmt.Printf("---- GO Variables ----\n")
	var greeting string
	greeting = "Hello variables"
	fmt.Println("message : " + greeting)
	var count int
	count = 10
	fmt.Println(count)

	var is_running bool
	is_running = true
	fmt.Println("Running Status : ", is_running)
	var firstName, lastName string
	firstName = "John"
	lastName = "Doe"
	fmt.Println(firstName, lastName)
}
