package main

import "fmt"



func main() {
	// var name = "Adrian"
	var age int32 = 37
	const isCool = true
	name := "Adrian"
	var size float32 = 2.3

	
	name, email := "Brad", "brad@gmail.com"
	fmt.Println(name, age, isCool, size, email)
	fmt.Printf("%T\n", size)

	
}
