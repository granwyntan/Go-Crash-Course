package main

import (
	"fmt"
	"strconv"
)

// Person struct definition
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	// Could also be done by:
	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value reciever)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age) + "."
}

// hasBirthday Method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer reciever)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender != "Male" {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Lorem", lastName: "Ipsum", city: "Generic City Name", gender: "Female", age: 25}
	person2 := Person{firstName: "Placeholder", lastName: "Text", city: "Generic City Name", gender: "Male", age: 27}

	// // Alternative
	// person1 := Person{"Lorem", "Ipsum", "Generic City Name", "Female", 25}

	// fmt.Println(person1)
	// fmt.Println(person1.firstName)
	// person1.age++
	// fmt.Println(person1)

	person1.hasBirthday()
	person1.getMarried(person2.lastName)
	person2.getMarried(person1.lastName)

	fmt.Println(person1.greet())
}
