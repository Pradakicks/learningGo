package main

import (
	"fmt"
	"strconv"
)

// Define Person struct

type Person struct {
	firstName string 
	lastName string
	city string
	gender string
	age int

}

// Greeting method (value receiver)

func (p Person) greet () string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)



func main() {
	// init person using struct

	 person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "f", age: 25}
	// // Alternative	
	// // person2 := Person{"Samantha", "Smith", "Boston", "f", 25}

	// // fmt.Println(person1)

	// fmt.Println(person1.firstName)
	// person1.age++
	// fmt.Println(person1)


	fmt.Println(person1.greet())
}
